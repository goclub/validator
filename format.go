package vd

import (
	"github.com/goclub/conv"
	"strings"
)

type Formatter interface {
	Pattern   (name string, value string, pattern []string, failPattern string) string
	BanPattern   (name string, value string, banPattern []string, failBanPattern string) string

	StringNotAllowEmpty(name string) string
	StringMinRuneLen(name string, value string, length int) string
	StringMaxRuneLen(name string, value string, length int) string
	StringEnum (name string, value string, enum []string) string

	IntNotAllowEmpty(name string) string
	IntMin(name string, value int, min int) string
	IntMax(name string, value int, max int) string

	FloatMin(name string, value float64, min float64) string
	FloatMax(name string, value float64, max float64) string

	ArrayMinLen(name string, len int, minLen int) string
	ArrayMaxLen(name string, len int, maxLen int) string
	ArrayNotAllowEmpty(name string) string
	ArrayUniqueStrings(name string, repeatElement string) string
}
type CNFormat struct {}
func (CNFormat) StringNotAllowEmpty(name string) string {
	return name  + "必填"
}
func (CNFormat) StringMinRuneLen(name string, value string, length int) string {
	return name + "长度不能小于" + xconv.IntString(length)
}
func (CNFormat) StringMaxRuneLen(name string, value string, length int) string {
	return name + "长度不能大于" + xconv.IntString(length)
}
func (CNFormat) Pattern(name string, value string, pattern []string, failPattern string) string {
	return name + "格式错误"
}
func (CNFormat) BanPattern(name string, value string, banPattern []string, failBanPattern string) string {
	return name + "格式错误"
}
func (CNFormat) StringEnum(name string, value string, enum []string) string {
	return name + "参数错误，只允许("+ strings.Join(enum, " ") + ")"
}
func (CNFormat) IntNotAllowEmpty(name string) string {
	return name + "不允许为0"
}
func (CNFormat) IntMin(name string, value int, min int) string {
	return name + "不能小于" + xconv.IntString(min)
}
func (CNFormat) IntMax(name string, value int, max int) string {
	return name + "不能大于" + xconv.IntString(max)
}
func (CNFormat) FloatMin(name string, value float64, min float64) string {
	return name + "不能小于" + xconv.Float64String(min)
}
func (CNFormat) FloatMax(name string, value float64, max float64) string {
	return name + "不能大于" + xconv.Float64String(max)
}



func (CNFormat) ArrayMinLen(name string, len int, minLen int) string {
	return name + "长度不能小于" + xconv.IntString(minLen)
}
func (CNFormat) ArrayMaxLen(name string, len int, maxLen int) string {
	return name + "长度不能大于" + xconv.IntString(maxLen)
}
func (CNFormat) ArrayNotAllowEmpty(name string) string {
	return name + "不能为空"
}

func (CNFormat) ArrayUniqueStrings(name string, repeatElement string) string {
	return name + "中(" + repeatElement + ")重复"
}