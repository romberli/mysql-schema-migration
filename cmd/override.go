package cmd

import (
	"strings"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/viper"
	"github.com/spf13/cast"

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

	// override rsa
	overrideRSAByCLI()
	// override sm2
	overrideSM2ByCLI()
	// override input
	overrideInputByCLI()
	// override convert
	err = overrideConvertByCLI()
	if err != nil {
		return err
	}

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

// overrideRSAByCLI overrides the rsa section by command line interface
func overrideRSAByCLI() {
	if rsaPrivate != constant.DefaultRandomString {
		viper.Set(config.RSAPrivateKey, rsaPrivate)
	}
	if rsaPublic != constant.DefaultRandomString {
		viper.Set(config.RSAPublicKey, rsaPublic)
	}
}

// overrideSM2ByCLI overrides the sm2 section by command line interface
func overrideSM2ByCLI() {
	if sm2Private != constant.DefaultRandomString {
		viper.Set(config.SM2PrivateKey, sm2Private)
	}
	if sm2Public != constant.DefaultRandomString {
		viper.Set(config.SM2PublicKey, sm2Public)
	}
}

// overrideInputByCLI overrides the input section by command line interface
func overrideInputByCLI() {
	if input != constant.DefaultRandomString {
		viper.Set(config.InputKey, input)
	}
}

// overrideConvertByCLI overrides the convert section by command line interface
func overrideConvertByCLI() error {
	if convertYAMLEnabledStr != constant.DefaultRandomString {
		convertYAMLEnabled, err := cast.ToBoolE(convertYAMLEnabledStr)
		if err != nil {
			return err
		}
		viper.Set(config.ConvertYAMLEnabledKey, convertYAMLEnabled)
	}
	if convertYAMLPath != constant.DefaultRandomString {
		viper.Set(config.ConvertYAMLPathKey, convertYAMLPath)
	}
	if convertYAMLNestedPath != constant.DefaultRandomString {
		viper.Set(config.ConvertYAMLNestedPathKey, convertYAMLNestedPath)
	}
	if convertInsightEnabledStr != constant.DefaultRandomString {
		convertInsightEnabled, err := cast.ToBoolE(convertInsightEnabledStr)
		if err != nil {
			return err
		}
		viper.Set(config.ConvertInsightEnabledKey, convertInsightEnabled)
	}
	if convertTenantEnabledStr != constant.DefaultRandomString {
		convertTenantEnabled, err := cast.ToBoolE(convertTenantEnabledStr)
		if err != nil {
			return err
		}
		viper.Set(config.ConvertTenantEnabledKey, convertTenantEnabled)
	}
	if convertPAMEnabledStr != constant.DefaultRandomString {
		convertPAMEnabled, err := cast.ToBoolE(convertPAMEnabledStr)
		if err != nil {
			return err
		}
		viper.Set(config.ConvertPAMEnabledKey, convertPAMEnabled)
	}
	if convertDBMySQLAddr != constant.DefaultRandomString {
		viper.Set(config.ConvertDBMySQLAddrKey, convertDBMySQLAddr)
	}
	if convertDBMySQLName != constant.DefaultRandomString {
		viper.Set(config.ConvertDBMySQLNameKey, convertDBMySQLName)
	}
	if convertDBMySQLUser != constant.DefaultRandomString {
		viper.Set(config.ConvertDBMySQLUserKey, convertDBMySQLUser)
	}
	if convertDBMySQLPass != constant.DefaultRandomString {
		viper.Set(config.ConvertDBMySQLPassKey, convertDBMySQLPass)
	}

	return nil
}
