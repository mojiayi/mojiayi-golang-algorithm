package algorithm_test

import (
	"mojiayi-golang-algorithm/tree"
	"strconv"
	"testing"
)

func TestDepthFirstTraversal(t *testing.T) {
	var multipleTree = buildMultipleTree()
	var exceptedResult []string = []string{"k", "l", "m", "h", "j", "i", "e", "f", "b", "g", "c", "n", "d", "a"}
	var index = 0
	multipleTree.DepthFirstTraversal(func(args interface{}) {
		if args.(string) != exceptedResult[index] {
			t.Error("第" + strconv.Itoa(index) + "个应该是" + exceptedResult[index])
		}

		index++
	})
}

func TestBreadthFirstTraversal(t *testing.T) {
	var multipleTree = buildMultipleTree()
	var exceptedResult []string = []string{"a", "b", "c", "d", "e", "f", "g", "n", "h", "i", "k", "l", "m", "j"}
	var index = 0
	multipleTree.BreadthFirstTraversal(func(args interface{}) {
		if args.(string) != exceptedResult[index] {
			t.Error("第" + strconv.Itoa(index) + "个应该是" + exceptedResult[index])
		}

		index++
	})
}

func buildMultipleTree() *tree.MultipleTree {
	var multipleTree = new(tree.MultipleTree)
	multipleTree.Root = &tree.MultipleTreeNode{ID: 0, Data: "a"}

	nodeB, _ := multipleTree.Root.AddChild("b")
	nodeC, _ := multipleTree.Root.AddChild("c")
	nodeD, _ := multipleTree.Root.AddChild("d")
	nodeD.AddChild("n")

	nodeE, _ := nodeB.AddChild("e")
	nodeB.AddChild("f")

	nodeC.AddChild("g")

	nodeH, _ := nodeE.AddChild("h")

	nodeI, _ := nodeE.AddChild("i")

	nodeI.AddChild("j")

	nodeH.AddChild("k")
	nodeH.AddChild("l")
	nodeH.AddChild("m")

	return multipleTree
}
