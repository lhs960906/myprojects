package main

import (
	"github.com/lhs960906/myprojects/bean/identity"
	"github.com/lhs960906/myprojects/bean/user"
	"github.com/lhs960906/myprojects/utils"
)

func main() {
	u1 := user.User{
		Id:       "20230709140000",
		Email:    "592081451@qq.com",
		Password: "lhs960906",
		Identity: identity.Identity{
			Id:   "362202199609061531",
			Name: "赖宏松",
		},
	}
	utils.PrintUserIdentity(u1)
}
