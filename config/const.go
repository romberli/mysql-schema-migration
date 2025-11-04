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
	"github.com/romberli/go-util/constant"
)

// global constant
const (
	DefaultCommandName = "crypto"
	DefaultBaseDir     = constant.CurrentDir

	DefaultSourceType = "file"
	DefaultTargetType = "file"
)

// configuration constant
const (
	// log
	LogLevelKey  = "log.level"
	LogFormatKey = "log.format"
	// table
	TableIncludeKey = "table.include"
	TableExcludeKey = "table.exclude"
	// source
	SourceTypeKey = "source.type"
	SourceFileKey = "source.file"
	SourceDBAddrKey
	SourceDBNameKey = "source.db.name"
	SourceDBUserKey = "source.db.user"
	SourceDBPassKey = "source.db.pass"
	// target
	TargetTypeKey   = "target.type"
	TargetFileKey   = "target.file"
	TargetDBAddrKey = "target.db.addr"
	TargetDBNameKey = "target.db.name"
	TargetDBUserKey = "target.db.user"
	TargetDBPassKey = "target.db.pass"
)
