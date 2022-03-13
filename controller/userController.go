package controller

import (
	"errors"

	"github.com/forum_server/model"
)

func GetUserById(userId int) model.User {
	return model.GetUserById(userId)
}

func GetUserByToken(token string) (model.User, error) {
	err := checkToken(token)
	if err != nil {
		return model.User{}, err
	}
	return model.GetUserByToken(token), err
}

func GetUserArr(idArr []int) []model.User {
	return model.GetUserArr(idArr)
}

func checkToken(token string) error {
	if len(token) != 32 {
		return errors.New("token is field")
	}
	return nil
}
