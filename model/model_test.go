package model

import (
	"fmt"
	"testing"
)

func TestGetIndex(t *testing.T) {
	fmt.Println(GetForumPostIndex(1, 0, 10))
}

func TestGetList(t *testing.T) {
	fmt.Println(GetForumPostList(1, 0, 2))
}

func TestUpdate(t *testing.T) {
	// UpdateForumPostCount(1, int(time.Now().Unix()))
}

func TestUser(t *testing.T) {
	// fmt.Println(GetUserById(2), GetUserByToken("233"), GetUserArr([]int{1, 2, 3}))
	fmt.Println(GetUserArr([]int{1, 2, 3}))
}

func TestGetLastList(t *testing.T) {
	fmt.Println(GetLastPostList([]int{0, 1}, 2))
}
