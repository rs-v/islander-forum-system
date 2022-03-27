package controller

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
