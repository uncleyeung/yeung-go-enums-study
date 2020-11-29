package bo

import (
	"fmt"
	"github.com/uncleyeung/enums/bo/base"
	"github.com/uncleyeung/enums/enum"
)

type User struct {
	//继承
	base.Base
	Name     string
	Age      int8
	Phone    string
	SexEnum  enum.SexEnum
	WorkEnum enum.WorkEnum
}

func (c User) String() string {
	return fmt.Sprintf("[id:%d, Name:%s, Age:%d, Sex:%s, UpdateTime:%d, Phone:%s,Work:%s]", c.Id, c.Name, c.Age, c.SexEnum, c.UpdateTime, c.Phone, c.WorkEnum)

}
