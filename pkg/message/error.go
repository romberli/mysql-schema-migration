package message

import (
	"github.com/romberli/go-util/config"
)

const (
	ErrPrintHelpInfo           = 400001
	ErrNotValidLogLevel        = 400004
	ErrNotValidLogFormat       = 400005
	ErrValidateConfig          = 400014
	ErrInitDefaultConfig       = 400015
	ErrOverrideCommandLineArgs = 400017
	ErrInitLogger              = 400019
	ErrBaseDir                 = 400020
	ErrInitConfig              = 400021
	ErrNotValidPath            = 400022
	ErrEmptyPath               = 400023
	ErrNotValidType            = 400024
	ErrEmptyDBAddr             = 400025
	ErrEmptyDBName             = 400026
	ErrEmptyDBUser             = 400027
	ErrEmptyDBPass             = 400028
	ErrMarshalJSON             = 400029
)

func initErrorMessage() {
	Messages[ErrPrintHelpInfo] = config.NewErrMessage(DefaultMessageHeader, ErrPrintHelpInfo, "got message when printing help information")
	Messages[ErrNotValidLogLevel] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogLevel, "log level must be one of [debug, info, warn, message, fatal], %s is not valid")
	Messages[ErrNotValidLogFormat] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFormat, "log level must be either text or json, %s is not valid")
	Messages[ErrValidateConfig] = config.NewErrMessage(DefaultMessageHeader, ErrValidateConfig, "validate config failed")
	Messages[ErrInitDefaultConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitDefaultConfig, "init default configuration failed")
	Messages[ErrOverrideCommandLineArgs] = config.NewErrMessage(DefaultMessageHeader, ErrOverrideCommandLineArgs, "override command line arguments failed")
	Messages[ErrInitLogger] = config.NewErrMessage(DefaultMessageHeader, ErrInitLogger, "initialize logger failed")
	Messages[ErrBaseDir] = config.NewErrMessage(DefaultMessageHeader, ErrBaseDir, "get base dir of %s failed")
	Messages[ErrInitConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitConfig, "init config failed")
	Messages[ErrNotValidPath] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidPath, "path must be either unix or windows path format, %s is not valid")
	Messages[ErrEmptyPath] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyPath, "when type is file, path should not be empty")
	Messages[ErrNotValidType] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidType, "type must be either file or db, %s is not valid")
	Messages[ErrEmptyDBAddr] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyDBAddr, "when type is db, db address should not be empty")
	Messages[ErrEmptyDBName] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyDBName, "when type is db, db name should not be empty")
	Messages[ErrEmptyDBUser] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyDBUser, "when type is db, db user should not be empty")
	Messages[ErrEmptyDBPass] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyDBPass, "when type is db, db pass should not be empty")
	Messages[ErrMarshalJSON] = config.NewErrMessage(DefaultMessageHeader, ErrMarshalJSON, "marshal json failed")
}
