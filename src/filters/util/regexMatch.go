package util

import (
	"mana/src/config"
	"regexp"
)

var _log = config.Log()

// RegexpMatchString 正则匹配
// str 匹配参数
// regexpMatch 正则表达式
func RegexpMatchString(str, regexpMatch string) bool {
	matched, err := regexp.MatchString(regexpMatch, str)
	if err != nil {
		_log.Error("GetResourceLinks system exception")
		return false
	}
	// 参数校验
	if !matched {
		_log.Error("Parameters of the abnormal, str: ", str, " , regexpMatch: ", regexpMatch)
		return false
	}
	return true
}
