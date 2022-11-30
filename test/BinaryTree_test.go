package algorithm_test

import (
	"mojiayi-golang-algorithm/tree"
	"strconv"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	var binaryTree = buildBinaryTree()

	var exceptedResult []string = []string{"m", "k", "h", "a", "b", "i", "c", "d", "l", "j", "e", "f"}
	var index = 0
	binaryTree.PreorderTraversal(func(args interface{}) {
		if args.(string) != exceptedResult[index] {
			t.Error("第" + strconv.Itoa(index) + "个应该是" + exceptedResult[index])
		}

		index++
	})
}

func TestMiddleOrderTraversal(t *testing.T) {
	var binaryTree = buildBinaryTree()

	var exceptedResult []string = []string{"a", "h", "b", "k", "c", "i", "d", "m", "l", "e", "j", "f"}
	var index = 0
	binaryTree.MiddleOrderTraversal(func(args interface{}) {
		if args.(string) != exceptedResult[index] {
			t.Error("第" + strconv.Itoa(index) + "个应该是" + exceptedResult[index])
		}

		index++
	})
}

func TestPostOrderTraversal(t *testing.T) {
	var binaryTree = buildBinaryTree()

	var exceptedResult []string = []string{"a", "b", "h", "c", "d", "i", "k", "e", "f", "j", "l", "m"}
	var index = 0
	binaryTree.PostorderTraversal(func(args interface{}) {
		if args.(string) != exceptedResult[index] {
			t.Error("第" + strconv.Itoa(index) + "个应该是" + exceptedResult[index])
		}

		index++
	})
}

func buildBinaryTree() *tree.BinaryTree {
	var binaryTree = new(tree.BinaryTree)

	binaryTree.Root = &tree.BinaryTreeNode{ID: 0, Data: "m"}
	var leftNode, _ = binaryTree.Root.AddToLeft("k")
	var rightNode, _ = leftNode.AddToRight("i")
	leftNode, _ = leftNode.AddToLeft("h")
	leftNode.AddToLeft("a")
	leftNode.AddToRight("b")
	rightNode.AddToLeft("c")
	rightNode.AddToRight("d")
	rightNode, _ = binaryTree.Root.AddToRight("l")
	rightNode, _ = rightNode.AddToRight("j")
	rightNode.AddToLeft("e")
	rightNode.AddToRight("f")

	return binaryTree
}
