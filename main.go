package main

import (
	"github.com/lhs960906/myprojects/bean/identity"
	"github.com/lhs960906/myprojects/bean/user"
	"github.com/lhs960906/myprojects/utils"
)

func main() {
	u1 := user.User{
		Id:       "20230709140000",
		Email:    "592081451@163.com",
		Password: "123456",
		Identity: identity.Identity{
			Id:   "362201200009061532",
			Name: "小赖",
		},
	}
	utils.PrintUserIdentity(u1)
}
