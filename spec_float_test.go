package vd

import (
	"testing"
)


type FloatMin struct {
	Age float64
}
func (v FloatMin) VD(r *Rule) (err error) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Min: Float(18.2),
	})
	return nil
}

func TestFloatMin(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatMin{Age:17}, Report{
		Fail:    true,
		Message: "年龄不能小于18.2",
	})
	CheckEqualAndNoError(t, checker, FloatMin{Age:18.1}, Report{
		Fail:    true,
		Message: "年龄不能小于18.2",
	})
	CheckEqualAndNoError(t, checker, FloatMin{Age:18.2}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMin{Age:18.3}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMin{Age:19}, Report{
		Fail:    false,
		Message: "",
	})
}

type FloatMinMessage struct {
	Age float64
}

func (v FloatMinMessage) VD(r *Rule) (err error) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Min: Float(18.2),
		MinMessage:"年龄不可以小于{{Min}}",
	})
	return nil
}
func TestFloatMinMessage(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatMinMessage{Age:17}, Report{
		Fail:    true,
		Message: "年龄不可以小于18.2",
	})
	CheckEqualAndNoError(t, checker, FloatMinMessage{Age:18}, Report{
		Fail:    true,
		Message: "年龄不可以小于18.2",
	})
	CheckEqualAndNoError(t, checker, FloatMinMessage{Age:19}, Report{
		Fail:    false,
		Message: "",
	})
}


type FloatMax struct {
	Age float64
}
func (v FloatMax) VD(r *Rule) (err error) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Max: Float(18.2),
	})
	return nil
}
func TestFloatMax(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatMax{Age:17}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMax{Age:18.2}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMax{Age:18.3}, Report{
		Fail:    true,
		Message: "年龄不能大于18.2",
	})
	CheckEqualAndNoError(t, checker, FloatMax{Age:18.4}, Report{
		Fail:    true,
		Message: "年龄不能大于18.2",
	})
	CheckEqualAndNoError(t, checker, FloatMax{Age:19}, Report{
		Fail:    true,
		Message: "年龄不能大于18.2",
	})
}

type FloatMaxMessage struct {
	Age float64
}
func (v FloatMaxMessage) VD(r *Rule) (err error) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Max: Float(18),
		MaxMessage:"年龄不可以大于{{Max}}",
	})
	return nil
}
func TestFloatMaxMessage(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatMaxMessage{Age:17}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMaxMessage{Age:18}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMaxMessage{Age:19}, Report{
		Fail:    true,
		Message: "年龄不可以大于18",
	})
}
type FloatMinMax struct {
	Age float64
}
func (v FloatMinMax) VD (r *Rule) (err error) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Min: Float(2),
		Max: Float(4),
	})
	return nil
}
func TestFloatMinMax(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatMinMax{Age: 0}, Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	CheckEqualAndNoError(t, checker, FloatMinMax{Age: 1}, Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	CheckEqualAndNoError(t, checker, FloatMinMax{Age: 2}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMinMax{Age: 3}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMinMax{Age: 4}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatMinMax{Age: 5}, Report{
		Fail:    true,
		Message: "年龄不能大于4",
	})

}
type FloatPattern struct {
	Number float64
}
func (v FloatPattern) VD (r *Rule) (err error) {
	r.Float(v.Number, FloatSpec{
		Name: "号码",
		Pattern: []string{`^138`},
		PatternMessage: "{{Name}}必须以138开头",
	})
	return nil
}
func TestFloatPattern(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatPattern{Number: 11384}, Report{
		Fail:    true,
		Message: "号码必须以138开头",
	})
	CheckEqualAndNoError(t, checker, FloatPattern{Number: 138}, Report{
		Fail:    false,
		Message: "",
	})
}

type FloatBanPattern struct {
	Number float64
}
func (v FloatBanPattern) VD (r *Rule) (err error) {
	r.Float(v.Number, FloatSpec{
		Name: "号码",
		BanPattern: []string{`^138`, `^178`},
		PatternMessage: "{{Name}}不允许以138和178开头",
	})
	return
}
func TestFloatBanPattern(t *testing.T) {

	checker := NewCN()
	CheckEqualAndNoError(t, checker, FloatBanPattern{Number: 11384}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, FloatBanPattern{Number: 138}, Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
	CheckEqualAndNoError(t, checker, FloatBanPattern{Number: 178}, Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
}