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
	Param []TreeNode
}

func NewTreeNode(atom string, value Value, param []TreeNode) *TreeNode {
	return &TreeNode{
		Atom:  atom,
		Value: value,
		Param: param,
	}
}

type stack struct {
	Type  string // 符号类型
	Index int    // 位置索引
}

func eval(expression string) {
	var wSwap []byte
	var pSwap []string
	for i := 1; i < len(expression)-1; i++ {
		// 空格分割
		if expression[i] != ' ' {
			for {
				wSwap = append(wSwap, expression[i])
				i += 1
				if expression[i] == ' ' || i >= len(expression)-1 {
					pSwap = append(pSwap, string(wSwap))
					wSwap = nil
					break
				}
			}
		}
		// "号分割
		// \号分割
	}
	fmt.Println(pSwap)
}

// 回查找的字符串，和游标index
func parseValue(str string, index int) (TreeNode, int) {
	// 跳过空格
	index = skipSpace(str, index)
	fmt.Println(string(str[index]), index)
	// 表达式
	if str[index] == '[' {
		return parseExpression(str, index)
	} else if str[index] == '"' { // 字符串
		return TreeNode{}, 0
	} else { // 数字和原子
		return parseAtom(str, index)
	}
}

// 返回复合表达式
func parseExpression(str string, index int) (TreeNode, int) {
	index += 1
	rootStatus := false
	var node TreeNode
	for {
		index = skipSpace(str, index)
		if str[index] == ']' {
			return node, index
		} else if !rootStatus {
			node, index = parseValue(str, index)
		} else {
			childNode, i := parseValue(str, index)
			node.Param = append(node.Param, childNode)
			index = i
		}
	}
	// return node
}

// 返回原子表达式
func parseAtom(str string, index int) (TreeNode, int) {
	index = skipSpace(str, index)
	buff := make([]byte, 0)
	node := TreeNode{}
	for {
		if str[index] == ' ' || str[index] == ']' {
			node.Atom = string(buff)
			return node, index
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
