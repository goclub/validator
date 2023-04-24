package vdg

import (
	vd "github.com/goclub/validator"
	"github.com/hoisie/mustache"
)

// SliceSpec 不实现 AllowEmpty 因为与 MinLen 实现重复。会增加使用者学习成本
type SliceSpec struct {
	Name          string
	Path          string
	MinLen        vd.OptionInt
	MinLenMessage string
	MaxLen        vd.OptionInt
	MaxLenMessage string
	Unique        bool
}

func Slice[T comparable](r *vd.Rule, slice []T, spec SliceSpec, elementCheck ...func(elem T, r *vd.Rule, index int) (err error)) {
	if r.Fail {
		return
	}
	length := len(slice)
	if spec.CheckMinLen(length, r) {
		return
	}
	if spec.CheckMaxLen(length, r) {
		return
	}
	if sliceSpecCheckUnique(spec, slice, r) {
		return
	}
	for _, check := range elementCheck {
		for i, v := range slice {
			if err := check(v, r, i); err != nil {
				r.Error(err)
				break
			}
		}
	}
}

type arraySpecRender struct {
	Value interface{}
	SliceSpec
}

func (spec SliceSpec) render(message string, value interface{}) string {
	context := arraySpecRender{
		Value:     value,
		SliceSpec: spec,
	}
	return mustache.Render(message, context)
}
func (spec SliceSpec) CheckMinLen(v int, r *vd.Rule) (fail bool) {
	if !spec.MinLen.Valid() {
		return
	}
	min := spec.MinLen.Unwrap()
	pass := v >= min
	if !pass {
		message := r.CreateMessage(spec.MinLenMessage, func() string {
			return r.Format.SliceMinLen(spec.Name, spec.Path, v, min)
		})
		r.Break(spec.render(message, v), spec.Path)
	}
	return r.Fail
}
func (spec SliceSpec) CheckMaxLen(v int, r *vd.Rule) (fail bool) {
	if !spec.MaxLen.Valid() {
		return
	}
	max := spec.MaxLen.Unwrap()
	pass := v <= max
	if !pass {
		message := r.CreateMessage(spec.MaxLenMessage, func() string {
			return r.Format.SliceMaxLen(spec.Name, spec.Path, v, max)
		})
		r.Break(spec.render(message, v), spec.Path)
	}
	return r.Fail
}
func sliceSpecCheckUnique[T comparable](spec SliceSpec, slice []T, r *vd.Rule) (fail bool) {
	isRepeat := false
	seen := make(map[T]struct{})
	var repeatElement interface{}
	for _, v := range slice {
		if _, ok := seen[v]; ok {
			isRepeat = true
			repeatElement = v
			break
		}
		seen[v] = struct{}{}
	}
	pass := !isRepeat
	if !pass {
		if !pass {
			message := r.Format.SliceUnique(spec.Name, spec.Path, repeatElement)
			r.Break(spec.render(message, slice), spec.Path)
		}
	}
	return r.Fail
}
