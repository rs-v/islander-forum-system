package controller

import (
	"fmt"
	"testing"
)

func TestGetForumIndex(t *testing.T) {
	fmt.Println(GetForumPostIndex(1, 0, 10))
}

// func TestGetForumPost(t *testing.T) {
// 	ReplyForumPost("test", 1, nil, 1, "", "name")
// 	PostForumPost("test", "test", nil, 1, 1, "", "name")
// }

func TestSage(t *testing.T) {
	// fmt.Println(SageAdd(1, 3))
}

func TestForum(t *testing.T) {
	fmt.Println(GetForumPlate())
}

func TestGetLast(t *testing.T) {
	fmt.Println(GetForumIndexLastTime(0, 10))
}
