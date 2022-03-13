package controller

import (
	"encoding/json"

	"github.com/forum_server/model"
)

func TransferForumPlateArrModel(data []model.ForumPlate) []ForumPlate {
	res := make([]ForumPlate, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = ForumPlate(data[i])
	}
	return res
}

func TransferForumPostList(data []ForumPost) []model.ForumPost {
	res := make([]model.ForumPost, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = transferForumPost(data[i])
	}
	return res
}

func TransferForumPostListModel(data []model.ForumPost) []ForumPost {
	res := make([]ForumPost, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = transferForumPostModel(data[i])
	}
	return res
}

func transferForumPost(data ForumPost) model.ForumPost {
	return model.ForumPost{
		Id:            data.Id,
		Title:         data.Title,
		Value:         data.Value,
		FollowId:      data.FollowId,
		PlateId:       data.PlateId,
		Status:        data.Status,
		ReplyArr:      intArr2json(data.ReplyArr),
		UserId:        data.UserId,
		Time:          data.Time,
		MediaUrl:      data.MediaUrl,
		ReplyCount:    data.ReplyCount,
		TopStatus:     data.TopStatus,
		LastReplyTime: data.LastReplyTime,
		SageAddId:     intArr2json(data.SageAddId),
		SageSubId:     intArr2json(data.SageSubId),
		Name:          data.Name,
	}
}

func transferForumPostModel(data model.ForumPost) ForumPost {
	res := ForumPost{
		Id:            data.Id,
		Title:         data.Title,
		Value:         data.Value,
		FollowId:      data.FollowId,
		PlateId:       data.PlateId,
		Status:        data.Status,
		ReplyArr:      json2intArr(data.ReplyArr),
		UserId:        data.UserId,
		Time:          data.Time,
		MediaUrl:      data.MediaUrl,
		ReplyCount:    data.ReplyCount,
		TopStatus:     data.TopStatus,
		LastReplyTime: data.LastReplyTime,
		SageAddId:     json2intArr(data.SageAddId),
		SageSubId:     json2intArr(data.SageSubId),
		Name:          data.Name,
	}
	res.SageAddCount = len(res.SageAddId)
	res.SageSubCount = len(res.SageSubId)
	return res
}

func intArr2json(data []int) string {
	if data == nil || len(data) == 0 {
		return "[]"
	}
	res, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func json2intArr(data string) []int {
	if data == "" {
		return nil
	}
	var res []int
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		panic(err)
	}
	return res
}
