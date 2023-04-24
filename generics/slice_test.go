package vdg

import (
	vd "github.com/goclub/validator"
	"testing"
)

type SliceMinLen struct {
	Skills []string
}

func (v SliceMinLen) VD(r *vd.Rule) (err error) {
	Slice(r, v.Skills, SliceSpec{
		Name:   "skills",
		MinLen: vd.Int(2),
	})
	return nil
}
func TestSliceMinLen(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceMinLen{Skills: []string{}}, vd.Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLen{Skills: []string{"c"}}, vd.Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLen{Skills: []string{"c", "d"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
}

type SliceMinLenMessage struct {
	Skills []string
}

func (v SliceMinLenMessage) VD(r *vd.Rule) (err error) {
	Slice(r, v.Skills, SliceSpec{
		Name:          "skills",
		MinLen:        vd.Int(2),
		MinLenMessage: "skills 长度必须 < {{MinLen}}",
	})
	return nil
}
func TestSliceMinLenMessage(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceMinLenMessage{Skills: []string{}}, vd.Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLenMessage{Skills: []string{"c"}}, vd.Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	CheckEqualAndNoError(t, checker, SliceMinLenMessage{Skills: []string{"c", "d"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
}

type SliceMaxLen struct {
	Skills []string
}

func (v SliceMaxLen) VD(r *vd.Rule) (err error) {
	Slice(r, v.Skills, SliceSpec{
		Name:   "skills",
		MaxLen: vd.Int(2),
	})
	return nil
}
func TestSliceMaxLen(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{"c"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{"c", "d"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLen{Skills: []string{"c", "d", "e"}}, vd.Report{
		Fail:    true,
		Message: "skills长度不能大于2",
	})
}

type SliceMaxLenMessage struct {
	Skills []string
}

func (v SliceMaxLenMessage) VD(r *vd.Rule) (err error) {
	Slice(r, v.Skills, SliceSpec{
		Name:          "skills",
		MaxLen:        vd.Int(2),
		MaxLenMessage: "skills 长度必须 > {{MaxLen}}",
	})
	return nil
}
func TestSliceMaxLenMessage(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{"c"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{"c", "d"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMaxLenMessage{Skills: []string{"c", "d", "e"}}, vd.Report{
		Fail:    true,
		Message: "skills 长度必须 > 2",
	})
}

type SliceMinMax struct {
	Skills []string
}

func (v SliceMinMax) VD(r *vd.Rule) (err error) {
	Slice(r, v.Skills, SliceSpec{
		Name:   "技能",
		MinLen: vd.Int(2),
		MaxLen: vd.Int(4),
	})
	return nil
}
func TestSliceMinMax(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{}}, vd.Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c"}}, vd.Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c", "d"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c", "d", "e"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c", "d", "e", "f"}}, vd.Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, SliceMinMax{Skills: []string{"c", "d", "e", "f", "g"}}, vd.Report{
		Fail:    true,
		Message: "技能长度不能大于4",
	})
}

type SliceElement struct {
	Skills []string
}

func (v SliceElement) VD(r *vd.Rule) (err error) {
	Slice(r, v.Skills, SliceSpec{
		MinLen: vd.Int(2),
	}, func(elem string, r *vd.Rule, i int) (err error) {
		r.String(elem, vd.StringSpec{
			BanPattern:     []string{"fuck"},
			PatternMessage: "不能有脏话",
		})
		return
	})
	return
}
func TestSliceElement(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceElement{Skills: []string{}}, vd.Report{
		Fail:    true,
		Message: "长度不能小于2",
	})
	CheckEqualAndNoError(t, checker, SliceElement{Skills: []string{"fuck", "abc"}}, vd.Report{
		Fail:    true,
		Message: "不能有脏话",
	})
}

type SliceElement2 struct {
	Skills []string
}

func (v SliceElement2) VD(r *vd.Rule) (err error) {
	for _, skill := range v.Skills {
		r.String(skill, vd.StringSpec{
			BanPattern:     []string{"fuck"},
			PatternMessage: "不能有脏话",
		})
	}
	return
}
func TestSliceElement2(t *testing.T) {

	checker := vd.NewCN()
	CheckEqualAndNoError(t, checker, SliceElement2{Skills: []string{"fuck", "abc"}}, vd.Report{
		Fail:    true,
		Message: "不能有脏话",
	})
}
