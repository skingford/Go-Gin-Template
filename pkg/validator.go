/*
 * @Author: kingford
 * @Date: 2023-03-11 01:17:38
 * @LastEditTime: 2023-03-11 01:18:05
 */
package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"gte=18"`
}

func NewValidate() {
	// 创建验证器
	validate := validator.New()

	// 定义待验证的数据
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       20,
	}

	// 验证数据
	err := validate.Struct(user)
	if err != nil {
		// 处理验证错误
		fmt.Println(err)
		return
	}

	// 执行正常逻辑
	fmt.Println("User is valid")
}
