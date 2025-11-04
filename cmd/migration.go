package cmd

import (
	"fmt"
	"os"

	"github.com/romberli/go-util/constant"
	"github.com/spf13/cobra"

	"github.com/romberli/mysql-schema-migration/module/migration"
	"github.com/romberli/mysql-schema-migration/pkg/message"

	msgMigration "github.com/romberli/mysql-schema-migration/pkg/message/migration"
)

func init() {
	showCmd.AddCommand(migrationCmd)
}

// showCmd represents the show command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration command",
	Long:  `show the schema migration sqls`,
	Run: func(cmd *cobra.Command, args []string) {
		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		c := migration.NewController()
		sqlList, err := c.GetSchemaMigrationSQLList()
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(msgMigration.ErrMigrationErrGetMigrationSQLList, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		for _, sql := range sqlList {
			fmt.Println(sql)
		}

		os.Exit(constant.DefaultNormalExitCode)
	},
}
