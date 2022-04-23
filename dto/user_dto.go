package dto

import "web.com/ginGormJwt/model"

//定义需要返回前端的
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//定义转换的函数

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
