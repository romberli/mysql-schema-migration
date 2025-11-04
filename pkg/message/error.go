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
}
