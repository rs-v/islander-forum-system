package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	UserName    string
	PassWord    string
	Ip          string
	Database    string
	SageNum     int
	RpcIp       string
	RedisIp     string
	BuffTime    int
	ImgUserName string
	ImgPassWord string
}

func GetConfig() Config {
	// file, err := ioutil.ReadFile("./conf/config.json")
	file, err := ioutil.ReadFile("../conf/config.json")
	if err != nil {
		log.Println(err)
	}
	var res Config
	json.Unmarshal(file, &res)
	return res
}
