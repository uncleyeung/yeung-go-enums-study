//性别枚举
package enum

type SexEnum int8

const (
	SexEnumMan SexEnum = iota
	SexEnumWoMan
)

//Values
func (c SexEnum) Values() []SexEnum {
	return initSexEnumArray()
}

//toString
func (c SexEnum) String() string {
	switch c {
	case SexEnumMan:
		return "男"
	case SexEnumWoMan:
		return "女"
	default:
		return "不知道"
	}
}

// init SexEnum
func initSexEnumArray() []SexEnum {
	return []SexEnum{SexEnumMan, SexEnumWoMan}
}
