package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBlank(t *testing.T) {
	// 初始化 assert 对象
	ast := assert.New(t)

	// 直接断言
	ast.True(IsBlank(""), "空字符串应该是 true")
	ast.True(IsBlank("   "), "只包含空格的字符串应该是 true")
	ast.True(IsBlank("\n\t"), "包含换行符的字符串应该是 true")
	
	ast.False(IsBlank("a"), "单个字符不应该是 blank")
	ast.False(IsBlank(" hello "), "普通字符串不应该是 blank")
}