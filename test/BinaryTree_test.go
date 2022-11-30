package algorithm_test

import (
	"mojiayi-golang-algorithm/tree"
	"strconv"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	var binaryTree = buildBinaryTree()

	var exceptedResult []string = []string{"a", "b", "e", "h", "k", "l", "i", "m", "j", "f", "c", "d", "n", "g"}
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

	var exceptedResult []string = []string{"k", "h", "l", "e", "m", "i", "j", "b", "f", "a", "d", "n", "c", "g"}
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

	var exceptedResult []string = []string{"k", "l", "h", "m", "j", "i", "e", "f", "b", "n", "d", "g", "c", "a"}
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

	binaryTree.Root = &tree.BinaryTreeNode{ID: 0, Data: "a"}
	nodeB, _ := binaryTree.Root.AddToLeft("b")
	nodeC, _ := binaryTree.Root.AddToRight("c")

	nodeE, _ := nodeB.AddToLeft("e")
	nodeB.AddToRight("f")
	nodeD, _ := nodeC.AddToLeft("d")
	nodeC.AddToRight("g")

	nodeH, _ := nodeE.AddToLeft("h")
	nodeI, _ := nodeE.AddToRight("i")
	nodeD.AddToRight("n")

	nodeH.AddToLeft("k")
	nodeH.AddToRight("l")
	nodeI.AddToLeft("m")
	nodeI.AddToRight("j")

	return binaryTree
}
