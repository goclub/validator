package vd

import (
	"testing"
)

type SliceMinLen struct {
	Skills []string
}
func (v SliceMinLen) VD(r *Rule)  (err error) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MinLen: Int(2),
	})
	return nil
}
func TestSliceMinLen(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, SliceMinLen{Skills: []string{}}, Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLen{Skills: []string{"c"}}, Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLen{Skills: []string{"c","d"}}, Report{
		Fail:    false,
		Message: "",
	})
}

type SliceMinLenMessage struct {
	Skills []string
}
func (v SliceMinLenMessage) VD(r *Rule) (err error) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MinLen: Int(2),
		MinLenMessage: "skills 长度必须 < {{MinLen}}",
	})
	return nil
}
func TestSliceMinLenMessage(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, SliceMinLenMessage{Skills: []string{}}, Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLenMessage{Skills: []string{"c"}}, Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLenMessage{Skills: []string{"c","d"}}, Report{
		Fail:    false,
		Message: "",
	})
}


type SliceMaxLen struct {
	Skills []string
}
func (v SliceMaxLen) VD(r *Rule) (err error) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MaxLen: Int(2),
	})
	return nil
}
func TestSliceMaxLen(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{"c"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{"c", "d"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{"c","d","e"}}, Report{
		Fail:    true,
		Message: "skills长度不能大于2",
	})
}

type SliceMaxLenMessage struct {
	Skills []string
}
func (v SliceMaxLenMessage) VD(r *Rule) (err error) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MaxLen: Int(2),
		MaxLenMessage: "skills 长度必须 > {{MaxLen}}",
	})
	return nil
}
func TestSliceMaxLenMessage(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{"c"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{"c", "d"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{"c","d","e"}}, Report{
		Fail:    true,
		Message: "skills 长度必须 > 2",
	})
}


type SliceMinMax struct {
	Skills []string
}
func (v SliceMinMax) VD(r *Rule) (err error) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "技能",
		MinLen: Int(2),
		MaxLen: Int(4),
	})
	return nil
}
func TestSliceMinMax(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{}}, Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c"}}, Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c", "d"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c","d","e"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c","d","e","f"}}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c","d","e","f", "g"}}, Report{
		Fail:    true,
		Message: "技能长度不能大于4",
	})
}