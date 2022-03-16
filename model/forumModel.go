package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ForumPost struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Value         string `json:"value"`
	FollowId      int    `json:"followId"`
	PlateId       int    `json:"plateId"`
	Status        int    `json:"status"`
	ReplyArr      string `json:"replyArr"`
	UserId        int    `json:"userId"`
	Time          int    `json:"time"`
	MediaUrl      string `json:"mediaUrl"`
	ReplyCount    int    `json:"replyCount"`
	TopStatus     int    `json:"topStatus"`
	LastReplyTime int    `json:"listReplayTime"`
	SageAddId     string `json:"sageAddId"`
	SageSubId     string `json:"sageSubId"`
	Name          string `json:"name"`
}

type ForumPlate struct {
	Id     int
	Name   string
	Status int
}

func GetForumPlate() ([]ForumPlate, error) {
	var res []ForumPlate
	db := newDB()
	err := db.Where("status = ?", 0).Find(&res).Error
	return res, err
}

func GetForumPost(postId int) (ForumPost, error) {
	var res ForumPost
	db := newDB()
	err := db.Where("id = ?", postId).Take(&res).Error
	return res, err
}

// 增加buff版本
func GetForumPostIndex(plateId int, page int, size int) ([]ForumPost, int) {
	first := page * size
	var res []ForumPost
	var count int64
	db := newDB()
	db.Limit(size).Offset(first).Where("plate_id = ? and follow_id = 0 and status = 0", plateId).Order("last_reply_time desc").Find(&res).Limit(-1).Offset(-1).Count(&count)
	return res, int(count)
}

func GetForumPostList(postId int, page int, size int) ([]ForumPost, int) {
	first := page * size
	var res []ForumPost
	var count int64
	db := newDB()
	db.Limit(size).Offset(first).Where("(follow_id = ? and status = 0) or id = ?", postId, postId).Order("time asc").Find(&res).Limit(-1).Offset(-1).Count(&count)
	return res, int(count)
}

// 增加buff版本
func GetLastPostList(followIdArr []int, count int) []ForumPost {
	var res []ForumPost
	db := newDB()
	db.Raw("select fp.* from (select fp1.*, (select count(*) + 1 from forum_post fp2 where fp2.follow_id = fp1.follow_id and fp2.time > fp1.time) top from forum_post fp1 where follow_id in ?) fp where top < (? + 1) order by fp.follow_id, top", followIdArr, count).Scan(&res)
	return res
}

func GetAlreadySagePost(page int, size int) []ForumPost {
	first := page * size
	var res []ForumPost
	db := newDB()
	db.Limit(size).Offset(first).Where("status = 1").Order("time desc").Find(&res)
	return res
}

func GetAlreadySageCount() int {
	var count int64
	db := newDB()
	db.Model(&ForumPost{}).Where("status = 1").Count(&count)
	return int(count)
}

// 新增buff版本，更新
func SaveForumPost(post ForumPost) int {
	db := newDB()
	db.Create(&post)
	return post.Id
}

// 新增buff版本
func SaveForumReply(post ForumPost) int {
	db := newDB()
	db.Create(&post)
	return post.Id
}

func setCount(key string, count int) {
	rdb := newRdb()
	rdb.Set(ctx, key, count, time.Second*30)
}

func initForumIndexBuff(postArr []ForumPost, plateId int) {
	rdb := newRdb()
	key := "postIndex" + strconv.Itoa(plateId)
	rdb.Del(ctx, key).Result()
	for i := 0; i < len(postArr); i++ {
		setForumIndexBuff(postArr[i])
	}
}

func initForumReplyBuff(postArr []ForumPost, followId int) {
	rdb := newRdb()
	key := "postReply" + strconv.Itoa(followId)
	rdb.Del(ctx, key).Result()
	for i := 0; i < len(postArr); i++ {
		setForumReplyBuff(postArr[i])
	}
}

func setForumIndexBuff(post ForumPost) {
	rdb := newRdb()
	score := post.LastReplyTime
	post.LastReplyTime = 0
	key := "postIndex:" + strconv.Itoa(post.PlateId)
	AddZsetBuff(key, score, post)
	rdb.Expire(ctx, key, time.Second*30)
}

func setForumReplyBuff(post ForumPost) {
	rdb := newRdb()
	key := "postReply:" + strconv.Itoa(post.FollowId)
	AddZsetBuff(key, post.Time, post)
	rdb.Expire(ctx, key, time.Second*30)
}

func UpdateForumPostCount(postId int, time int) {
	db := newDB()
	db.Model(&ForumPost{}).Where("id = ?", postId).Updates(ForumPost{LastReplyTime: time})
	db.Model(&ForumPost{}).Where("id = ?", postId).UpdateColumn("reply_count", gorm.Expr("reply_count + ?", 1))
}

func UpdateSageAdd(post ForumPost) {
	db := newDB()
	db.Model(&post).Update("sage_add_id", post.SageAddId)
}

func UpdateSageSub(post ForumPost) {
	db := newDB()
	db.Model(&post).Update("sage_sub_id", post.SageSubId)
}

func UpdateForumPostStatus(post ForumPost, status int) {
	db := newDB()
	db.Model(&post).Update("status", status)
}

func (ForumPost) TableName() string {
	return "forum_post"
}

func (ForumPlate) TableName() string {
	return "forum_plate"
}
