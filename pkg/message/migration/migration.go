package rsa

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/mysql-schema-migration/pkg/message"
)

func init() {
	initMigrationDebugMessage()
	initMigrationInfoMessage()
	initMigrationErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrMigrationErrGetDiff             = 401001
	ErrMigrationErrGetMigrationSQLList = 401002
)

func initMigrationDebugMessage() {

}

func initMigrationInfoMessage() {

}

func initMigrationErrorMessage() {
	message.Messages[ErrMigrationErrGetDiff] = config.NewErrMessage(message.DefaultMessageHeader, ErrMigrationErrGetDiff,
		"migration: show table differences failed.")
	message.Messages[ErrMigrationErrGetMigrationSQLList] = config.NewErrMessage(message.DefaultMessageHeader, ErrMigrationErrGetMigrationSQLList,
		"migration: get migration SQL list failed.")
}
