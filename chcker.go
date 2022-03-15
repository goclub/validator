package vd

import (
	"errors"
	"reflect"
)
type Checker struct {
	Format Formatter
}

type Data interface {
	VD(r *Rule) (err error)
}
type Report struct {
	Fail bool
	Message string
	Path string
}
func (checker Checker) Check(data Data) (report Report, err error) {
	rValue := reflect.ValueOf(data)
	rType := rValue.Type()
	if rType.Kind() == reflect.Ptr {
		err = errors.New("goclub/validator: Check(data) data ("+ rType.Name() + ") must be pointer")
		return
	}
	return checker.reflectCheck(rValue, rType)
}
func (checker Checker) reflectCheck(rValue reflect.Value, rType reflect.Type) (report Report, err error) {
	data := rValue.Interface()
	switch v := data.(type) {
	case Data:
		rule := Rule{
			Format: checker.Format,
		}
		err = v.VD(&rule) ; if err != nil {
		    return
		}
		if rule.Fail {
			report.Fail = true
			report.Message = rule.Message
			report.Path = rule.Path
			return
		}
	}
	for i:=0; i<rType.NumField();i++ {
		rValueItem := rValue.Field(i)
		structField := rType.Field(i)
		switch structField.Type.Kind() {
		case reflect.Struct:
			report, err = checker.reflectCheck(rValueItem, structField.Type) ; if err != nil {
			    return
			}
			if report.Fail {
				return
			}
		}
	}
	return
}

