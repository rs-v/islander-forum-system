package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/forum_server/config"
)

// 使用https://sm.ms图床
func GetImgToken() (string, error) {
	config := config.GetConfig()
	res, err := http.PostForm(
		"https://sm.ms/api/v2/token",
		url.Values{
			"username": {config.UserName},
			"password": {config.PassWord},
		})
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var resJson map[string]interface{}
	json.Unmarshal(body, &resJson)
	if resJson["success"] != true {
		return "", errors.New("get token field")
	}
	token := resJson["data"].(map[string]interface{})["token"].(string)
	return token, nil
}
