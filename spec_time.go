package vd

import (
	"time"
)

type TimeSpec struct {
	Name string
	AllowZero bool
	BeforeIt time.Time
	AfterIt time.Time
	BeforeOrEqualIt time.Time
	AfterOrEqualIt time.Time
}
type TimeRange struct {
	StartName string
	StartTime time.Time
	EndName string
	EndTime time.Time
}
// 等同于同时配置
// r.Time(v.StartTime, vd.TimeSpec{BeforeOrEqualIt: v.EndTime})
// 和 r.Time(v.EndTime, vd.TimeSpec{AfterOrEqualIt: v.StartTime})
func (r *Rule) TimeRange(data TimeRange) {
	r.coreTimeRange(false, data)
}
// 等同于同时配置
// r.Time(v.StartTime, vd.TimeSpec{BeforeOrEqualIt: v.EndTime, AllowZero: true})
// 和 r.Time(v.EndTime, vd.TimeSpec{AfterOrEqualIt: v.StartTime, AllowZero: true})
func (r *Rule) TimeRangeAllowZero(data TimeRange) {
	r.coreTimeRange(true, data)
}
func (r *Rule) coreTimeRange(allowZero bool, data TimeRange) {
	defaultStartTimeName, defaultEndTimeName :=  r.Format.TimeRangeDefaultName()
	if len(data.StartName) == 0 {
		data.StartName = defaultStartTimeName
	}
	if len(data.EndName) == 0 {
		data.EndName = defaultEndTimeName
	}
	r.Time(data.StartTime, TimeSpec{
		Name: data.StartName,
		AllowZero: allowZero,
		BeforeOrEqualIt: data.EndTime,
	})
	r.Time(data.EndTime, TimeSpec{
		Name: data.EndName,
		AllowZero: allowZero,
		AfterOrEqualIt: data.EndTime,
	})
}
func (r *Rule) Time(v time.Time, spec TimeSpec) {
	if r.Fail {return}
	if spec.AllowZero == false && v.IsZero() {
		r.Break(r.Format.TimeNotAllowZero(spec.Name))
		return
	}
	if spec.CheckBeforeIt(v, r) { return }
	if spec.CheckAfterIt(v, r) { return }
	if spec.CheckBeforeOrEqualIt(v, r) {return}
	if spec.CheckAfterOrEqualIt(v, r) {return}
}
func (spec TimeSpec) CheckBeforeIt(v time.Time, r *Rule) (fail bool) {
	if spec.BeforeIt.IsZero() {
		return false
	}
	if v.Before(spec.BeforeIt) == false {
		r.Break(r.Format.TimeBeforeIt(spec.Name, v, spec.BeforeIt))
	}
	return r.Fail
}

func (spec TimeSpec) CheckAfterIt(v time.Time, r *Rule) (fail bool) {
	if spec.AfterIt.IsZero() {
		return false
	}
	if v.After(spec.AfterIt) == false {
		r.Break(r.Format.TimeAfterIt(spec.Name, v, spec.AfterIt))
	}
	return r.Fail
}

func (spec TimeSpec) CheckBeforeOrEqualIt(v time.Time, r *Rule) (fail bool) {
	if spec.BeforeOrEqualIt.IsZero() {
		return false
	}
	if spec.BeforeOrEqualIt.Equal(v) {
		return false
	}
	if v.Before(spec.BeforeOrEqualIt) == false {
		r.Break(r.Format.TimeBeforeOrEqualIt(spec.Name, v, spec.BeforeOrEqualIt))
	}
	return r.Fail
}

func (spec TimeSpec) CheckAfterOrEqualIt(v time.Time, r *Rule) (fail bool) {
	if spec.AfterOrEqualIt.IsZero() {
		return false
	}
	if spec.AfterOrEqualIt.Equal(v) {
		return false
	}
	if v.After(spec.AfterOrEqualIt) == false {
		r.Break(r.Format.TimeAfterOrEqualIt(spec.Name, v, spec.AfterOrEqualIt))
	}
	return r.Fail
}