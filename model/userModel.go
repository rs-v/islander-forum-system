package model

import (
	"log"
	"net/rpc"

	"github.com/forum_server/config"
)

type User struct {
	Id           int
	Name         string
	RegisterTime int
}

type UserQuery struct {
	UserId    int
	UserIdArr []int
	Token     string
}

func rpcClient() *rpc.Client {
	conf := config.GetConfig()
	client, err := rpc.DialHTTP("tcp", conf.RpcIp)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetUserById(userId int) User {
	query := &UserQuery{UserId: userId}
	client := rpcClient()
	res := new(User)
	err := client.Call("UserRpcServer.GetUserById", query, res)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func GetUserByToken(token string) User {
	query := &UserQuery{Token: token}
	client := rpcClient()
	res := new(User)
	err := client.Call("UserRpcServer.GetUserByToken", query, res)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func GetUserArr(idArr []int) []User {
	query := &UserQuery{UserIdArr: idArr}
	client := rpcClient()
	res := new([]User)
	err := client.Call("UserRpcServer.GetUserArr", query, res)
	if err != nil {
		log.Println(err)
	}
	return *res
}
