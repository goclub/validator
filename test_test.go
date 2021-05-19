package vd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func CheckEqualAndNoError(t *testing.T, checker Checker, data Data, report Report) {
	r, err := checker.Check(data) ; assert.NoError(t, err)
	assert.Equal(t, r, report)
}
