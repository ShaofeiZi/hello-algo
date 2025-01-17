// File: binary_tree_dfs_test.go
// Created Time: 2022-11-26
// Author: Reanon (793584285@qq.com)

package chapter_tree

import (
	"fmt"
	"testing"

	. "github.com/krahets/hello-algo/pkg"
)

func TestPreInPostOrderTraversal(t *testing.T) {
	/* 初始化二叉树 */
	// 这里借助了一个从数组直接生成二叉树的函数
	root := ArrayToTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("初始化二叉树: ")
	PrintTree(root)

	// 前序遍历
	nums = nil
	preOrder(root)
	fmt.Println("前序遍历的结点打印序列 =", nums)

	// 中序遍历
	nums = nil
	inOrder(root)
	fmt.Println("中序遍历的结点打印序列 =", nums)

	// 后序遍历
	nums = nil
	postOrder(root)
	fmt.Println("后序遍历的结点打印序列 =", nums)
}
