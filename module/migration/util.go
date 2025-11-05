package migration

import (
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/viper"

	"github.com/romberli/mysql-schema-migration/config"
)

func IsTableIncluded(tableName string) bool {
	includeTables := viper.GetStringSlice(config.TableIncludeKey)
	excludeTables := viper.GetStringSlice(config.TableExcludeKey)

	if (len(includeTables) == constant.ZeroInt || common.ElementInSlice(includeTables, tableName)) &&
		!common.ElementInSlice(excludeTables, tableName) {
		return true
	}

	return false
}
