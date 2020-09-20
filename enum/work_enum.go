//工作枚举
package enum

type WorkEnum int8

const (
	WorkEnumOnTheJob      WorkEnum = iota //在职
	WorkEnumResign                        //离职
	WorkEnumAlreadyOffer                  //已有offer
	WorkEnumLookingForJob                 //正在找工作

)

//Values
func (c WorkEnum) Values() []WorkEnum {
	return initWorkEnumArray()
}

//toString
func (c WorkEnum) String() string {
	switch c {
	case WorkEnumOnTheJob:
		return "在职"
	case WorkEnumResign:
		return "离职"
	case WorkEnumAlreadyOffer:
		return "已有offer"
	case WorkEnumLookingForJob:
		return "正在找工作"
	default:
		return "不知道"
	}
}

// init WorkEnum
func initWorkEnumArray() []WorkEnum {
	return []WorkEnum{WorkEnumOnTheJob, WorkEnumResign, WorkEnumAlreadyOffer, WorkEnumLookingForJob}
}
