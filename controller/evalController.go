package controller

import "fmt"

type Value struct {
	Num int
	Str string
}
type TreeNode struct {
	Oper  string
	Value Value
	Param []TreeNode
}

func NewTreeNode(oper string, value Value, param []TreeNode) *TreeNode {
	return &TreeNode{
		Oper:  oper,
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
