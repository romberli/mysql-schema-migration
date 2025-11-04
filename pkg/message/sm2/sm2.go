package rsa

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/mysql-schema-migration/pkg/message"
)

func init() {
	initSM2DebugMessage()
	initSM2InfoMessage()
	initSM2ErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrSM2Encrypt = 401001
	ErrSM2Decrypt = 401002
)

func initSM2DebugMessage() {

}

func initSM2InfoMessage() {

}

func initSM2ErrorMessage() {
	message.Messages[ErrSM2Encrypt] = config.NewErrMessage(message.DefaultMessageHeader, ErrSM2Encrypt,
		"sm2: encrypt failed. privateKey: %s, publicKey: %s, input: %s")
	message.Messages[ErrSM2Decrypt] = config.NewErrMessage(message.DefaultMessageHeader, ErrSM2Decrypt,
		"sm2: decrypt failed. privateKey: %s, publicKey: %s, input: %s")
}
