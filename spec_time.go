package vd

import (
	"time"
)

type TimeSpec struct {
	Name string
	Path string
	AllowZero bool
	BeforeIt time.Time
	AfterIt time.Time
	BeforeOrEqualIt time.Time
	AfterOrEqualIt time.Time
}
func (r *Rule) Time(v time.Time, spec TimeSpec) {
	if r.Fail {return}
	if spec.AllowZero == false && v.IsZero() {
		r.Break(r.Format.TimeNotAllowZero(spec.Name, spec.Path), spec.Path)
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
		r.Break(r.Format.TimeBeforeIt(spec.Name, spec.Path, v, spec.BeforeIt), spec.Path)
	}
	return r.Fail
}

func (spec TimeSpec) CheckAfterIt(v time.Time, r *Rule) (fail bool) {
	if spec.AfterIt.IsZero() {
		return false
	}
	if v.After(spec.AfterIt) == false {
		r.Break(r.Format.TimeAfterIt(spec.Name, spec.Path, v, spec.AfterIt), spec.Path)
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
		r.Break(r.Format.TimeBeforeOrEqualIt(spec.Name, spec.Path, v, spec.BeforeOrEqualIt), spec.Path)
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
		r.Break(r.Format.TimeAfterOrEqualIt(spec.Name, spec.Path, v, spec.AfterOrEqualIt), spec.Path)
	}
	return r.Fail
}