package vd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	{
		v := []string{"a", "b", "a"}
		isRepeat, repeatElement := uniqueStrings(v)
		assert.Equal(t, isRepeat, true)
		assert.Equal(t, repeatElement, "a")
	}
	{
		v := []string{"a","b"}
		isRepeat, repeatElement := uniqueStrings(v)
		assert.Equal(t, isRepeat, false)
		assert.Equal(t, repeatElement, "")
	}
}
