package stringutils

import "strings"

// IsBlank 判断字符串是否为空
func IsBlank(s string) bool {
    return len(strings.TrimSpace(s)) == 0
}