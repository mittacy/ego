package utils

import "strings"

// StringFirstUpper 字符串首字母大写
// @param s
// @return string
func StringFirstUpper(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

// StringFirstLower 字符串首字母小写
// @param s
// @return string
func StringFirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
