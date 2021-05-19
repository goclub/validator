package vd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestUtilEnumValues struct {
	Type string
}
func (TestUtilEnumValues) Dict() (dict struct{
	Type struct {
		Normal string
		Danger string
	}
}) {
	dict.Type.Normal = "normal"
	dict.Type.Danger = "danger"
	return
}
func (v TestUtilEnumValues) VD(r *Rule) (err error) {
	r.String(v.Type, StringSpec{
		Name:              "类型",
		Enum:              EnumValues(v.Dict().Type),
	})
	return nil
}
func Test_UtilEnumValues(t *testing.T) {
	data := TestUtilEnumValues{}
	
	assert.Equal(t, EnumValues(data.Dict().Type), []string{"normal", "danger"})
}
