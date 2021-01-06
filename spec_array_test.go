package vd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArrayMinLen struct {
	Skills []string
}
func (v ArrayMinLen) VD(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MinLen: Int(2),
	})
}
func TestArrayMinLen(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(ArrayMinLen{Skills: []string{}}), Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	assert.Equal(t, checker.Check(ArrayMinLen{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	assert.Equal(t, checker.Check(ArrayMinLen{Skills: []string{"c","d"}}), Report{
		Fail:    false,
		Message: "",
	})
}

type ArrayMinLenMessage struct {
	Skills []string
}
func (v ArrayMinLenMessage) VD(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MinLen: Int(2),
		MinLenMessage: "skills 长度必须 < {{MinLen}}",
	})
}
func TestArrayMinLenMessage(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(ArrayMinLenMessage{Skills: []string{}}), Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	assert.Equal(t, checker.Check(ArrayMinLenMessage{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	assert.Equal(t, checker.Check(ArrayMinLenMessage{Skills: []string{"c","d"}}), Report{
		Fail:    false,
		Message: "",
	})
}


type ArrayMaxLen struct {
	Skills []string
}
func (v ArrayMaxLen) VD(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MaxLen: Int(2),
	})
}
func TestArrayMaxLen(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(ArrayMaxLen{Skills: []string{}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMaxLen{Skills: []string{"c"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMaxLen{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMaxLen{Skills: []string{"c","d","e"}}), Report{
		Fail:    true,
		Message: "skills长度不能大于2",
	})
}

type ArrayMaxLenMessage struct {
	Skills []string
}
func (v ArrayMaxLenMessage) VD(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MaxLen: Int(2),
		MaxLenMessage: "skills 长度必须 > {{MaxLen}}",
	})
}
func TestArrayMaxLenMessage(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(ArrayMaxLenMessage{Skills: []string{}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMaxLenMessage{Skills: []string{"c"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMaxLenMessage{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMaxLenMessage{Skills: []string{"c","d","e"}}), Report{
		Fail:    true,
		Message: "skills 长度必须 > 2",
	})
}


type ArrayMinMax struct {
	Skills []string
}
func (v ArrayMinMax) VD(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "技能",
		MinLen: Int(2),
		MaxLen: Int(4),
	})
}
func TestArrayMinMax(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(ArrayMinMax{Skills: []string{}}), Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	assert.Equal(t, checker.Check(ArrayMinMax{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	assert.Equal(t, checker.Check(ArrayMinMax{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMinMax{Skills: []string{"c","d","e"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMinMax{Skills: []string{"c","d","e","f"}}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(ArrayMinMax{Skills: []string{"c","d","e","f", "g"}}), Report{
		Fail:    true,
		Message: "技能长度不能大于4",
	})
}