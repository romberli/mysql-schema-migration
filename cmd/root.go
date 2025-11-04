/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/viper"
	"github.com/romberli/log"
	"github.com/spf13/cobra"

	"github.com/romberli/mysql-schema-migration/config"
	"github.com/romberli/mysql-schema-migration/pkg/message"
)

const (
	defaultConfigFileType = "yaml"
)

var (
	// config
	baseDir string
	cfgFile string

	// log
	logLevel  string
	logFormat string
	// rsa
	rsaPrivate string
	rsaPublic  string
	// sm2
	sm2Private string
	sm2Public  string
	// input
	input string
	// convert
	convertYAMLEnabledStr    string
	convertYAMLPath          string
	convertYAMLNestedPath    string
	convertInsightEnabledStr string
	convertTenantEnabledStr  string
	convertPAMEnabledStr     string
	convertDBMySQLAddr       string
	convertDBMySQLName       string
	convertDBMySQLUser       string
	convertDBMySQLPass       string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-crypto",
	Short: "go-crypto",
	Long:  `go-crypto is a encryption and decryption tool written in go.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if no subcommand is set, it will print help information.
		if len(args) == 0 {
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(fmt.Sprintf(constant.LogWithStackString, errors.Trace(err)))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func init() {
	// set usage template
	rootCmd.SetUsageTemplate(UsageTemplateWithoutDefault())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", constant.DefaultRandomString, "config file path")
	// log
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", constant.DefaultRandomString, fmt.Sprintf("specify the log level(default: %s)", log.DefaultLogLevel))
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", constant.DefaultRandomString, fmt.Sprintf("specify the log format(default: %s)", log.DefaultLogFormat))
	// rsa
	rootCmd.PersistentFlags().StringVar(&rsaPrivate, "rsa-private", constant.DefaultRandomString, fmt.Sprintf("specify rsa private key"))
	rootCmd.PersistentFlags().StringVar(&rsaPublic, "rsa-public", constant.DefaultRandomString, fmt.Sprintf("specify rsa public key"))
	rootCmd.PersistentFlags().StringVar(&sm2Private, "sm2-private", constant.DefaultRandomString, fmt.Sprintf("specify sm2 private key"))
	// sm2
	rootCmd.PersistentFlags().StringVar(&sm2Public, "sm2-public", constant.DefaultRandomString, fmt.Sprintf("specify sm2 public key"))
	// input
	rootCmd.PersistentFlags().StringVar(&input, "input", constant.DefaultRandomString, "specify the input string")
	// convert
	rootCmd.PersistentFlags().StringVar(&convertYAMLEnabledStr, "convert-yaml-enabled", constant.DefaultRandomString, "specify whether to convert yaml")
	rootCmd.PersistentFlags().StringVar(&convertYAMLPath, "convert-yaml-path", constant.DefaultRandomString, "specify the path of yaml file")
	rootCmd.PersistentFlags().StringVar(&convertYAMLNestedPath, "convert-yaml-nested-path", constant.DefaultRandomString, "specify the nested path")
	rootCmd.PersistentFlags().StringVar(&convertInsightEnabledStr, "convert-insight-enabled", constant.DefaultRandomString, "specify whether to convert insight")
	rootCmd.PersistentFlags().StringVar(&convertTenantEnabledStr, "convert-tenant-enabled", constant.DefaultRandomString, "specify whether to convert tenant")
	rootCmd.PersistentFlags().StringVar(&convertPAMEnabledStr, "convert-pam-enabled", constant.DefaultRandomString, "specify whether to convert pam")
	rootCmd.PersistentFlags().StringVar(&convertDBMySQLAddr, "convert-db-mysql-addr", constant.DefaultRandomString, "specify the mysql address")
	rootCmd.PersistentFlags().StringVar(&convertDBMySQLName, "convert-db-mysql-name", constant.DefaultRandomString, "specify the mysql name")
	rootCmd.PersistentFlags().StringVar(&convertDBMySQLUser, "convert-db-mysql-user", constant.DefaultRandomString, "specify the mysql user")
	rootCmd.PersistentFlags().StringVar(&convertDBMySQLPass, "convert-db-mysql-pass", constant.DefaultRandomString, "specify the mysql password")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	var err error

	// init default config
	err = initDefaultConfig()
	if err != nil {
		return message.NewMessage(message.ErrInitDefaultConfig, err.Error())
	}

	// read config with config file
	err = ReadConfigFile()
	if err != nil {
		return message.NewMessage(message.ErrInitDefaultConfig, err)
	}

	// override config with command line arguments
	err = OverrideConfigByCLI()
	if err != nil {
		return message.NewMessage(message.ErrOverrideCommandLineArgs, err)
	}

	// init log
	level := viper.GetString(config.LogLevelKey)
	format := viper.GetString(config.LogFormatKey)

	logger, zapProperties, err := log.InitStdoutLogger(level, format)
	if err != nil {
		return message.NewMessage(message.ErrInitLogger, err)
	}

	log.ReplaceGlobals(logger, zapProperties)
	log.SetDisableDoubleQuotes(true)
	log.SetDisableEscape(true)

	return nil
}

// initDefaultConfig initiate default configuration
func initDefaultConfig() (err error) {
	// get base dir
	baseDir, err = filepath.Abs(config.DefaultBaseDir)
	if err != nil {
		return message.NewMessage(message.ErrBaseDir, errors.Trace(err), config.DefaultCommandName)
	}
	// set default config value
	config.SetDefaultConfig(baseDir)
	err = config.ValidateConfiguration()
	if err != nil {
		return err
	}

	return nil
}

// ReadConfigFile read configuration from config file, it will override the init configuration
func ReadConfigFile() (err error) {
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType(defaultConfigFileType)
		err = viper.ReadInConfig()
		if err != nil {
			return errors.Trace(err)
		}
		err = config.ValidateConfiguration()
		if err != nil {
			return message.NewMessage(message.ErrValidateConfig, err)
		}
	}

	return nil
}

// UsageTemplateWithoutDefault returns a usage template which does not contain default part
func UsageTemplateWithoutDefault() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
