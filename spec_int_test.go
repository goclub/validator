package vd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// type IntNotAllowZero struct {
// 	Age int
// }
// func (v IntNotAllowZero) VD(r *Rule) {
// 	r.Int(v.Age, IntSpec{
// 		Name: "年龄",
// 	})
// }
// func TestIntNotAllowZero(t *testing.T) {
// 	
// 	checker := NewCN()
// 	assert.Equal(t, checker.Check(IntNotAllowZero{Age:0}), Report{
// 		Fail:    true,
// 		Message: "年龄不允许为0",
// 	})
// 	assert.Equal(t, checker.Check(IntNotAllowZero{Age:1}), Report{
// 		Fail:    false,
// 		Message: "",
// 	})
// }

// type IntAllowZero struct {
// 	Age int
// }
// func (v IntAllowZero) VD(r *Rule) {
// 	r.Int(v.Age, IntSpec{
// 		Name: "年龄",
// 		AllowZero: true,
// 	})
// }
// func TestIntAllowZero(t *testing.T) {
// 	
// 	checker := NewCN()
// 	assert.Equal(t, checker.Check(IntAllowZero{Age:0}), Report{
// 		Fail:    false,
// 		Message: "",
// 	})
// 	assert.Equal(t, checker.Check(IntAllowZero{Age:1}), Report{
// 		Fail:    false,
// 		Message: "",
// 	})
// }

type IntMin struct {
	Age int
}
func (v IntMin) VD(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: Int(18),
	})
}
func TestIntMin(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(IntMin{Age:17}), Report{
		Fail:    true,
		Message: "年龄不能小于18",
	})
	assert.Equal(t, checker.Check(IntMin{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMin{Age:19}), Report{
		Fail:    false,
		Message: "",
	})
}

type IntMinMessage struct {
	Age int
}

func (v IntMinMessage) VD(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: Int(18),
		MinMessage:"年龄不可以小于{{Min}}",
	})
}
func TestIntMinMessage(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(IntMinMessage{Age:17}), Report{
		Fail:    true,
		Message: "年龄不可以小于18",
	})
	assert.Equal(t, checker.Check(IntMinMessage{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMinMessage{Age:19}), Report{
		Fail:    false,
		Message: "",
	})
}


type IntMax struct {
	Age int
}
func (v IntMax) VD(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Max: Int(18),
	})
}
func TestIntMax(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(IntMax{Age:17}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMax{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMax{Age:19}), Report{
		Fail:    true,
		Message: "年龄不能大于18",
	})
}

type IntMaxMessage struct {
	Age int
}
func (v IntMaxMessage) VD(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Max: Int(18),
		MaxMessage:"年龄不可以大于{{Max}}",
	})
}
func TestIntMaxMessage(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(IntMaxMessage{Age:17}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMaxMessage{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMaxMessage{Age:19}), Report{
		Fail:    true,
		Message: "年龄不可以大于18",
	})
}
type IntMinMax struct {
	Age int
}
func (v IntMinMax) VD (r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: Int(2),
		Max: Int(4),
	})
}
func TestIntMinMax(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(IntMinMax{Age: 0}), Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	assert.Equal(t, checker.Check(IntMinMax{Age: 1}), Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	assert.Equal(t, checker.Check(IntMinMax{Age: 2}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMinMax{Age: 3}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMinMax{Age: 4}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntMinMax{Age: 5}), Report{
		Fail:    true,
		Message: "年龄不能大于4",
	})

}
type IntPattern struct {
	Number int
}
func (v IntPattern) VD (r *Rule) {
	r.Int(v.Number, IntSpec{
		Name: "号码",
		Pattern: []string{`^138`},
		PatternMessage: "{{Name}}必须以138开头",
	})
}
func TestIntPattern(t *testing.T) {
	
	checker := NewCN()
	assert.Equal(t, checker.Check(IntPattern{Number: 11384}), Report{
		Fail:    true,
		Message: "号码必须以138开头",
	})
	assert.Equal(t, checker.Check(IntPattern{Number: 138}), Report{
		Fail:    false,
		Message: "",
	})
}

type IntBanPattern struct {
	Number int
}
func (v IntBanPattern) VD (r *Rule) {
	r.Int(v.Number, IntSpec{
		Name: "号码",
		BanPattern: []string{`^138`, `^178`},
		PatternMessage: "{{Name}}不允许以138和178开头",
	})
}
func TestIntBanPattern(t *testing.T) {
	checker := NewCN()
	assert.Equal(t, checker.Check(IntBanPattern{Number: 11384}), Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, checker.Check(IntBanPattern{Number: 138}), Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
	assert.Equal(t, checker.Check(IntBanPattern{Number: 178}), Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
}