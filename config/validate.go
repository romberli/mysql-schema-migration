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
	// validate table
	err = ValidateTable()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate source
	err = ValidateSource()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate target
	err = ValidateTarget()
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

// ValidateTable validates if table section is valid.
func ValidateTable() error {
	merr := &multierror.Error{}

	// validate table.include
	_, err := cast.ToStringSliceE(viper.Get(TableIncludeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	// validate table.exclude
	_, err = cast.ToStringSliceE(viper.Get(TableExcludeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	return merr.ErrorOrNil()
}

// ValidateSource validates if source section is valid.
func ValidateSource() error {
	merr := &multierror.Error{}
	// validate source.type
	sourceType, err := cast.ToStringE(viper.Get(SourceTypeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if !common.ElementInSlice(ValidTypes, sourceType) {
			merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidType, sourceType))
		}
	}
	// validate source.file
	path, err := cast.ToStringE(viper.Get(SourceFileKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if sourceType == TypeFile {
			if path == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyPath))
			} else {
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
		}
	}
	// validate source.db.addr
	sourceDBAddr, err := cast.ToStringE(viper.Get(SourceDBAddrKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if sourceType == TypeDB {
			if sourceDBAddr == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBAddr))
			}
		}
	}
	// validate source.db.name
	sourceDBName, err := cast.ToStringE(viper.Get(SourceDBNameKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if sourceType == TypeDB {
			if sourceDBName == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBName))
			}
		}
	}
	// validate source.db.user
	sourceDBUser, err := cast.ToStringE(viper.Get(SourceDBUserKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if sourceType == TypeDB {
			if sourceDBUser == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBUser))
			}
		}
	}
	// validate source.db.pass
	sourceDBPass, err := cast.ToStringE(viper.Get(SourceDBPassKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if sourceType == TypeDB {
			if sourceDBPass == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBPass))
			}
		}
	}

	return merr.ErrorOrNil()
}

// ValidateTarget validates if target section is valid.
func ValidateTarget() error {
	merr := &multierror.Error{}
	// validate target.type
	targetType, err := cast.ToStringE(viper.Get(TargetTypeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if !common.ElementInSlice(ValidTypes, targetType) {
			merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidType, targetType))
		}
	}
	// validate target.file
	path, err := cast.ToStringE(viper.Get(TargetFileKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if targetType == TypeFile {
			if path == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyPath))
			} else {
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
		}
	}
	// validate target.db.addr
	targetDBAddr, err := cast.ToStringE(viper.Get(TargetDBAddrKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if targetType == TypeDB {
			if targetDBAddr == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBAddr))
			}
		}
	}
	// validate target.db.name
	targetDBName, err := cast.ToStringE(viper.Get(TargetDBNameKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if targetType == TypeDB {
			if targetDBName == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBName))
			}
		}
	}
	// validate target.db.user
	targetDBUser, err := cast.ToStringE(viper.Get(TargetDBUserKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if targetType == TypeDB {
			if targetDBUser == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBUser))
			}
		}
	}
	// validate target.db.pass
	targetDBPass, err := cast.ToStringE(viper.Get(TargetDBPassKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		if targetType == TypeDB {
			if targetDBPass == constant.EmptyString {
				merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyDBPass))
			}
		}
	}

	return merr.ErrorOrNil()
}
