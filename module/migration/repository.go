package migration

import (
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/mysql"
)

const (
	ShowCreateTableSQLTemplate = `SHOW CREATE TABLE %s`
)

type Repository struct {
	Conn *mysql.Conn
}

func NewRepository(conn *mysql.Conn) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Close() error {
	return r.Conn.Close()
}

func (r *Repository) Execute(sql string, args ...interface{}) (middleware.Result, error) {
	return r.Conn.Execute(sql, args...)
}

func (r *Repository) GetTableNames() ([]string, error) {
	sql := `SHOW TABLES`

	result, err := r.Execute(sql)
	if err != nil {
		return nil, err
	}

	tableNames := make([]string, result.RowNumber())
	for i := constant.ZeroInt; i < result.RowNumber(); i++ {
		tableName, err := result.GetString(i, constant.ZeroInt)
		if err != nil {
			return nil, err
		}
		tableNames[i] = tableName
	}

	return tableNames, nil
}

func (r *Repository) GetCreateTableSQL(tableName string) (string, error) {
	sql := fmt.Sprintf(ShowCreateTableSQLTemplate, tableName)

	result, err := r.Execute(sql)
	if err != nil {
		return constant.EmptyString, err
	}

	createSQL, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	if err != nil {
		return constant.EmptyString, err
	}

	return createSQL + constant.SemicolonString, nil
}
