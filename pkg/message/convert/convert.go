package convert

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/mysql-schema-migration/pkg/message"
)

func init() {
	initConvertDebugMessage()
	initConvertInfoMessage()
	initConvertErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrConvertRSAoSM2     = 402001
	ErrConvertGPConfig    = 402002
	ErrConvertInsightPass = 402003
	ErrConvertTenantPass  = 402004
	ErrConvertPAMConfig   = 402005
)

func initConvertDebugMessage() {

}

func initConvertInfoMessage() {

}

func initConvertErrorMessage() {
	message.Messages[ErrConvertRSAoSM2] = config.NewErrMessage(message.DefaultMessageHeader, ErrConvertRSAoSM2,
		"convert the input from RSA to SM2 failed")
	message.Messages[ErrConvertGPConfig] = config.NewErrMessage(message.DefaultMessageHeader, ErrConvertGPConfig,
		"convert the encrypted data that is in the GP yaml file from rsa to sm2 failed")
	message.Messages[ErrConvertInsightPass] = config.NewErrMessage(message.DefaultMessageHeader, ErrConvertInsightPass,
		"convert the encrypted insight password from rsa to sm2 failed")
	message.Messages[ErrConvertTenantPass] = config.NewErrMessage(message.DefaultMessageHeader, ErrConvertTenantPass,
		"convert the encrypted tenant password from rsa to sm2 failed")
	message.Messages[ErrConvertPAMConfig] = config.NewErrMessage(message.DefaultMessageHeader, ErrConvertPAMConfig,
		"convert the encrypted pam config from rsa to sm2 failed")
}
