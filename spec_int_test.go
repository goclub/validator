package vd

import (
	"testing"
)

// type IntNotAllowZero struct {
// 	Age int
// }
// func (v IntNotAllowZero) VD(r *Rule) (err error) {
// 	r.Int(v.Age, IntSpec{
// 		Name: "年龄",
// 	})
// }
// func TestIntNotAllowZero(t *testing.T) {
// 	
// 	checker := NewCN()
// 	CheckEqualAndNoError(t, checker, IntNotAllowZero{Age:0}, Report{
// 		Fail:    true,
// 		Message: "年龄不允许为0",
// 	})
// 	CheckEqualAndNoError(t, checker, IntNotAllowZero{Age:1}, Report{
// 		Fail:    false,
// 		Message: "",
// 	})
// }

// type IntAllowZero struct {
// 	Age int
// }
// func (v IntAllowZero) VD(r *Rule) (err error) {
// 	r.Int(v.Age, IntSpec{
// 		Name: "年龄",
// 		AllowZero: true,
// 	})
// }
// func TestIntAllowZero(t *testing.T) {
// 	
// 	checker := NewCN()
// 	CheckEqualAndNoError(t, checker, IntAllowZero{Age:0}, Report{
// 		Fail:    false,
// 		Message: "",
// 	})
// 	CheckEqualAndNoError(t, checker, IntAllowZero{Age:1}, Report{
// 		Fail:    false,
// 		Message: "",
// 	})
// }

type IntMin struct {
	Age int
}
func (v IntMin) VD(r *Rule) (err error) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: Int(18),
	})
	return nil
}
func TestIntMin(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntMin{Age:17}, Report{
		Fail:    true,
		Message: "年龄不能小于18",
	})
	CheckEqualAndNoError(t, checker, IntMin{Age:18}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMin{Age:19}, Report{
		Fail:    false,
		Message: "",
	})
}

type IntMinMessage struct {
	Age int
}

func (v IntMinMessage) VD(r *Rule) (err error) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: Int(18),
		MinMessage:"年龄不可以小于{{Min}}",
	})
	return nil
}
func TestIntMinMessage(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntMinMessage{Age:17}, Report{
		Fail:    true,
		Message: "年龄不可以小于18",
	})
	CheckEqualAndNoError(t, checker, IntMinMessage{Age:18}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMinMessage{Age:19}, Report{
		Fail:    false,
		Message: "",
	})
}


type IntMax struct {
	Age int
}
func (v IntMax) VD(r *Rule) (err error) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Max: Int(18),
	})
	return nil
}
func TestIntMax(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntMax{Age:17}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMax{Age:18}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMax{Age:19}, Report{
		Fail:    true,
		Message: "年龄不能大于18",
	})
}

type IntMaxMessage struct {
	Age int
}
func (v IntMaxMessage) VD(r *Rule) (err error) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Max: Int(18),
		MaxMessage:"年龄不可以大于{{Max}}",
	})
	return nil
}
func TestIntMaxMessage(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntMaxMessage{Age:17}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMaxMessage{Age:18}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMaxMessage{Age:19}, Report{
		Fail:    true,
		Message: "年龄不可以大于18",
	})
}
type IntMinMax struct {
	Age int
}
func (v IntMinMax) VD (r *Rule) (err error) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: Int(2),
		Max: Int(4),
	})
	return nil
}
func TestIntMinMax(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntMinMax{Age: 0}, Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	CheckEqualAndNoError(t, checker, IntMinMax{Age: 1}, Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	CheckEqualAndNoError(t, checker, IntMinMax{Age: 2}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMinMax{Age: 3}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMinMax{Age: 4}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntMinMax{Age: 5}, Report{
		Fail:    true,
		Message: "年龄不能大于4",
	})

}
type IntPattern struct {
	Number int
}
func (v IntPattern) VD (r *Rule) (err error) {
	r.Int(v.Number, IntSpec{
		Name: "号码",
		Pattern: []string{`^138`},
		PatternMessage: "{{Name}}必须以138开头",
	})
	return nil
}
func TestIntPattern(t *testing.T) {
	
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntPattern{Number: 11384}, Report{
		Fail:    true,
		Message: "号码必须以138开头",
	})
	CheckEqualAndNoError(t, checker, IntPattern{Number: 138}, Report{
		Fail:    false,
		Message: "",
	})
}

type IntBanPattern struct {
	Number int
}
func (v IntBanPattern) VD (r *Rule) (err error) {
	r.Int(v.Number, IntSpec{
		Name: "号码",
		BanPattern: []string{`^138`, `^178`},
		PatternMessage: "{{Name}}不允许以138和178开头",
	})
	return nil
}
func TestIntBanPattern(t *testing.T) {
	checker := NewCN()
	CheckEqualAndNoError(t, checker, IntBanPattern{Number: 11384}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, checker, IntBanPattern{Number: 138}, Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
	CheckEqualAndNoError(t, checker, IntBanPattern{Number: 178}, Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
}