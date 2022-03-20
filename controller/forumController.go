package controller

import (
	"errors"
	"time"

	"github.com/forum_server/config"
	"github.com/forum_server/model"
)

type ForumPost struct {
	Id            int         `json:"id"`
	Title         string      `json:"title"`
	Value         string      `json:"value"`
	FollowId      int         `json:"followId"`
	PlateId       int         `json:"plateId"`
	Status        int         `json:"status"`
	ReplyArr      []int       `json:"replyArr"`
	UserId        int         `json:"userId"`
	Time          int         `json:"time"`
	MediaUrl      string      `json:"mediaUrl"`
	ReplyCount    int         `json:"replyCount"`
	TopStatus     int         `json:"topStatus"`
	LastReplyTime int         `json:"lastReplyTime"`
	SageAddId     []int       `json:"sageAddId"`
	SageSubId     []int       `json:"sageSubId"`
	SageAddCount  int         `json:"sageAddCount"`
	SageSubCount  int         `json:"sageSubCount"`
	Name          string      `json:"name"`
	LastReplyArr  []ForumPost `json:"lastReplyArr"`
}

type ForumPlate struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Value  string `json:"value"`
}

func NewForumPost(value string, plateId int, userId int) *ForumPost {
	return &ForumPost{
		Value:     value,
		PlateId:   plateId,
		UserId:    userId,
		Status:    0,
		Time:      int(time.Now().Unix()),
		SageAddId: make([]int, 0),
		SageSubId: make([]int, 0),
	}
}

func GetForumPlate() ([]ForumPlate, error) {
	res, err := model.GetForumPlate()
	if err != nil {
		return nil, err
	}
	return TransferForumPlateArrModel(res), nil
}

// 获取单个串
func GetForumPost(postId int) (ForumPost, error) {
	resModel, err := model.GetForumPost(postId)
	res := transferForumPostModel(resModel)
	if err != nil {
		return res, err
	}
	res.Name = model.GetUserById(res.UserId).Name
	return res, nil
}

// 获取板块首页
func GetForumPostIndex(plateId int, page int, size int) ([]ForumPost, int) {
	// 缓存版
	modelRes, count := model.GetForumPostIndexBuff(plateId, page, size)
	// 无缓存版
	// modelRes, count := model.GetForumPostIndex(plateId, page, size)
	res := TransferForumPostListModel(modelRes)
	// 获取最晚回复，最晚回复时间非自己
	// 用一次遍历获取所需内存分配
	replyIndexCount := 0
	for i := 0; i < len(res); i++ {
		if res[i].Time != res[i].LastReplyTime {
			replyIndexCount += 1
		}
	}
	followIdArr := make([]int, replyIndexCount)
	replyIndex := 0
	resMap := make(map[int]*ForumPost)
	for i := 0; i < len(res); i++ {
		if res[i].Time != res[i].LastReplyTime {
			followIdArr[replyIndex] = res[i].Id
			resMap[res[i].Id] = &res[i]
			replyIndex += 1
		}
	}
	// 缓存版
	lastRes := TransferForumPostListModel(model.GetLastPostListBuff(followIdArr, 5))
	// 无缓存版
	// lastRes := TransferForumPostListModel(model.GetLastPostList(followIdArr, 5))
	for i := len(lastRes) - 1; i > -1; i-- {
		resMap[lastRes[i].FollowId].LastReplyArr = append(resMap[lastRes[i].FollowId].LastReplyArr, lastRes[i])
	}
	return res, count
}

// 获取串页
func GetForumPostList(postId int, page int, size int) ([]ForumPost, int, error) {
	if _, err := GetForumPost(postId); err != nil {
		return nil, 0, err
	}
	modelRes, count := model.GetForumPostList(postId, page, size)
	res := TransferForumPostListModel(modelRes)
	return res, count, nil
}

// 发串
func PostForumPost(value string, title string, replyArr []int, plateId int, userId int, mediaUrl string, name string) error {
	if value == "" || len(value) > 1024 || len(title) > 128 || len(mediaUrl) > 2048 {
		return errors.New("input too long or value is null")
	}
	post := NewForumPost(value, plateId, userId)
	post.Title = title
	post.MediaUrl = mediaUrl
	post.LastReplyTime = post.Time
	post.Name = name
	if replyArr == nil {
		post.ReplyArr = make([]int, 0)
	} else {
		post.ReplyArr = replyArr
	}
	model.SaveForumPost(transferForumPost(*post))

	return nil
}

// 回串
func ReplyForumPost(value string, followId int, replyArr []int, userId int, mediaUrl string, name string) error {
	if value == "" || len(value) > 1024 || len(mediaUrl) > 2048 {
		return errors.New("input too long")
	}
	mainPost, err := GetForumPost(followId)
	if err != nil || mainPost.FollowId != 0 {
		return errors.New("replyId is illegal")
	}
	post := NewForumPost(value, mainPost.PlateId, userId)
	post.FollowId = mainPost.Id
	post.MediaUrl = mediaUrl
	post.LastReplyTime = post.Time
	post.Name = name
	if replyArr == nil {
		post.ReplyArr = make([]int, 0)
	} else {
		post.ReplyArr = replyArr
	}
	model.SaveForumReply(transferForumPost(*post))
	if mainPost.Status == 0 { // SAGE贴自沉
		model.UpdateForumPostCount(mainPost.Id, post.Time)
	} else {
		model.UpdateForumPostCount(mainPost.Id, mainPost.Time)
	}

	return nil
}

// sage添加
func SageAdd(postId int, userId int) error {
	post, err := GetForumPost(postId)
	if err != nil {
		return err
	}
	if post.Status != 0 {
		return errors.New("it's sage")
	}
	ok := findIntArr(post.SageAddId, userId)
	if !ok {
		post.SageAddId = append(post.SageAddId, userId)
		model.UpdateSageAdd(transferForumPost(post))
		SageSet(post)
	} else {
		// 取消sage
		return errors.New("you already sage it")
	}
	return nil
}

// 反sage添加
func SageSub(postId int, userId int) error {
	post, err := GetForumPost(postId)
	if err != nil {
		return err
	}
	ok := findIntArr(post.SageSubId, userId)
	if !ok {
		post.SageSubId = append(post.SageSubId, userId)
		model.UpdateSageSub(transferForumPost(post))
	} else {
		// 取消反sage
		return errors.New("you already unsage it")
	}
	return nil
}

func SageSet(post ForumPost) {
	conf := config.GetConfig()
	if len(post.SageAddId)-len(post.SageSubId) > conf.SageNum {
		model.UpdateForumPostStatus(transferForumPost(post), 1)
	}
}

func GetAlreadySagePost(page int, size int) []ForumPost {
	res := model.GetAlreadySagePost(page, size)
	return TransferForumPostListModel(res)
}

func ChangePostPlate() {

}

func findIntArr(arr []int, id int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == id {
			return true
		}
	}
	return false
}
