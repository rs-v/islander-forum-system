package controller

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 操作
type Operate struct {
	Id       int `json:"id"`
	Type     int `json:"type"`
	AgreeNum int `json:"agreeNum"`
	DenyNum  int `json:"denyNum"`
	Status   int `json:"status"`
}

// 操作的类型
type OperateType struct {
	Id      int
	Value   string
	Operate func(postId int, userIdArr []int)
}

// 字符串操作
// [operate 123 123] 或 [operate?123&123]
func strOperate(str string) string {
	// 通过[]查找操作
	reg := regexp.MustCompile(`\[(.*?)\]`)
	if reg == nil {
		return str
	}
	res := reg.FindAllStringSubmatch(str, -1)
	// 操作数大于10不允许操作
	if len(res) > 10 {
		return str
	}
	// 通过?分割操作和参数
	for _, value := range res {
		fmt.Println(value)
		strOperateCase(value[1])
	}
	// 通过&分割参数

	return str
}

// 字符串操作选择
func strOperateCase(str string) {
	split := strings.Split(str, " ")
	operate := split[0]
	param := split[1:]
	fmt.Println(operate, param)
	switch operate {
	case "roll":
		rollOperate(param)
	}
}

func rollOperate(param []string) {
	// 验参
	if len(param) != 2 {
		return
	}
	start, err := strconv.Atoi(param[0])
	if err != nil {
		return
	}
	end, err := strconv.Atoi(param[1])
	if err != nil {
		return
	}
	rand.Seed(time.Now().Unix())
	num := rand.Intn(end-start) + start
	fmt.Println(num)
}
