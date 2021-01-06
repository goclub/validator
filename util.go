package vd

import (
	"errors"
	"fmt"
	"reflect"
)

func EnumValues (v interface{}) (enum []string) {
	rValue := reflect.ValueOf(v)
	for i:=0;i<rValue.NumField();i++ {
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
		panic(errors.New("typjson vd.EnumValues(v)  v values length is zero!"))
	}
	return
}
