package vdg

import (
	vd "github.com/goclub/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CheckEqualAndNoError(t *testing.T, checker vd.Checker, data vd.Data, report vd.Report) {
	r, err := checker.Check(data)
	assert.NoError(t, err)
	assert.Equal(t, r, report)
}
