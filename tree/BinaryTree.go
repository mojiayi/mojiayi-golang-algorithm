package tree

import (
	"errors"
	"math/rand"
	"strconv"
)

/**
* 二叉树
 */
type BinaryTree struct {
	Root *BinaryTreeNode
}

/**
* 二叉树节点
 */
type BinaryTreeNode struct {
	ID    int
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

/**
* 随机生成整数类型的id
 */
func (treeNode *BinaryTreeNode) generateId() int {
	id := rand.Intn(9999)
	if id == 0 {
		return 9999
	}
	return id
}

/**
* 添加左子节点
 */
func (treeNode *BinaryTreeNode) AddToLeft(args interface{}) (*BinaryTreeNode, error) {
	if treeNode.Left != nil {
		return nil, errors.New("节点（id=" + strconv.Itoa(treeNode.ID) + "）已经有左子节点")
	}
	var leftNode = BinaryTreeNode{ID: treeNode.generateId(), Data: args}
	treeNode.Left = &leftNode
	return &leftNode, nil
}

/**
* 添加右子节点
 */
func (treeNode *BinaryTreeNode) AddToRight(args interface{}) (*BinaryTreeNode, error) {
	if treeNode.Right != nil {
		return nil, errors.New("节点（id=" + strconv.Itoa(treeNode.ID) + "）已经有右子节点")
	}
	var rightNode = BinaryTreeNode{ID: treeNode.generateId(), Data: args}
	treeNode.Right = &rightNode
	return &rightNode, nil
}

/**
* 前序遍历二叉树
 */
func (tree *BinaryTree) PreorderTraversal(handler func(args interface{})) {
	var preOrder func(treeNode *BinaryTreeNode)
	preOrder = func(treeNode *BinaryTreeNode) {
		if treeNode == nil {
			return
		}
		// 对当前节点进行业务处理
		handler(treeNode.Data)
		preOrder(treeNode.Left)
		preOrder(treeNode.Right)
	}

	preOrder(tree.Root)
}

/**
* 中序遍历二叉树
 */
func (tree *BinaryTree) MiddleOrderTraversal(handler func(args interface{})) {
	var middleOrder func(treeNode *BinaryTreeNode)
	middleOrder = func(treeNode *BinaryTreeNode) {
		if treeNode == nil {
			return
		}
		middleOrder(treeNode.Left)
		// 对当前节点进行业务处理
		handler(treeNode.Data)
		middleOrder(treeNode.Right)
	}

	middleOrder(tree.Root)
}

/**
* 后序遍历二叉树
 */
func (tree *BinaryTree) PostorderTraversal(handler func(args interface{})) {
	var postOrder func(treeNode *BinaryTreeNode)
	postOrder = func(treeNode *BinaryTreeNode) {
		if treeNode == nil {
			return
		}
		postOrder(treeNode.Left)
		postOrder(treeNode.Right)
		// 对当前节点进行业务处理
		handler(treeNode.Data)
	}

	postOrder(tree.Root)
}
