package vd

import (
	"github.com/goclub/conv"
	xtime "github.com/goclub/time"
	"strings"
	"time"
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

	SliceMinLen(name string, len int, minLen int) string
	SliceMaxLen(name string, len int, maxLen int) string
	SliceNotAllowEmpty(name string) string
	SliceUniqueStrings(name string, repeatElement string) string

	TimeRangeDefaultName() (startName string, endTime string)

	TimeNotAllowZero(name string) string
	TimeBeforeIt(name string, value time.Time, beforeIt time.Time) string
	TimeAfterIt(name string, value time.Time, afterIt time.Time) string
	TimeBeforeOrEqualIt(name string, value time.Time, beforeOrEqualIt time.Time) string
	TimeAfterOrEqualIt(name string, value time.Time, afterOrEqualIt time.Time) string

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



func (CNFormat) SliceMinLen(name string, len int, minLen int) string {
	return name + "长度不能小于" + xconv.IntString(minLen)
}
func (CNFormat) SliceMaxLen(name string, len int, maxLen int) string {
	return name + "长度不能大于" + xconv.IntString(maxLen)
}
func (CNFormat) SliceNotAllowEmpty(name string) string {
	return name + "不能为空"
}
func (CNFormat) SliceUniqueStrings(name string, repeatElement string) string {
	return name + "中(" + repeatElement + ")重复"
}

func (CNFormat) TimeNotAllowZero(name string) string {
	return name + "不能为空"
}

func (CNFormat) TimeBeforeIt(name string, value time.Time, beforeIt time.Time) string {
	return name + xtime.FormatChinaTime(value) + "必须在" + xtime.FormatChinaTime(beforeIt) + "之前"
}
func (CNFormat) TimeAfterIt(name string, value time.Time, afterIt time.Time) string {
	return name + xtime.FormatChinaTime(value) + "必须在" + xtime.FormatChinaTime(afterIt) + "之后"
}
func (CNFormat) TimeBeforeOrEqualIt(name string, value time.Time, beforeOrEqualIt time.Time) string {
	return name + xtime.FormatChinaTime(value) + "必须等于" + xtime.FormatChinaTime(beforeOrEqualIt) + "或之前"
}
func (CNFormat) TimeAfterOrEqualIt(name string, value time.Time, afterOrEqualIt time.Time) string {
	return name + xtime.FormatChinaTime(value) + "必须等于" + xtime.FormatChinaTime(afterOrEqualIt) + "或之后"
}

func (CNFormat) TimeRangeDefaultName() (startName string, endTime string){
	return "开始时间", "结束时间"
}