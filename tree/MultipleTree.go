package tree

import "math/rand"

/**
* 多叉树
 */
type MultipleTree struct {
	Root *MultipleTreeNode
}

/**
* 多叉树节点
 */
type MultipleTreeNode struct {
	ID       int
	Data     interface{}
	Children []*MultipleTreeNode
}

/**
* 随机生成整数类型的id
 */
func (treeNode *MultipleTreeNode) generateId() int {
	id := rand.Intn(9999)
	if id == 0 {
		return 9999
	}
	return id
}

/**
* 添加子节点
 */
func (treeNode *MultipleTreeNode) AddChild(args interface{}) (*MultipleTreeNode, error) {
	var children = treeNode.Children
	if len(children) == 0 {
		children = make([]*MultipleTreeNode, 0)
	}
	var child = MultipleTreeNode{ID: treeNode.generateId(), Data: args}
	treeNode.Children = append(children, &child)
	return &child, nil
}

/**
* 深度优先遍历多叉树
 */
func (tree *MultipleTree) DepthFirstTraversal(handler func(args interface{})) {
	var depthFirst func(treeNode *MultipleTreeNode)

	depthFirst = func(treeNode *MultipleTreeNode) {
		if treeNode == nil {
			return
		}
		for _, child := range treeNode.Children {
			depthFirst(child)
			handler(child.Data)
		}
	}

	depthFirst(tree.Root)
	handler(tree.Root.Data)
}

/**
* 广度优先遍历多叉树
 */
func (tree *MultipleTree) BreadthFirstTraversal(handler func(args interface{})) {
	var breadthFirst func(treeNode *MultipleTreeNode)

	breadthFirst = func(treeNode *MultipleTreeNode) {
		if treeNode == nil {
			return
		}
		var siblings = make([]*MultipleTreeNode, 0)
		for _, child := range treeNode.Children {
			handler(child.Data)
			siblings = append(siblings, child.Children...)
		}
		for _, child := range siblings {
			handler(child.Data)
		}
		for _, child := range siblings {
			breadthFirst(child)
		}
	}

	handler(tree.Root.Data)
	breadthFirst(tree.Root)
}
