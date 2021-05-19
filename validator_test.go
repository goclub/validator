package vd

import (
	"testing"
)

type RequiredOne struct {
	Name string
}
func (v RequiredOne) VD(r *Rule) error {
	r.String(v.Name, StringSpec{
		Name: "姓名",
	})
	return nil
}
func Test_RequiredOne (t *testing.T) {
	c := NewCN()
	
	CheckEqualAndNoError(t, c, RequiredOne{}, Report{
		Fail:    true,
		Message: "姓名必填",
	})
	CheckEqualAndNoError(t, c, RequiredOne{Name:"n"}, Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredTwo struct {
	Name string
	Title string
}
func (v RequiredTwo) VD(r *Rule) error {
	r.String(v.Name, StringSpec{
		Name: "姓名",
	})
	r.String(v.Title, StringSpec{
		Name: "标题",
	})
	return nil
}
func Test_RequiredTwo (t *testing.T) {
	c := NewCN()
	
	CheckEqualAndNoError(t, c, RequiredTwo{}, Report{
		Fail:    true,
		Message: "姓名必填",
	})
	CheckEqualAndNoError(t, c, RequiredTwo{Name:"n"}, Report{
		Fail:    true,
		Message: "标题必填",
	})
	CheckEqualAndNoError(t, c, RequiredTwo{Name:"n",Title:"1"}, Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredThree struct {
	Name string
	Title string
}
func (v RequiredThree) VD(r *Rule) error {
	r.String(v.Name, StringSpec{
		Name: "姓名",
		AllowEmpty: true,
	})
	r.String(v.Title, StringSpec{
		Name: "标题",
	})
	return nil
}
func Test_RequiredThree (t *testing.T) {
	c := NewCN()
	
	CheckEqualAndNoError(t, c, RequiredThree{}, Report{
		Fail:    true,
		Message: "标题必填",
	})
	CheckEqualAndNoError(t, c, RequiredThree{Name:"n",Title:"1"}, Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredFour struct {
	Name  string
	Title string
}
func (v RequiredFour) VD(r *Rule) error {
	r.String(v.Name, StringSpec{
		Name: "姓名",
	})
	r.String(v.Title, StringSpec{
		Name: "标题",
	})
	return nil
}
func Test_RequiredFour (t *testing.T) {
	c := NewCN()
	
	CheckEqualAndNoError(t, c, RequiredFour{}, Report{
		Fail:    true,
		Message: "姓名必填",
	})
	CheckEqualAndNoError(t, c, RequiredFour{Name:"n",Title:""}, Report{
		Fail:    true,
		Message: "标题必填",
	})
}
