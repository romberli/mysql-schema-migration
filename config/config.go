/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"strings"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/viper"
	"github.com/romberli/log"
)

const (
	TypeFile = "file"
	TypeDB   = "db"
)

var (
	ValidLogLevels  = []string{"debug", "info", "warn", "warning", "error", "fatal"}
	ValidLogFormats = []string{"text", "json"}
	ValidTypes      = []string{TypeFile, TypeDB}
)

// SetDefaultConfig set default configuration, it is the lowest priority
func SetDefaultConfig(baseDir string) {
	// log
	viper.SetDefault(LogLevelKey, log.DefaultLogLevel)
	viper.SetDefault(LogFormatKey, log.DefaultLogFormat)
	// table
	viper.SetDefault(TableIncludeKey, constant.EmptyString)
	viper.SetDefault(TableExcludeKey, constant.EmptyString)
	// source
	viper.SetDefault(SourceTypeKey, DefaultSourceType)
	viper.SetDefault(SourceFileKey, constant.CurrentDir)
	viper.SetDefault(SourceDBAddrKey, constant.DefaultMySQLAddr)
	viper.SetDefault(SourceDBNameKey, constant.EmptyString)
	viper.SetDefault(SourceDBUserKey, constant.DefaultRootUserName)
	viper.SetDefault(SourceDBPassKey, constant.DefaultRootUserPass)
	// target
	viper.SetDefault(TargetTypeKey, DefaultTargetType)
	viper.SetDefault(TargetFileKey, constant.CurrentDir)
	viper.SetDefault(TargetDBAddrKey, constant.DefaultMySQLAddr)
	viper.SetDefault(TargetDBNameKey, constant.EmptyString)
	viper.SetDefault(TargetDBUserKey, constant.DefaultRootUserName)
	viper.SetDefault(TargetDBPassKey, constant.DefaultRootUserPass)
}

// TrimSpaceOfArg trims spaces of given argument
func TrimSpaceOfArg(arg string) string {
	args := strings.SplitN(arg, constant.EqualString, 2)

	switch len(args) {
	case 1:
		return strings.TrimSpace(args[0])
	case 2:
		argName := strings.TrimSpace(args[0])
		argValue := strings.TrimSpace(args[1])
		return fmt.Sprintf("%s=%s", argName, argValue)
	default:
		return arg
	}
}
