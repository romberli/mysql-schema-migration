package rsa

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/mysql-schema-migration/pkg/message"
)

func init() {
	initRSADebugMessage()
	initRSAInfoMessage()
	initRSAErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrRSAEncrypt = 402001
	ErrRSADecrypt = 402002
)

func initRSADebugMessage() {

}

func initRSAInfoMessage() {

}

func initRSAErrorMessage() {
	message.Messages[ErrRSAEncrypt] = config.NewErrMessage(message.DefaultMessageHeader, ErrRSAEncrypt,
		"rsa: encrypt failed. privateKey: %s, publicKey: %s, input: %s")
	message.Messages[ErrRSADecrypt] = config.NewErrMessage(message.DefaultMessageHeader, ErrRSADecrypt,
		"rsa: decrypt failed. privateKey: %s, publicKey: %s, input: %s")
}
