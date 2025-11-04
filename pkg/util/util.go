package util

import (
	"github.com/hokaccha/go-prettyjson"
	"github.com/romberli/go-util/constant"
)

func PrettyJSONString(jsonStr string) (string, error) {
	prettyJSONStr, err := prettyjson.Format([]byte(jsonStr))
	if err != nil {
		return constant.EmptyString, err
	}

	return string(prettyJSONStr), nil
}
