package controller

import (
	"fmt"
)

type Value struct {
	Num int
	Str string
}
type TreeNode struct {
	Atom  string
	Value Value
	Param []*TreeNode
}

func Eval(expression string) {
}

// 回查找的字符串，和游标index
func parseValue(str string, index int) (*TreeNode, int) {
	// 跳过空格
	index = skipSpace(str, index)
	fmt.Println(string(str[index]), index)
	// 表达式
	if str[index] == '[' {
		return parseExpression(str, index)
	} else if str[index] == '"' { // 字符串
		return parseStr(str, index)
	} else { // 数字和原子
		return parseAtom(str, index)
	}
}

// 返回复合表达式
func parseExpression(str string, index int) (*TreeNode, int) {
	index += 1
	rootStatus := false
	var node *TreeNode
	for {
		index = skipSpace(str, index)
		if str[index] == ']' {
			index += 1
			return node, index
		} else if !rootStatus {
			node, index = parseValue(str, index)
			rootStatus = true
		} else {
			childNode, i := parseValue(str, index)
			node.Param = append(node.Param, childNode)
			index = i
		}
	}
	// return node
}

// 返回原子表达式
func parseAtom(str string, index int) (*TreeNode, int) {
	index = skipSpace(str, index)
	buff := make([]byte, 0)
	node := TreeNode{}
	for {
		if str[index] == ' ' || str[index] == ']' {
			if str[index] == ' ' {
				index += 1 // 跳过最后的空格
			}
			node.Atom = string(buff)
			return &node, index
		}
		buff = append(buff, str[index])
		index += 1
	}
}

// 返回字符串表达式
func parseStr(str string, index int) (*TreeNode, int) {
	buff := make([]byte, 1)
	buff[0] = str[index]
	index += 1
	node := TreeNode{}
	for {
		if str[index] == '"' {
			buff = append(buff, str[index])
			node.Atom = string(buff)
			index += 1
			return &node, index
		}
		buff = append(buff, str[index])
		index += 1
	}
}

func skipSpace(str string, index int) int {
	// 跳过空格
	for {
		if str[index] == ' ' && index < len(str) {
			index += 1
		} else {
			break
		}
	}
	return index
}

func printTree(node *TreeNode) {
	fmt.Printf("%p", node)
	fmt.Println(node)
	for i := 0; i < len(node.Param); i++ {
		printTree(node.Param[i])
	}
}
