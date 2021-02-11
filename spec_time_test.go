package vd_test

import (
	xtime "github.com/goclub/time"
	vd "github.com/goclub/validator"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
var now = time.Now()
type TimeNotAllowZero struct {
	A time.Time
	B time.Time
}
func (v TimeNotAllowZero) VD(r *vd.Rule) {
	r.Time(v.A, vd.TimeSpec{
		Name:"时间1",
		AllowZero: true,
	})
	r.Time(v.B, vd.TimeSpec{
		Name:"时间2",
		AllowZero: false,
	})
}
func TestRule_TimeAlowZero(t *testing.T) {
	c := vd.NewCN()

	assert.Equal(t, c.Check(TimeNotAllowZero{
		A: time.Time{},
		B: time.Time{},
	}), vd.Report{
		Fail:    true,
		Message: "时间2不能为空",
	})
	assert.Equal(t, c.Check(TimeNotAllowZero{
		A: time.Now(),
		B: time.Time{},
	}), vd.Report{
		Fail:    true,
		Message: "时间2不能为空",
	})
	assert.Equal(t, c.Check(TimeNotAllowZero{
		A: time.Now(),
		B: time.Now(),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, c.Check(TimeNotAllowZero{
		A: time.Time{},
		B: time.Now(),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
}

type TimeRange1 struct {
	StartTime time.Time
	EndTime time.Time
}
func (v TimeRange1) VD(r *vd.Rule) {
	r.Time(v.StartTime, vd.TimeSpec{
		Name:"开始时间",
		BeforeIt: v.EndTime,
	})
	r.Time(v.EndTime, vd.TimeSpec{
		Name:"结束时间",
		AfterIt: v.StartTime,
	})
}

type TimeRange2 struct {
	StartTime time.Time
	EndTime time.Time
}
func (v TimeRange2) VD(r *vd.Rule) {
	r.TimeRange(vd.TimeRange{"开始时间", v.StartTime,"结束时间", v.EndTime,})
}

func TestRule_TimeRange2(t *testing.T) {
	c := vd.NewCN()
	// start end 互相约束
	assert.Equal(t, c.Check(TimeRange2{
		StartTime: now,
		EndTime: now.Add(time.Second*2),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, c.Check(TimeRange2{
		StartTime: now,
		EndTime: now,
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, c.Check(TimeRange2{
		StartTime: now.Add(time.Second),
		EndTime: now,
	}), vd.Report{
		Fail:    true,
		Message: "开始时间" + xtime.FormatChinaTime(now.Add(time.Second)) + "必须等于" + xtime.FormatChinaTime(now) + "或之前",
	})
	assert.Equal(t, c.Check(TimeRange2{
		StartTime: now,
		EndTime: now.Add(-time.Second),
	}), vd.Report{
		Fail:    true,
		Message: "开始时间" + xtime.FormatChinaTime(now) + "必须等于" + xtime.FormatChinaTime(now.Add(-time.Second)) + "或之前",
	})
}

type TestAfterIt struct {
	V time.Time
}

func (v TestAfterIt) VD(r *vd.Rule) {
	r.Time(v.V, vd.TimeSpec{
		Name:"v",
		AfterIt: now,
	})
}

func TestRule_TimeAfterIt(t *testing.T) {
	c := vd.NewCN()
	assert.Equal(t, c.Check(TestAfterIt{
		now,
	}), vd.Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now) + "必须在" + xtime.FormatChinaTime(now) + "之后",
	})
	assert.Equal(t, c.Check(TestAfterIt{
		now.Add(-time.Second),
	}), vd.Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(-time.Second)) + "必须在" + xtime.FormatChinaTime(now) + "之后",
	})
	assert.Equal(t, c.Check(TestAfterIt{
		now.Add(time.Second*1),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
}

type TestAfterOrEqualIt struct {
	V time.Time
}

func (v TestAfterOrEqualIt) VD(r *vd.Rule) {
	r.Time(v.V, vd.TimeSpec{
		Name:"v",
		AfterOrEqualIt: now,
	})
}


func TestRule_TimeAfterOrEqualIt(t *testing.T) {
	c := vd.NewCN()
	assert.Equal(t, c.Check(TestAfterOrEqualIt{
		now,
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, c.Check(TestAfterOrEqualIt{
		now.Add(-time.Second),
	}), vd.Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(-time.Second)) + "必须等于" + xtime.FormatChinaTime(now) + "或之后",
	})
	assert.Equal(t, c.Check(TestAfterOrEqualIt{
		now.Add(time.Second*1),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
}


type TestBeforeIt struct {
	V time.Time
}

func (v TestBeforeIt) VD(r *vd.Rule) {
	r.Time(v.V, vd.TimeSpec{
		Name:"v",
		BeforeIt: now,
	})
}

func TestRule_TimeBeforeIt(t *testing.T) {
	c := vd.NewCN()
	assert.Equal(t, c.Check(TestBeforeIt{
		now,
	}), vd.Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now) + "必须在" + xtime.FormatChinaTime(now) + "之前",
	})
	assert.Equal(t, c.Check(TestBeforeIt{
		now.Add(time.Second),
	}), vd.Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(time.Second)) + "必须在" + xtime.FormatChinaTime(now) + "之前",
	})
	assert.Equal(t, c.Check(TestBeforeIt{
		now.Add(-time.Second*1),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
}

type TestBeforeOrEqualIt struct {
	V time.Time
}

func (v TestBeforeOrEqualIt) VD(r *vd.Rule) {
	r.Time(v.V, vd.TimeSpec{
		Name:"v",
		BeforeOrEqualIt: now,
	})
}


func TestRule_TimeBeforeOrEqualIt(t *testing.T) {
	c := vd.NewCN()
	assert.Equal(t, c.Check(TestBeforeOrEqualIt{
		now,
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
	assert.Equal(t, c.Check(TestBeforeOrEqualIt{
		now.Add(time.Second),
	}), vd.Report{
		Fail:    true,
		Message: "v" + xtime.FormatChinaTime(now.Add(time.Second)) + "必须等于" + xtime.FormatChinaTime(now) + "或之前",
	})
	assert.Equal(t, c.Check(TestBeforeOrEqualIt{
		now.Add(-time.Second*1),
	}), vd.Report{
		Fail:    false,
		Message: "",
	})
}