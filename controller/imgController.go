package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
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

func PostImgUpload(r *http.Request) {
	// 获取文件
	file, header, _ := r.FormFile("file")
	buff := new(bytes.Buffer)
	w := multipart.NewWriter(buff)
	createFormFile, err := w.CreateFormFile("smfile", header.Filename)
	if err == nil {
		readAll, _ := ioutil.ReadAll(file)
		createFormFile.Write(readAll)
	}
	w.Close()
	// 写入文件
	req, _ := http.NewRequest(http.MethodPost, "https://sm.ms/api/v2/upload", buff)

	token, _ := GetImgToken()
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	// 转发文件
	res, err := client.Do(req)
	if err != nil {
		return
	}
	fmt.Println(res)
}
