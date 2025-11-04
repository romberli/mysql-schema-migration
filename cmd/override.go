package cmd

import (
	"strings"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/viper"

	"github.com/romberli/mysql-schema-migration/config"
	"github.com/romberli/mysql-schema-migration/pkg/message"
)

// OverrideConfigByCLI read configuration from command line interface, it will override the config file configuration
func OverrideConfigByCLI() error {
	// override log
	err := overrideLogByCLI()
	if err != nil {
		return err
	}

	// override table
	overrideTableByCLI()
	// override source
	overrideSourceByCLI()
	// override target
	overrideTargetByCLI()

	// validate configuration
	err = config.ValidateConfiguration()
	if err != nil {
		return message.NewMessage(message.ErrValidateConfig, err)
	}

	return nil
}

// overrideLogByCLI overrides the log section by command line interface
func overrideLogByCLI() error {
	// log.level
	if logLevel != constant.DefaultRandomString {
		logLevel = strings.ToLower(logLevel)
		viper.Set(config.LogLevelKey, logLevel)
	}
	// log.format
	if logFormat != constant.DefaultRandomString {
		logLevel = strings.ToLower(logFormat)
		viper.Set(config.LogFormatKey, logFormat)
	}

	return nil
}

// overrideTableByCLI overrides the table section by command line interface
func overrideTableByCLI() {
	if tableInclude != constant.DefaultRandomString {
		viper.Set(config.TableIncludeKey, tableInclude)
	}
	if tableExclude != constant.DefaultRandomString {
		viper.Set(config.TableExcludeKey, tableExclude)
	}
}

// overrideSourceByCLI overrides the source section by command line interface
func overrideSourceByCLI() {
	if sourceType != constant.DefaultRandomString {
		viper.Set(config.SourceTypeKey, sourceType)
	}
	if sourceFile != constant.DefaultRandomString {
		viper.Set(config.SourceFileKey, sourceFile)
	}
	if sourceDBAddr != constant.DefaultRandomString {
		viper.Set(config.SourceDBAddrKey, sourceDBAddr)
	}
	if sourceDBName != constant.DefaultRandomString {
		viper.Set(config.SourceDBNameKey, sourceDBName)
	}
	if sourceDBUser != constant.DefaultRandomString {
		viper.Set(config.SourceDBUserKey, sourceDBUser)
	}
	if sourceDBPass != constant.DefaultRandomString {
		viper.Set(config.SourceDBPassKey, sourceDBPass)
	}
}

// overrideTargetByCLI overrides the target section by command line interface
func overrideTargetByCLI() {
	if targetType != constant.DefaultRandomString {
		viper.Set(config.TargetTypeKey, targetType)
	}
	if targetFile != constant.DefaultRandomString {
		viper.Set(config.TargetFileKey, targetFile)
	}
	if targetDBAddr != constant.DefaultRandomString {
		viper.Set(config.TargetDBAddrKey, targetDBAddr)
	}
	if targetDBName != constant.DefaultRandomString {
		viper.Set(config.TargetDBNameKey, targetDBName)
	}
	if targetDBUser != constant.DefaultRandomString {
		viper.Set(config.TargetDBUserKey, targetDBUser)
	}
	if targetDBPass != constant.DefaultRandomString {
		viper.Set(config.TargetDBPassKey, targetDBPass)
	}
}
