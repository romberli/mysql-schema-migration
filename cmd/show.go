package cmd

import (
	"fmt"
	"os"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/cobra"

	"github.com/romberli/mysql-schema-migration/pkg/message"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show command",
	Long:  `show the schema differences or migration sqls.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == constant.ZeroInt {
			err := cmd.Help()
			if err != nil {
				fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrPrintHelpInfo, errors.Trace(err))))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			os.Exit(constant.DefaultNormalExitCode)
		}

		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
	},
}
