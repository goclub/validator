package vd

import (
	"testing"
)

type SpecStringMinLen struct {
	Name string
}
func (s SpecStringMinLen) VD(r *Rule) (err error){
	r.String(s.Name, StringSpec{
		Name:              "姓名",
		MinRuneLen:        4,
	})
	return nil
};
type SpecStringMinLenCustomMessage struct {
	Name string
}
func (s SpecStringMinLenCustomMessage) VD(r *Rule) (err error){
	r.String(s.Name, StringSpec{
		Name:              "姓名",
		MinRuneLen:        4,
		MinRuneLenMessage: "姓名长度不能小于{{MinRuneLen}}位,你输入的是{{Value}}",
	})
	return nil
}
func Test_SpecString_MinLen(t *testing.T) {
	c := NewCN()
	
	CheckEqualAndNoError(t, c, SpecStringMinLen{Name:"ni"}, Report{
		Fail:    true,
		Message: "姓名长度不能小于4",
	})
	CheckEqualAndNoError(t, c, SpecStringMinLen{Name:"nim"}, Report{
		Fail:    true,
		Message: "姓名长度不能小于4",
	})
	CheckEqualAndNoError(t, c, SpecStringMinLen{Name:"nimo"}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMinLen{Name:"nimoc"}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMinLenCustomMessage{Name:"ni"}, Report{
		Fail:    true,
		Message: "姓名长度不能小于4位,你输入的是ni",
	})
}

