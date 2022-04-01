package vd

import (
	xconv "github.com/goclub/conv"
	"github.com/hoisie/mustache"

)

type FloatSpec struct {
	Name string
	Path string
	// AllowZero bool // 暂时取消 AllowZero，目的是降低使用者学习成本，观察一段时间后再决定是否完全去掉 (2020年08月07日 by @nimoc)
	Min OptionFloat
	MinMessage string
	Max OptionFloat
	MaxMessage string
	Pattern []string
	BanPattern []string
	PatternMessage string
}
type FloatSpecRender struct {
	Value interface{}
	FloatSpec
}
func (spec FloatSpec) render (message string, value interface{}) string {
	context := FloatSpecRender{
		Value: value,
		FloatSpec: spec,
	}
	return mustache.Render(message, context)
}
func (r *Rule) Float(v float64, spec FloatSpec) {
	if r.Fail {return}
	// if v == 0 && !spec.AllowZero {
	// 	r.Break(r.Format.FloatNotAllowEmpty(spec.Name))
	// 	return
	// }
	if spec.CheckMin(v, r) { return }
	if spec.CheckMax(v ,r) { return }
	if spec.CheckPattern(v, r) {return}
	if spec.CheckBanPattern(v, r) {return}
	return
}
func (spec FloatSpec) CheckMin(v float64, r *Rule) (fail bool) {
	if !spec.Min.Valid() {
		return
	}
	min := spec.Min.Unwrap()
	pass := v >= min
	if !pass {
		message := r.CreateMessage(spec.MinMessage, func() string {
			return r.Format.FloatMin(spec.Name, spec.Path, v, min)
		})
		r.Break(spec.render(message, v), spec.Path)
	}
	return
}
func (spec FloatSpec) CheckMax(v float64, r *Rule) (fail bool) {
	if !spec.Max.Valid() {
		return
	}
	max := spec.Max.Unwrap()
	pass := v <= max
	if !pass {
		message := r.CreateMessage(spec.MaxMessage, func() string {
			return r.Format.FloatMax(spec.Name, spec.Path, v, max)
		})
		r.Break(spec.render(message, v), spec.Path)
	}
	return
}
func (spec FloatSpec) CheckPattern(v float64, r *Rule) (fail bool) {
	return checkPattern(patternData{
		Pattern:        spec.Pattern,
		PatternMessage: spec.PatternMessage,
		Name:           spec.Name,
		Path:           spec.Path,
	}, spec.render, xconv.Float64String(v), r, spec.Path)
}

func (spec FloatSpec) CheckBanPattern(v float64, r *Rule) (fail bool) {
	return checkBanPattern(banPatternData{
		BanPattern:        spec.BanPattern,
		PatternMessage: spec.PatternMessage,
		Name:           spec.Name,
		Path:           spec.Path,
	}, spec.render, xconv.Float64String(v), r, spec.Path)
}
