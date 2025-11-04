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

	DefaultKeyTypePrivate = "private"
	DefaultKeyTypePublic  = "public"

	DefaultConvertDBMySQLName = "gp"
)

// configuration constant
const (
	// log
	LogLevelKey  = "log.level"
	LogFormatKey = "log.format"
	// rsa
	RSAPrivateKey = "rsa.private"
	RSAPublicKey  = "rsa.public"
	// sm2
	SM2PrivateKey = "sm2.private"
	SM2PublicKey  = "sm2.public"
	// input
	InputKey = "input"
	// convert
	ConvertYAMLEnabledKey    = "convert.yaml.enabled"
	ConvertYAMLPathKey       = "convert.yaml.path"
	ConvertYAMLNestedPathKey = "convert.yaml.nestedPath"
	ConvertInsightEnabledKey = "convert.insightEnabled"
	ConvertTenantEnabledKey  = "convert.tenantEnabled"
	ConvertPAMEnabledKey     = "convert.pamEnabled"
	ConvertDBMySQLAddrKey    = "convert.db.mysql.addr"
	ConvertDBMySQLNameKey    = "convert.db.mysql.name"
	ConvertDBMySQLUserKey    = "convert.db.mysql.user"
	ConvertDBMySQLPassKey    = "convert.db.mysql.pass"
)
