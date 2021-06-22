package model

import baseModel "origin/router/model"

type User struct {
	baseModel.Model
	UserId  int    `json:"user_id"`
	UnionId string `json:"union_id"`
	Openid  string `json:"openid"`
}

func (u *User) TableName() string {
	return "user"
}
