// 枚举尝试
package main

import (
	"encoding/json"
	"fmt"
	"uncle.yeung.com/enums/bo"
	"uncle.yeung.com/enums/enum"
)

func main() {
	user := bo.User{
		Age:      12,
		Name:     "yanbo",
		Phone:    "17688533666",
		SexEnum:  enum.SexEnumWoMan,
		WorkEnum: enum.WorkEnumOnTheJob,
	}
	user.InitData()
	fmt.Println("1基础数据:", user)

	jsonValue, _ := json.Marshal(user)
	fmt.Println("2生成json: ", string(jsonValue))

	json2Obj := bo.User{}
	_ = json.Unmarshal(jsonValue, &json2Obj)
	fmt.Println("3json2Obj数据:", json2Obj)

	jsonStr := "{\"Id\":123456,\"CreamTime\":\"2020-09-20T13:28:23.4511096+08:00\",\"UpdateTime\":\"2020-09-20T13:28:23.4511096+08:00\",\"Name\":\"yanbo\",\"Age\":12,\"Phone\":\"17688533666\",\"SexEnum\":1,\"WorkEnum\":3}"

	json2Obj1 := bo.User{}
	err := json.Unmarshal([]byte(jsonStr), &json2Obj1)
	fmt.Println(err)
	fmt.Println("4json2Obj数据:", json2Obj1)

	values := user.SexEnum.Values()
	for _, v := range values {
		fmt.Println("枚举数据:", v)
		fmt.Println("枚举数据:", v.String())

	}

	values1 := user.WorkEnum.Values()
	for _, v := range values1 {
		fmt.Println("枚举数据:", v)
		fmt.Println("枚举数据:", v.String())

	}
}
