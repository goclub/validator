package vd

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Checker struct {
	Format Formatter
}

type Data interface {
	VD(r *Rule) (err error)
}
type Report struct {
	Fail    bool
	Message string
	Path    string
}

func (checker Checker) Check(data Data) (report Report, err error) {
	rValue := reflect.ValueOf(data)
	rType := rValue.Type()
	if rType.Kind() == reflect.Ptr {
		err = errors.New("goclub/validator: Check(data) data (" + rType.Name() + ") must be pointer")
		return
	}
	return checker.reflectCheck(rValue, rType, []string{})
}
func (checker Checker) reflectCheck(rValue reflect.Value, rType reflect.Type, path []string) (report Report, err error) {
	data := rValue.Interface()
	switch v := data.(type) {
	case Data:
		rule := Rule{
			Format: checker.Format,
			Path:   path,
		}
		err = v.VD(&rule)
		if err != nil {
			return
		}
		if rule.Fail {
			report.Fail = true
			report.Message = rule.Message
			report.Path = strings.Join(rule.Path, ".")
			return
		}
	}
	for i := 0; i < rType.NumField(); i++ {
		rValueItem := rValue.Field(i)
		structField := rType.Field(i)
		var oldPath []string
		copy(path, oldPath)
		switch structField.Type.Kind() {
		case reflect.Slice:
			sliceLen := rValueItem.Len()
			for i := 0; i < sliceLen; i++ {
				sliceItem := rValueItem.Index(i)
				sliceItemType := sliceItem.Type()
				if sliceItemType.Kind() == reflect.Struct {
					vdpath, hasVdpath := structField.Tag.Lookup("json")
					if hasVdpath == false {
						vdpath = strings.ToLower(structField.Name)
					}
					path = append(path, vdpath)
					path = append(path, strconv.FormatInt(int64(i), 10))
					report, err = checker.reflectCheck(sliceItem, sliceItemType, path)
					if err != nil {
						return
					}
					if report.Fail {
						return
					}
					path = path[0 : len(path)-2]
				}
			}

		case reflect.Struct:
			vdpath, hasVdpath := structField.Tag.Lookup("json")
			if hasVdpath {
				path = append(path, vdpath)
			} else {
				vdpath = strings.ToLower(structField.Name)
			}
			report, err = checker.reflectCheck(rValueItem, structField.Type, path)
			if err != nil {
				return
			}
			if report.Fail {
				return
			}
		}
		path = oldPath
	}
	return
}
