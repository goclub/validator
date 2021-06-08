# validator

> 类型安全的 Go Struct 验证器

```go
import vd "github.com/goclub/validator"
```

go 主流的结构体校验器使用结构体标签配置规则，这种方式类型不安全，一旦结构体标签写错将无法通过编译期检查出来。
例如：

```go
type exampleStruct2 struct {
  Name  string `valid:"-"`
  Email string `valid:"email"` // 一旦单词写错只有运行期才能发现错误，且不容易记住各种语法
}
```

`goclub/validator` 提供类型安全高性能的验证器.

结构体 `CreateUser`

```go
type CreateUser struct {
	Email string
	NikeName string
	Age int
}
```
实现 `CreateUser.VD(r *vd.Rule)` 方法

```go
func (v RequestCreateUser) VD(r *vd.Rule) error {
	r.String(v.Email, vd.StringSpec{
		Name:"邮箱地址",
		Ext: []vd.StringSpec{vd.Email()},
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
	return nil
}
```

验证

```go
checker := vd.NewCN()
req := CreateUser{
    Email: "xxx@domain.com",
    Nickname: "三儿",
    Age: 20,
}
report := checker.Check(req)
if report.Fail {
    log.Print(report.Message)
} else {
    log.Print("验证通过")
}
```

初看会觉得 `goclub/validator` 的验证方式没有结构体标签验证便捷，用过一段时间后会体会到类型安全的好处。
不需要反复查阅文档确认结构体标签的规则是否正确，通过 查看 `vd.StringSpec` 结构体的字段即可知道有哪些配置方式。

更多示例：[example_test.go](https://github.com/goclub/validator/blob/main/example_test.go)