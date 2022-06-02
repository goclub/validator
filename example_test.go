package vd_test

import (
	"context"
	"fmt"
	vd "github.com/goclub/validator"
	"log"
	"testing"
)

func TestExampleQuickStart(t *testing.T) {
	ExampleChecker_Check()
}

func ExampleChecker_Check() {
	ctx := context.Background()
	_ = ctx
	err := func() (err error) {
		checker := vd.NewCN()
		createUser := RequestCreateUser{
			Email:    "xxx@domain.com",
			Name:     "张三",
			Nickname: "三儿",
			Age:      20,
			Skills:   []string{"clang", "go"},
			Address: RequestCreateUserAddress{
				Province: "上海",
				Detail:   "", //
			},
			AddressList: []RequestCreateUserAddress{
				{
					Province: "上海",
					Detail:   "人民广场一号",
				},
				{
					Province: "上海",
					Detail:   "", // 人民广场一号
				},
			},
		}
		report, err := checker.Check(createUser)
		if err != nil {
			return
		}
		if report.Fail {
			log.Print("fail")
			log.Print("path:", report.Path)
			log.Print("message:", report.Message)
		} else {
			log.Print("验证通过")
		}
		return
	}()
	if err != nil {
		log.Printf("%+v", err)
	}
}

type RequestCreateUser struct {
	Email       string
	Name        string
	Nickname    string
	Age         int
	Skills      []string
	Address     RequestCreateUserAddress `json:"address"`
	AddressList []RequestCreateUserAddress
}

func (v RequestCreateUser) VD(r *vd.Rule) (err error) {
	r.String(v.Email, vd.StringSpec{
		Name: "邮箱地址",
		Path: "email",
		Ext:  []vd.StringSpec{vd.ExtString{}.Email()},
	})
	r.String(v.Name, vd.StringSpec{
		Name:       "姓名",
		Path:       "name",
		MinRuneLen: 2,
		MaxRuneLen: 20,
	})
	r.String(v.Nickname, vd.StringSpec{
		Name:           "昵称",
		Path:           "nickname",
		AllowEmpty:     true, // 昵称非必填
		BanPattern:     []string{`\d`},
		PatternMessage: "昵称不允许包含数字",
		MinRuneLen:     2,
		MaxRuneLen:     10,
	})
	r.Int(v.Age, vd.IntSpec{
		Name:       "年龄",
		Path:       "age",
		Min:        vd.Int(18),
		MinMessage: "只允许成年人注册",
	})
	r.Slice(len(v.Skills), vd.SliceSpec{
		Name:          "技能",
		Path:          "skills",
		MaxLen:        vd.Int(10),
		MaxLenMessage: "最多填写{{MaxLen}}项",
		UniqueStrings: v.Skills,
	})
	for index, skill := range v.Skills {
		r.String(skill, vd.StringSpec{
			Name: "技能项",
			Path: vd.PathIndex("skill", index),
		})
	}
	// Address由 RequestCreateUserAddress{}.VD() 实现
	return nil
}

type RequestCreateUserAddress struct {
	Province string
	Detail   string
}

func (v RequestCreateUserAddress) VD(r *vd.Rule) (err error) {
	r.String(v.Province, vd.StringSpec{
		Path: "province",
		Name: "省",
	})
	r.String(v.Detail, vd.StringSpec{
		Name:           "详细地址",
		Path:           "detail",
		Pattern:        []string{`号`},
		PatternMessage: "地址必须包含门牌号,例如:某路110号",
	})
	return
}

type LogKind uint8

func (v LogKind) Validator() error {
	switch v {
	case 1:
	case 2:
	default:
		return fmt.Errorf("kind can not be %v", v)
	}
	return nil
}

type Log struct {
	Kind LogKind
}

func (v Log) VD(r *vd.Rule) {
	r.Validator(v.Kind, "类型格式错误", "")
}
