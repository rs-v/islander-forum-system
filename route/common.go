package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func remoteIp(req *http.Request) string {
	ipAddr := remoteIpItem(req)
	for i := 0; i < len(ipAddr); i++ {
		if ipAddr[i] == ':' {
			return ipAddr[:i]
		}
	}
	return ""
}

func remoteIpItem(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if remoteAddr != "" {
		return remoteAddr
	}
	remoteAddr = req.Header.Get("ipv4")
	if remoteAddr != "" {
		return remoteAddr
	}
	remoteAddr = req.Header.Get("XForwardedFor")
	if remoteAddr != "" {
		return remoteAddr
	}
	remoteAddr = req.Header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	remoteAddr = req.Header.Get("X-Real_Ip")
	if remoteAddr != "" {
		return remoteAddr
	} else {
		return "127.0.0.1"
	}
}

func write(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(data)
	enc.Encode(res)
}

func writeList(w http.ResponseWriter, list interface{}, count int) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := resList(list, count)
	enc.Encode(res)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.Encode(map[string]interface{}{"code": code, "msg": msg})
}

func res(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	if data != nil {
		res["data"] = data
	}
	res["msg"] = "ok"
	return res
}

func resList(list interface{}, count int) map[string]interface{} {
	data := make(map[string]interface{})
	data["list"] = list
	data["count"] = count
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	res["data"] = data
	res["msg"] = "ok"
	return res
}

func get(r *http.Request) map[string]string {
	var res = make(map[string]string)
	keys := r.URL.Query()
	for index, value := range keys {
		res[index] = value[0]
	}

	return res
}

func postJson(r *http.Request, obj interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// 重新写入
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, obj)
	// err := json.NewDecoder(r.Body).Decode(obj) // 会导致r.Body读取完后无法重新写入
	if err != nil {
		panic(err)
	}
}

func typeof(data interface{}) string {
	return reflect.TypeOf(data).String()
}

func struct2map(input interface{}) map[string]interface{} {
	data, _ := json.Marshal(&input)
	res := make(map[string]interface{})
	json.Unmarshal(data, &res)
	return res
}

func structList2map(input interface{}) []map[string]interface{} {
	data, _ := json.Marshal(&input)
	res := make([]map[string]interface{}, 0)
	json.Unmarshal(data, &res)
	return res
}
