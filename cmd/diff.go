package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/cobra"

	"github.com/romberli/mysql-schema-migration/module/migration"
	"github.com/romberli/mysql-schema-migration/pkg/message"

	msgMigration "github.com/romberli/mysql-schema-migration/pkg/message/migration"
)

func init() {
	showCmd.AddCommand(diffCmd)
}

// showCmd represents the show command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "diff command",
	Long:  `show the schema differences`,
	Run: func(cmd *cobra.Command, args []string) {
		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		c := migration.NewController()
		diffList, err := c.GetDiff()
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(msgMigration.ErrMigrationErrGetDiff, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		jsonBytes, err := json.Marshal(diffList)
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrMarshalJSON, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		fmt.Println(common.BytesToString(jsonBytes))
		os.Exit(constant.DefaultNormalExitCode)
	},
}
