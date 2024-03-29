package vd

import (
	xtime "github.com/goclub/time"
	"testing"
	"time"
)
var now = time.Now()
type TimeNotAllowZero struct {
	A time.Time
	B time.Time
}
func (v TimeNotAllowZero) VD(r *Rule) (err error) {
	r.Time(v.A, TimeSpec{
		Name:"时间1",
		AllowZero: true,
	})
	r.Time(v.B, TimeSpec{
		Name:"时间2",
		AllowZero: false,
	})
	return nil
}
func TestRule_TimeAlowZero(t *testing.T) {
	c := NewCN()

	CheckEqualAndNoError(t, c, TimeNotAllowZero{
		A: time.Time{},
		B: time.Time{},
	}, Report{
		Fail:    true,
		Message: "时间2不能为空",
	})
	CheckEqualAndNoError(t, c, TimeNotAllowZero{
		A: time.Now(),
		B: time.Time{},
	}, Report{
		Fail:    true,
		Message: "时间2不能为空",
	})
	CheckEqualAndNoError(t, c, TimeNotAllowZero{
		A: time.Now(),
		B: time.Now(),
	}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, TimeNotAllowZero{
		A: time.Time{},
		B: time.Now(),
	}, Report{
		Fail:    false,
		Message: "",
	})
}

type TimeRange1 struct {
	StartTime time.Time
	EndTime time.Time
}
func (v TimeRange1) VD(r *Rule) {
	r.Time(v.StartTime, TimeSpec{
		Name:"开始时间",
		BeforeIt: v.EndTime,
	})
}

type TestAfterIt struct {
	V time.Time
}

func (v TestAfterIt) VD(r *Rule) error {
	r.Time(v.V, TimeSpec{
		Name:"v",
		AfterIt: now,
	})
	return nil
}

func TestRule_TimeAfterIt(t *testing.T) {
	c := NewCN()
	CheckEqualAndNoError(t, c, TestAfterIt{
		now,
	}, Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now) + "必须在" + xtime.FormatChinaTime(now) + "之后",
	})
	CheckEqualAndNoError(t, c, TestAfterIt{
		now.Add(-time.Second),
	}, Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(-time.Second)) + "必须在" + xtime.FormatChinaTime(now) + "之后",
	})
	CheckEqualAndNoError(t, c, TestAfterIt{
		now.Add(time.Second*1),
	}, Report{
		Fail:    false,
		Message: "",
	})
}

type TestAfterOrEqualIt struct {
	V time.Time
}

func (v TestAfterOrEqualIt) VD(r *Rule) error {
	r.Time(v.V, TimeSpec{
		Name:"v",
		AfterOrEqualIt: now,
	})
	return nil
}


func TestRule_TimeAfterOrEqualIt(t *testing.T) {
	c := NewCN()
	CheckEqualAndNoError(t, c, TestAfterOrEqualIt{
		now,
	}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, TestAfterOrEqualIt{
		now.Add(-time.Second),
	}, Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(-time.Second)) + "必须在" + xtime.FormatChinaTime(now) + "之后，或等于",
	})
	CheckEqualAndNoError(t, c, TestAfterOrEqualIt{
		now.Add(time.Second*1),
	}, Report{
		Fail:    false,
		Message: "",
	})
}


type TestBeforeIt struct {
	V time.Time
}

func (v TestBeforeIt) VD(r *Rule) error {
	r.Time(v.V, TimeSpec{
		Name:"v",
		BeforeIt: now,
	})
	return nil
}

func TestRule_TimeBeforeIt(t *testing.T) {
	c := NewCN()
	CheckEqualAndNoError(t, c, TestBeforeIt{
		now,
	}, Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now) + "必须在" + xtime.FormatChinaTime(now) + "之前",
	})
	CheckEqualAndNoError(t, c, TestBeforeIt{
		now.Add(time.Second),
	}, Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(time.Second)) + "必须在" + xtime.FormatChinaTime(now) + "之前",
	})
	CheckEqualAndNoError(t, c, TestBeforeIt{
		now.Add(-time.Second*1),
	}, Report{
		Fail:    false,
		Message: "",
	})
}

type TestBeforeOrEqualIt struct {
	V time.Time
}

func (v TestBeforeOrEqualIt) VD(r *Rule) (err error) {
	r.Time(v.V, TimeSpec{
		Name:"v",
		BeforeOrEqualIt: now,
	})
	return nil
}


func TestRule_TimeBeforeOrEqualIt(t *testing.T) {
	c := NewCN()
	CheckEqualAndNoError(t, c, TestBeforeOrEqualIt{
		now,
	}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, TestBeforeOrEqualIt{
		now.Add(time.Second),
	}, Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(time.Second)) + "必须在" + xtime.FormatChinaTime(now) + "之前，或等于",
	})
	CheckEqualAndNoError(t, c, TestBeforeOrEqualIt{
		now.Add(-time.Second*1),
	}, Report{
		Fail:    false,
		Message: "",
	})
}