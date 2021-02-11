package vd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SliceMinLen struct {
	Skills []string
}
func (v SliceMinLen) VD(r *Rule) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MinLen: Int(2),
	})
}
func TestSliceMinLen(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(SliceMinLen{Skills: []string{}}), Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	assert.Equal(t, checker.Check(SliceMinLen{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	assert.Equal(t, checker.Check(SliceMinLen{Skills: []string{"c","d"}}), Report{
		Fail:    false,
		Message: "",
	})
}

type SliceMinLenMessage struct {
	Skills []string
}
func (v SliceMinLenMessage) VD(r *Rule) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MinLen: Int(2),
		MinLenMessage: "skills 长度必须 < {{MinLen}}",
	})
}
func TestSliceMinLenMessage(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(SliceMinLenMessage{Skills: []string{}}), Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	assert.Equal(t, checker.Check(SliceMinLenMessage{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	assert.Equal(t, checker.Check(SliceMinLenMessage{Skills: []string{"c","d"}}), Report{
		Fail:    false,
		Message: "",
	})
}


type SliceMaxLen struct {
	Skills []string
}
func (v SliceMaxLen) VD(r *Rule) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MaxLen: Int(2),
	})
}
func TestSliceMaxLen(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(SliceMaxLen{Skills: []string{}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMaxLen{Skills: []string{"c"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMaxLen{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMaxLen{Skills: []string{"c","d","e"}}), Report{
		Fail:    true,
		Message: "skills长度不能大于2",
	})
}

type SliceMaxLenMessage struct {
	Skills []string
}
func (v SliceMaxLenMessage) VD(r *Rule) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "skills",
		MaxLen: Int(2),
		MaxLenMessage: "skills 长度必须 > {{MaxLen}}",
	})
}
func TestSliceMaxLenMessage(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(SliceMaxLenMessage{Skills: []string{}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMaxLenMessage{Skills: []string{"c"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMaxLenMessage{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMaxLenMessage{Skills: []string{"c","d","e"}}), Report{
		Fail:    true,
		Message: "skills 长度必须 > 2",
	})
}


type SliceMinMax struct {
	Skills []string
}
func (v SliceMinMax) VD(r *Rule) {
	r.Slice(len(v.Skills), SliceSpec{
		Name: "技能",
		MinLen: Int(2),
		MaxLen: Int(4),
	})
}
func TestSliceMinMax(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(SliceMinMax{Skills: []string{}}), Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	assert.Equal(t, checker.Check(SliceMinMax{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	assert.Equal(t, checker.Check(SliceMinMax{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMinMax{Skills: []string{"c","d","e"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMinMax{Skills: []string{"c","d","e","f"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(SliceMinMax{Skills: []string{"c","d","e","f", "g"}}), Report{
		Fail:    true,
		Message: "技能长度不能大于4",
	})
}