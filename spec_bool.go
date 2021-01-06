package vd
//
// type OptionBool struct {
// 	valid bool
// 	bool bool
// }
// func (o OptionBool) String() string {
// 	if !o.valid {return ""}
// 	if o.bool {return "true"} else { return "false"}
// }
// func (o OptionBool) Valid() bool { return o.valid }
// func (o OptionBool) Unwrap() bool {
// 	if !o.valid {panic("OptionBool: can not wrap invalid OptionBool")}
// 	return o.bool
// }
// func Bool(b bool) OptionBool {
// 	return OptionBool{true, b}
// }
type BoolSpec struct {
	Name string
}

func (r *Rule) Bool(v bool, spec BoolSpec) {

}
