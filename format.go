package vd

import (
	"github.com/goclub/conv"
	xtime "github.com/goclub/time"
	"strconv"
	"strings"
	"time"
)

type Formatter interface {
	Pattern   (name string, path string, value string, pattern []string, failPattern string) string
	BanPattern   (name string, path string, value string, banPattern []string, failBanPattern string) string

	StringNotAllowEmpty(name string, path string) string
	StringRuneLen(name string, path string, value string, length uint64) string
	StringMinRuneLen(name string, path string, value string, length uint64) string
	StringMaxRuneLen(name string, path string, value string, length uint64) string
	StringEnum (name string, path string, value string, enum []string) string

	IntNotAllowEmpty(name string, path string) string
	IntMin(name string, path string, value int, min int) string
	IntMax(name string, path string, value int, max int) string

	FloatMin(name string, path string, value float64, min float64) string
	FloatMax(name string, path string, value float64, max float64) string

	SliceMinLen(name string, path string, len int, minLen int) string
	SliceMaxLen(name string, path string, len int, maxLen int) string
	SliceNotAllowEmpty(name string, path string) string
	SliceUniqueStrings(name string, path string, repeatElement string) string

	TimeRangeDefaultName() (beginName string, endTime string)
	DateRangeDefaultName() (beginName string, endTime string)

	TimeNotAllowZero(name string, path string) string
	TimeBeforeIt(name string, path string, value time.Time, beforeIt time.Time) string
	TimeAfterIt(name string, path string, value time.Time, afterIt time.Time) string
	TimeBeforeOrEqualIt(name string, path string, value time.Time, beforeOrEqualIt time.Time) string
	TimeAfterOrEqualIt(name string, path string, value time.Time, afterOrEqualIt time.Time) string

}
type CNFormat struct {}
func nameOrPath(name string, path string) string {
	if name != "" { return name}
	return path
}
func (CNFormat) StringNotAllowEmpty(name string, path string) string {
	return nameOrPath(name, path)  + "必填"
}
func (f CNFormat) StringRuneLen(name string, path string, value string, length uint64) string {
	return nameOrPath(name, path) + "长度需等于" + strconv.FormatUint(length, 10)
}
func (CNFormat) StringMinRuneLen(name string, path string, value string, length uint64) string {
	return nameOrPath(name, path) + "长度不能小于" + strconv.FormatUint(length, 10)
}
func (CNFormat) StringMaxRuneLen(name string, path string, value string, length uint64) string {
	return nameOrPath(name, path) + "长度不能大于" + strconv.FormatUint(length, 10)
}
func (CNFormat) Pattern(name string, path string, value string, pattern []string, failPattern string) string {
	return nameOrPath(name, path) + "格式错误"
}
func (CNFormat) BanPattern(name string, path string, value string, banPattern []string, failBanPattern string) string {
	return nameOrPath(name, path) + "格式错误"
}
func (CNFormat) StringEnum(name string, path string, value string, enum []string) string {
	return nameOrPath(name, path) + "参数错误，只允许("+ strings.Join(enum, " ") + ")"
}
func (CNFormat) IntNotAllowEmpty(name string, path string) string {
	return nameOrPath(name, path) + "不允许为0"
}
func (CNFormat) IntMin(name string, path string, value int, min int) string {
	return nameOrPath(name, path) + "不能小于" + xconv.IntString(min)
}
func (CNFormat) IntMax(name string, path string, value int, max int) string {
	return nameOrPath(name, path) + "不能大于" + xconv.IntString(max)
}
func (CNFormat) FloatMin(name string, path string, value float64, min float64) string {
	return nameOrPath(name, path) + "不能小于" + xconv.Float64String(min)
}
func (CNFormat) FloatMax(name string, path string, value float64, max float64) string {
	return nameOrPath(name, path) + "不能大于" + xconv.Float64String(max)
}



func (CNFormat) SliceMinLen(name string, path string, len int, minLen int) string {
	return nameOrPath(name, path) + "长度不能小于" + xconv.IntString(minLen)
}
func (CNFormat) SliceMaxLen(name string, path string, len int, maxLen int) string {
	return nameOrPath(name, path) + "长度不能大于" + xconv.IntString(maxLen)
}
func (CNFormat) SliceNotAllowEmpty(name string, path string) string {
	return nameOrPath(name, path) + "不能为空"
}
func (CNFormat) SliceUniqueStrings(name string, path string, repeatElement string) string {
	return nameOrPath(name, path) + "中(" + repeatElement + ")重复"
}

func (CNFormat) TimeNotAllowZero(name string, path string) string {
	return nameOrPath(name, path) + "不能为空"
}

func (CNFormat) TimeBeforeIt(name string, path string, value time.Time, beforeIt time.Time) string {
	return nameOrPath(name, path) + xtime.FormatChinaTime(value) + "必须在" + xtime.FormatChinaTime(beforeIt) + "之前"
}
func (CNFormat) TimeAfterIt(name string, path string, value time.Time, afterIt time.Time) string {
	return nameOrPath(name, path) + xtime.FormatChinaTime(value) + "必须在" + xtime.FormatChinaTime(afterIt) + "之后"
}
func (CNFormat) TimeBeforeOrEqualIt(name string, path string, value time.Time, beforeOrEqualIt time.Time) string {
	return nameOrPath(name, path) + xtime.FormatChinaTime(value) + "必须在" + xtime.FormatChinaTime(beforeOrEqualIt) + "之前，或等于"
}
func (CNFormat) TimeAfterOrEqualIt(name string, path string, value time.Time, afterOrEqualIt time.Time) string {
	return nameOrPath(name, path) + xtime.FormatChinaTime(value) + "必须在" + xtime.FormatChinaTime(afterOrEqualIt) + "之后，或等于"
}

func (CNFormat) TimeRangeDefaultName() (beginName string, endTime string){
	return "开始时间", "结束时间"
}
func (CNFormat) DateRangeDefaultName() (beginName string, endTime string){
	return "开始日期", "结束日期"
}