package model

import (
	"fmt"
	"testing"
)

// func TestGetIndex(t *testing.T) {
// 	fmt.Println(GetForumPostIndex(1, 0, 10))
// }

// func TestGetList(t *testing.T) {
// 	fmt.Println(GetForumPostList(1, 0, 2))
// }

// func TestUpdate(t *testing.T) {
// 	// UpdateForumPostCount(1, int(time.Now().Unix()))
// }

// func TestUser(t *testing.T) {
// 	// fmt.Println(GetUserById(2), GetUserByToken("233"), GetUserArr([]int{1, 2, 3}))
// 	fmt.Println(GetUserArr([]int{1, 2, 3}))
// }

// func TestGetLastList(t *testing.T) {
// 	fmt.Println(GetLastPostList([]int{0, 1}, 2))
// }

func TestGetBuff(t *testing.T) {
	fmt.Println(GetForumPostIndexBuff(1, 0, 10))
}

func TestGetReplyBuff(t *testing.T) {
	fmt.Println(GetLastPostListBuff([]int{1, 18, 24, 28, 40, 43}, 5))
	rdb := newRdb()
	res, _ := rdb.Keys(ctx, "*").Result()
	fmt.Println(res)
}

func TestLastTime(t *testing.T) {
	fmt.Println(GetForumIndexLastTime(0, 10, []int{}))
}

func TestChangePost(t *testing.T) {
	ChangePostPlate(57, 1)
	ChangeFollowPostPlate(57, 1)
}
