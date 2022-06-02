package vd

import (
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"strconv"
)

func EnumValues(v interface{}) (enum []string) {
	rValue := reflect.ValueOf(v)
	for i := 0; i < rValue.NumField(); i++ {
		itemValue := rValue.Field(i)
		var value string
		if itemValue.Type().Kind() == reflect.String {
			value = itemValue.String()
		} else {
			value = fmt.Sprintf("%v", itemValue.Interface())
		}
		enum = append(enum, value)
	}
	if len(enum) == 0 {
		// 此处传递 err 是没有用的，因为调用方无法处理，只能修改代码重新发布。
		// 如果 panic 会导致服务中断，所以才去返回错误信息在 enum 让验证不会通过，这样好过项目 panic
		errorMessage := "goclub/validator: vd.EnumValues(v)  v (" + rValue.Type().Name() + ")values length is zero!"
		log.Print(errorMessage)
		debug.PrintStack()
		return []string{errorMessage}
	}
	return
}
func PathIndex(path string, index int) string {
	return path + "." + strconv.FormatInt(int64(index), 10)

}