type SpecStringMaxLen struct {
	Name string 
}
func (s SpecStringMaxLen) VD(r *Rule) (err error){
	r.String(s.Name, StringSpec{
		Name:              "姓名",
		MaxRuneLen:        4,
	})
	return nil
}
type SpecStringMaxLenCustomMessage struct {
	Name string
}
func (s SpecStringMaxLenCustomMessage) VD(r *Rule) (err error){
	r.String(s.Name, StringSpec{
		Name:              "姓名",
		MaxRuneLen:        4,
		MaxRuneLenMessage: "姓名长度不能大于{{MaxRuneLen}}位,你输入的是{{Value}}",
	})
	return nil
}
func Test_SpecString_MaxLen(t *testing.T) {
	c := NewCN()
	
	CheckEqualAndNoError(t, c, SpecStringMaxLen{Name:"nimoc"}, Report{
		Fail:    true,
		Message: "姓名长度不能大于4",
	})
	CheckEqualAndNoError(t, c, SpecStringMaxLen{Name:"nimo"}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMaxLen{Name:"nim"}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMaxLenCustomMessage{Name:"nimoc"}, Report{
		Fail:    true,
		Message: "姓名长度不能大于4位,你输入的是nimoc",
	})
}
type SpecStringPattern struct {
	Name string
	Title string
	More string 
}
func (s SpecStringPattern) VD (r *Rule) (err error){
	r.String(s.Name, StringSpec{
		Name:              "姓名",
		Pattern:		   []string{"^nimo"},
	})
	r.String(s.Title, StringSpec{
		Name: "标题",
		Pattern: []string{`abc$`},
		PatternMessage: "{{Name}}必须以abc为结尾",
	})
	r.String(s.More, StringSpec{
		AllowEmpty: true,
		Name: "更多",
		Pattern:[]string{`^a`, `a$`},
		PatternMessage: "{{Name}}开始结尾必须是a",
	})
	return nil
}
func TestSpecStringPattern(t *testing.T) {
	
	c := NewCN()
	{
		CheckEqualAndNoError(t, c, SpecStringPattern{
			Name: "nimo",
			Title: "abc",
		}, Report{
			Fail:    true,
			Message: "更多开始结尾必须是a",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringPattern{
			Name: "xnimo",
			Title: "abc",
		}, Report{
			Fail:    true,
			Message: "姓名格式错误",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringPattern{
			Name: "nimo",
			Title: "abcd",
		}, Report{
			Fail:    true,
			Message: "标题必须以abc为结尾",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringPattern{
			Name: "nimo",
			Title: "abcd",
			More: "c",
		}, Report{
			Fail:    true,
			Message: "标题必须以abc为结尾",
		})
	}
}

type SpecStringBanPattern struct {
	Name string
	Title string
	More string
}
func (s SpecStringBanPattern) VD (r *Rule)(err error){
	r.String(s.Name, StringSpec{
		Name:              "姓名",
		BanPattern:		   []string{"fuck"},
		PatternMessage: "{{Name}}不允许出现敏感词",
	})
	r.String(s.Title, StringSpec{
		Name: "标题",
		BanPattern: []string{`fuck`},
		PatternMessage: "{{Name}}不允许出现敏感词",
	})
	r.String(s.More, StringSpec{
		AllowEmpty: true,
		Name: "更多",
		BanPattern: []string{`fuck`, `dick`},
		PatternMessage: "{{Name}}不允许出现敏感词:{{BanPattern}}",
	})
	return nil
}
func TestSpecStringBanPattern(t *testing.T) {
	
	c := NewCN()
	{
		CheckEqualAndNoError(t, c, SpecStringBanPattern{
			Name: "nimo",
			Title: "nimo",
			More: "nimo",
		}, Report{
			Fail:    false,
			Message: "",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringBanPattern{
			Name: "fuck",
			Title: "nimo",
			More: "nimo",
		}, Report{
			Fail:    true,
			Message: "姓名不允许出现敏感词",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringBanPattern{
			Name: "nimo",
			Title: "fuck",
			More: "nimo",
		}, Report{
			Fail:    true,
			Message: "标题不允许出现敏感词",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringBanPattern{
			Name: "nimo",
			Title: "nimo",
			More: "fuck",
		}, Report{
			Fail:    true,
			Message: "更多不允许出现敏感词:[fuck dick]",
		})
	}
	{
		CheckEqualAndNoError(t, c, SpecStringBanPattern{
			Name: "nimo",
			Title: "nimo",
			More: "dick",
		}, Report{
			Fail:    true,
			Message: "更多不允许出现敏感词:[fuck dick]",
		})
	}
}
type SpecStringEnum struct {
	Type SomeType
}
type SomeType string
func (v SomeType) String() string {
	return string(v)
}
func (SomeType) Enum() (e struct{
	Normal SomeType
	Danger SomeType
}) {
	e.Normal = "normal"
	e.Danger = "danger"
	return
}
func (s SpecStringEnum) VD(r *Rule) (err error){
	r.String(s.Type.String(), StringSpec{
		Name: "类型",
		Enum: EnumValues(s.Type.Enum()),
	})
	return nil
}
func TestStringSpec_CheckEnum (t *testing.T) {
	
	c := NewCN()
	CheckEqualAndNoError(t, c, SpecStringEnum{
		Type: "normal1",
	}, Report{
		Fail:    true,
		Message: "类型参数错误，只允许(normal danger)",
	})
}
type SpecStringMinMax struct {
	Name string
}
func (v SpecStringMinMax) VD(r *Rule) (err error){
	r.String(v.Name, StringSpec{
		Name:              "姓名",
		AllowEmpty: 	   true,
		MinRuneLen:        2,
		MaxRuneLen:        4,
	})
	return nil
}
func TestSpectStringMinMax(t *testing.T) {
	
	c := NewCN()
	CheckEqualAndNoError(t, c, SpecStringMinMax{
		Name: "",
	}, Report{
		Fail:    true,
		Message: "姓名长度不能小于2",
	})
	CheckEqualAndNoError(t, c, SpecStringMinMax{
		Name: "1",
	}, Report{
		Fail:    true,
		Message: "姓名长度不能小于2",
	})
	CheckEqualAndNoError(t, c, SpecStringMinMax{
		Name: "12",
	}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMinMax{
		Name: "123",
	}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMinMax{
		Name: "1234",
	}, Report{
		Fail:    false,
		Message: "",
	})
	CheckEqualAndNoError(t, c, SpecStringMinMax{
		Name: "12345",
	}, Report{
		Fail:    true,
		Message: "姓名长度不能大于4",
	})

}

type SpecStringEmail struct {
	Email string
	OtherEmail string
}
func (v SpecStringEmail) VD(r *Rule) (err error){
	r.String(v.Email, StringSpec{
		Name: "邮箱",
		Ext:  []StringSpec{
			ExtString{}.Email(),
		},
	})
	r.String(v.OtherEmail, ExtString{}.Email().NameIs("附属邮箱"))
	return nil
}
func TestStringEmail(t *testing.T) {
	c := NewCN()
	CheckEqualAndNoError(t, c, SpecStringEmail{
		Email: "12345",
		OtherEmail: "mail@github.com",
	}, Report{
		Fail:    true,
		Message: "邮箱格式错误",
	})
	CheckEqualAndNoError(t, c, SpecStringEmail{
		Email: "12345@qq.com",
		OtherEmail: "mailithub.com",
	}, Report{
		Fail:    true,
		Message: "附属邮箱格式错误",
	})
}
