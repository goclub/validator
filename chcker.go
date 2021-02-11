package vd

import (
	"errors"
	"log"
	"reflect"
	"runtime/debug"
)
type Checker struct {
	Format Formatter
}

type Data interface {
	VD(r *Rule)
}
type Report struct {
	Fail bool
	Message string
}
func (checker Checker) Check(data Data) (report Report) {
	rValue := reflect.ValueOf(data)
	rType := rValue.Type()
	if rType.Kind() == reflect.Ptr {
		// 为了避免服务中断，采用返回验证失败的锤了
		debug.PrintStack()
		errorMessage := "goclub/validator: Check(data) data ("+ rType.Name() + ") must be pointer"
		log.Print(errors.New(errorMessage))
		return Report{
			Fail: true,
			Message: errorMessage,
		}
	}
	return checker.reflectCheck(rValue, rType)
}
func (checker Checker) reflectCheck(rValue reflect.Value, rType reflect.Type) (report Report) {
	checkMethod := rValue.MethodByName("VD")
	if !checkMethod.IsValid() {
		{
			TjMethod := rValue.MethodByName("Vd")
			if TjMethod.IsValid() {
				checkMethod = TjMethod
				log.Print("typejson: you write error method name" +rType.Name() + ".Vd()")
			}
		}
	}
	if checkMethod.IsValid() {
		rule := Rule{
			Format: checker.Format,
		}
		var rValueList []reflect.Value
		rValueList = append(rValueList, reflect.ValueOf(&rule))
		checkMethod.Call(rValueList)
		if rule.Fail {
			report.Fail = true
			report.Message = rule.Message
			return
		}
	}
	for i:=0; i<rType.NumField();i++ {
		rValueItem := rValue.Field(i)
		structField := rType.Field(i)
		switch structField.Type.Kind() {
		case reflect.Struct:
			report = checker.reflectCheck(rValueItem, structField.Type)
			if report.Fail {
				return
			}
		}
	}
	return
}

