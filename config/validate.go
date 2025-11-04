package config

import (
	"path/filepath"

	"github.com/asaskevich/govalidator"
	"github.com/pingcap/errors"
	"github.com/romberli/go-multierror"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/viper"
	"github.com/spf13/cast"

	"github.com/romberli/mysql-schema-migration/pkg/message"
)

// ValidateConfiguration validates if the configuration is valid
func ValidateConfiguration() (err error) {
	merr := &multierror.Error{}

	// validate log section
	err = ValidateLog()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate rsa section
	err = ValidateRSA()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate sm2 section
	err = ValidateSM2()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate input section
	err = ValidateInput()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate convert section
	err = ValidateConvert()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return errors.Trace(merr.ErrorOrNil())
}

// ValidateLog validates if log section is valid.
func ValidateLog() error {
	merr := &multierror.Error{}

	// validate log.level
	logLevel, err := cast.ToStringE(viper.Get(LogLevelKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if !common.ElementInSlice(ValidLogLevels, logLevel) {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogLevel, logLevel))
	}
	// validate log.format
	logFormat, err := cast.ToStringE(viper.Get(LogFormatKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if !common.ElementInSlice(ValidLogFormats, logFormat) {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogFormat, logFormat))
	}

	return merr.ErrorOrNil()
}

// ValidateRSA validates if rsa section is valid.
func ValidateRSA() error {
	merr := &multierror.Error{}

	// validate rsa.private
	_, err := cast.ToStringE(viper.Get(RSAPrivateKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate rsa.public
	_, err = cast.ToStringE(viper.Get(RSAPublicKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	return merr.ErrorOrNil()
}

// ValidateSM2 validates if sm2 section is valid.
func ValidateSM2() error {
	merr := &multierror.Error{}

	// validate sm2.private
	_, err := cast.ToStringE(viper.Get(SM2PrivateKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate sm2.public
	_, err = cast.ToStringE(viper.Get(SM2PublicKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	return merr.ErrorOrNil()
}

// ValidateInout validates if input section is valid.
func ValidateInput() error {
	merr := &multierror.Error{}
	// validate input
	_, err := cast.ToStringE(viper.Get(InputKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	return merr.ErrorOrNil()
}

// ValidateConvert validates if convert section is valid.
func ValidateConvert() error {
	merr := &multierror.Error{}

	// validate convert.yaml.enabled
	_, err := cast.ToBoolE(viper.Get(ConvertYAMLEnabledKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.yaml.path
	path, err := cast.ToStringE(viper.Get(ConvertYAMLPathKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if path != constant.EmptyString {
		isAbs := filepath.IsAbs(path)
		if !isAbs {
			path, err = filepath.Abs(path)
			if err != nil {
				merr = multierror.Append(merr, errors.Trace(err))
			}
		}
		valid, _ := govalidator.IsFilePath(path)
		if !valid {
			merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidPath, path))
		}
	}
	// validate convert.yaml.nested
	_, err = cast.ToStringE(viper.Get(ConvertYAMLNestedPathKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.insight.enabled
	_, err = cast.ToBoolE(viper.Get(ConvertInsightEnabledKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.instance.enabled
	_, err = cast.ToBoolE(viper.Get(ConvertTenantEnabledKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.pam.enabled
	_, err = cast.ToBoolE(viper.Get(ConvertPAMEnabledKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.db.mysql.addr
	_, err = cast.ToStringE(viper.Get(ConvertDBMySQLAddrKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.db.mysql.name
	_, err = cast.ToStringE(viper.Get(ConvertDBMySQLNameKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.db.mysql.user
	_, err = cast.ToStringE(viper.Get(ConvertDBMySQLUserKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate convert.db.mysql.pass
	_, err = cast.ToStringE(viper.Get(ConvertDBMySQLPassKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	return merr.ErrorOrNil()
}
