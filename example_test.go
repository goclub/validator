package vd_test

import (
	vd "github.com/goclub/validator"
	"log"
	"testing"
)
func Test_ExampleQuickStart(t *testing.T) {
	ExampleQuickStart()
}
func ExampleQuickStart() {
	checker := vd.NewCN()
	createUser := RequestCreateUser{
		Email: "xxx@domain.com",
		Name: "张三",
		Nickname: "三儿",
		Age: 20,
		Skills: []string{"clang", "go"},
		Address: RequestCreateUserAddress{
			Province: "上海",
			Detail:   "人民广场一号",
		},
	}
	report, err := checker.Check(createUser) ; if err != nil {
	    log.Print(err);return
	}
	if report.Fail {
		log.Print(report.Message)
	} else {
		log.Print("验证通过")
	}
}

type RequestCreateUser struct {
	Email string
	Name string
	Nickname string
	Age int
	Skills []string
	Address RequestCreateUserAddress
}
func (v RequestCreateUser) VD(r *vd.Rule) (err error) {
	r.String(v.Email, vd.StringSpec{
		Name:"邮箱地址",
		Ext: []vd.StringSpec{vd.Email()},
	})
	r.String(v.Name, vd.StringSpec{
		Name:              "姓名",
		MinRuneLen:        2,
		MaxRuneLen:        20,
	})
	r.String(v.Nickname, vd.StringSpec{
		Name:              "昵称",
		AllowEmpty:        true, // 昵称非必填
		BanPattern: []string{`\d`},
		PatternMessage: "昵称不允许包含数字",
		MinRuneLen:        2,
		MaxRuneLen:        10,
	})
	r.Int(v.Age, vd.IntSpec{
		Name:           "年龄",
		Min:            vd.Int(18),
		MinMessage:     "只允许成年人注册",
	})
	r.Slice(len(v.Skills), vd.SliceSpec{
		Name:          "技能",
		MaxLen:        vd.Int(10),
		MaxLenMessage: "最多填写{{MaxLen}}项",
		UniqueStrings: v.Skills,
	})
	for _, skill := range v.Skills {
		r.String(skill, vd.StringSpec{
			Name:              "技能项",
		})
	}
	// Address由 RequestCreateUserAddress{}.VD() 实现
	return nil
}
type RequestCreateUserAddress struct {
	Province string
	Detail string
}
func (v RequestCreateUserAddress) VD(r *vd.Rule) {
	r.String(v.Province, vd.StringSpec{
		Name:              "省",
	})
	r.String(v.Detail, vd.StringSpec{
		Name: "详细地址",
		Pattern: []string{`号`},
		PatternMessage: "地址必须包含门牌号",
	})
}
