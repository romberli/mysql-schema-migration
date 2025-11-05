package migration

import (
	"os"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/sql/parser"
	"github.com/romberli/go-util/viper"
	"github.com/romberli/log"

	"github.com/romberli/mysql-schema-migration/config"
)

type Controller struct {
	Parser *parser.Parser
}

func NewController() *Controller {
	p := parser.NewParserWithDefault()
	p.SetParseTableDefinition(true)

	return &Controller{
		Parser: p,
	}
}

func (c *Controller) GetSchemaMigrationSQLList() ([]string, error) {
	diffList, err := c.GetDiff()
	if err != nil {
		return nil, err
	}

	var sqlList []string
	for _, diff := range diffList {
		sql := diff.GetTableMigrationSQL()
		if sql != constant.EmptyString {
			sqlList = append(sqlList, sql)
		}
	}

	return sqlList, nil
}

func (c *Controller) GetDiff() ([]*parser.TableDefinitionDiff, error) {
	sourceTableDefinitions, err := c.GetSourceTableDefinitions()
	if err != nil {
		return nil, err
	}
	targetTableDefinitions, err := c.GetTargetTableDefinitions()
	if err != nil {
		return nil, err
	}

	var diffList []*parser.TableDefinitionDiff

Loop:
	for _, sourceTableDefinition := range sourceTableDefinitions {
		var targetTableDefinition *parser.TableFullDefinition
		for _, targetTableDefinition = range targetTableDefinitions {
			if sourceTableDefinition.Table.TableName == targetTableDefinition.Table.TableName {
				tableDiff := targetTableDefinition.Diff(sourceTableDefinition)
				diffList = append(diffList, tableDiff)
				continue Loop
			}
		}

		tableDiff := parser.NewTableDiff(parser.TableDiffTypeDrop, sourceTableDefinition.Table, nil)
		diffList = append(diffList, parser.NewTableDefinitionDiff(sourceTableDefinition.Table.GetFullTableName(), constant.EmptyString, tableDiff, nil, nil))
	}

	for _, targetTableDefinition := range targetTableDefinitions {
		if !c.tableDefinitionExists(sourceTableDefinitions, targetTableDefinition) {
			tableDiff := parser.NewTableDiff(parser.TableDiffTypeCreate, nil, targetTableDefinition.Table)
			diffList = append(diffList, parser.NewTableDefinitionDiff(constant.EmptyString, targetTableDefinition.Table.GetFullTableName(), tableDiff, nil, nil))
		}
	}

	return diffList, nil
}

func (c *Controller) GetSourceTableDefinitions() ([]*parser.TableFullDefinition, error) {
	var (
		sqlList []string
		err     error
	)

	t := viper.GetString(config.SourceTypeKey)
	switch t {
	case config.TypeFile:
		sqlList, err = c.GetCreateTableSQLFromFile(viper.GetString(config.SourceFileKey))
		if err != nil {
			return nil, err
		}
	case config.TypeDB:
		sqlList, err = c.GetCreateTableSQLFromDB(
			viper.GetString(config.SourceDBAddrKey),
			viper.GetString(config.SourceDBNameKey),
			viper.GetString(config.SourceDBUserKey),
			viper.GetString(config.SourceDBPassKey),
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("unknown source type. type: %s", t)
	}

	tableDefinitions, err := c.ParseTableDefinitions(sqlList)
	if err != nil {
		return nil, err
	}

	return c.FilterTables(tableDefinitions), nil
}

func (c *Controller) GetTargetTableDefinitions() ([]*parser.TableFullDefinition, error) {
	var (
		sqlList []string
		err     error
	)

	t := viper.GetString(config.TargetTypeKey)
	switch t {
	case config.TypeFile:
		sqlList, err = c.GetCreateTableSQLFromFile(viper.GetString(config.TargetFileKey))
		if err != nil {
			return nil, err
		}
	case config.TypeDB:
		sqlList, err = c.GetCreateTableSQLFromDB(
			viper.GetString(config.TargetDBAddrKey),
			viper.GetString(config.TargetDBNameKey),
			viper.GetString(config.TargetDBUserKey),
			viper.GetString(config.TargetDBPassKey),
		)
	default:
		return nil, errors.Errorf("unknown target type. type: %s", t)
	}

	tableDefinitions, err := c.ParseTableDefinitions(sqlList)
	if err != nil {
		return nil, err
	}

	return c.FilterTables(tableDefinitions), nil
}

func (c *Controller) GetCreateTableSQLFromFile(filePath string) ([]string, error) {
	exists, err := linux.PathExists(filePath)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}

	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return c.Parser.Split(string(sqlBytes))
}

func (c *Controller) GetCreateTableSQLFromDB(addr, dbName, dbUser, dbPass string) ([]string, error) {
	conn, err := mysql.NewConn(addr, dbName, dbUser, dbPass)
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := conn.Close()
		if closeErr != nil {
			log.Errorf("close db connection failed. addr: %s, dbName: %s, dbUser: %s, error:\n%+v",
				addr, dbName, dbUser, closeErr)
		}
	}()
	r := NewRepository(conn)

	tables, err := r.GetTableNames()
	if err != nil {
		return nil, err
	}

	var sqlList []string
	for _, table := range tables {
		sql, err := r.GetCreateTableSQL(table)
		if err != nil {
			return nil, err
		}
		sqlList = append(sqlList, sql)
	}

	return sqlList, nil
}

func (c *Controller) ParseTableDefinitions(sqls []string) ([]*parser.TableFullDefinition, error) {
	var tableDefinitions []*parser.TableFullDefinition
	for _, sql := range sqls {
		tableDefinition, err := c.Parser.ParseTableDefinition(sql)
		if err != nil {
			return nil, err
		}

		tableDefinitions = append(tableDefinitions, tableDefinition)
	}

	return tableDefinitions, nil
}

func (c *Controller) FilterTables(tableDefinitions []*parser.TableFullDefinition) []*parser.TableFullDefinition {
	var filteredTables []*parser.TableFullDefinition
	for _, tableDefinition := range tableDefinitions {
		if IsTableIncluded(tableDefinition.Table.TableName) {
			filteredTables = append(filteredTables, tableDefinition)
		}
	}

	return filteredTables
}

func (c *Controller) tableDefinitionExists(tableDefinitions []*parser.TableFullDefinition,
	tableDefinition *parser.TableFullDefinition) bool {
	for _, table := range tableDefinitions {
		if table.Table.TableName == tableDefinition.Table.TableName {
			return true
		}
	}

	return false
}
