package vd

import (
	"github.com/hoisie/mustache"
)

// 不实现 AllowEmpty 因为与 MinLen 实现重复。会增加使用者学习成本
type SliceSpec struct {
	Name string
	Path string
	MinLen OptionInt
	MinLenMessage string
	MaxLen OptionInt
	MaxLenMessage string
	UniqueStrings []string
}
func (r *Rule) Slice(length int, spec SliceSpec){
	if r.Fail { return }
	if spec.CheckMinLen(length, r) {return}
	if spec.CheckMaxLen(length, r) {return}
	if spec.CheckUniqueStrings(spec.UniqueStrings, r) {return}
}
type arraySpecRender struct {
	Value interface{}
	SliceSpec
}
func (spec SliceSpec) render (message string, value interface{}) string {
	context := arraySpecRender{
		Value: value,
		SliceSpec: spec,
	}
	return mustache.Render(message, context)
}
func (spec SliceSpec) CheckMinLen(v int, r *Rule) (fail bool) {
	if !spec.MinLen.Valid() {
		return
	}
	min := spec.MinLen.Unwrap()
	pass := v >= min
	if !pass {
		message := r.CreateMessage(spec.MinLenMessage, func() string {
			return r.Format.SliceMinLen(spec.Name, v, min)
		})
		r.Break(spec.render(message, v), spec.Path)
	}
	return r.Fail
}
func (spec SliceSpec) CheckMaxLen(v int, r *Rule) (fail bool) {
	if !spec.MaxLen.Valid() {
		return
	}
	max := spec.MaxLen.Unwrap()
	pass := v <= max
	if !pass {
		message := r.CreateMessage(spec.MaxLenMessage, func() string {
			return r.Format.SliceMaxLen(spec.Name, v, max)
		})
		r.Break(spec.render(message, v), spec.Path)
	}
	return r.Fail
}
func (spec SliceSpec) CheckUniqueStrings(v []string, r *Rule) (fail bool) {
	isRepeat, repeatElement := uniqueStrings(v)
	pass := !isRepeat
	if !pass {
		if !pass {
			message := r.Format.SliceUniqueStrings(spec.Name, repeatElement)
			r.Break(spec.render(message, v), spec.Path)
		}
	}
	return r.Fail
}