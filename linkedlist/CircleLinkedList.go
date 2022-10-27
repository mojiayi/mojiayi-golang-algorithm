package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	ID   int
	Data interface{}
	Next *Node
}

type CircleLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (c *CircleLinkedList) AddToHead(newNode *Node) {
	if c.Size == 0 {
		newNode.Next = newNode
		c.Tail = newNode
	} else {
		c.Tail.Next = newNode
		newNode.Next = c.Head
	}
	c.Head = newNode
	c.Size++
}

func (c *CircleLinkedList) AddToTail(newNode *Node) {
	if c.Size == 0 {
		newNode.Next = newNode
		c.Tail = newNode
		c.Head = newNode
	} else {
		c.Tail.Next = newNode
		newNode.Next = c.Head
		c.Tail = newNode
	}
	c.Size++
}

func (c *CircleLinkedList) Add(newNode *Node, previousNode *Node) {
	if previousNode.ID == c.Tail.ID {
		c.AddToTail(newNode)
		return
	}
	newNode.Next = previousNode.Next
	previousNode.Next = newNode
	c.Size++
}

func (c *CircleLinkedList) DeleteHead() (bool, error) {
	if c.Size == 0 {
		return false, errors.New("链表中已经没有元素")
	}
	if c.Size > 1 {
		c.Head = c.Head.Next
		c.Tail.Next = c.Head
	} else {
		c.Head = nil
		c.Tail = nil
	}

	c.Size--
	return true, nil
}

func (c *CircleLinkedList) DeleteTail() (bool, error) {
	if c.Size == 0 {
		return false, errors.New("链表中已经没有元素")
	}
	if c.Size > 1 {
		node := c.Head
		for node.Next.ID != c.Tail.ID {
			node = node.Next
		}
		node.Next = c.Head
	} else if c.Size == 1 {
		c.Head = nil
		c.Tail = nil
	}
	c.Size--
	return true, nil
}

func (c *CircleLinkedList) Delete(node *Node) (bool, error) {
	if c.Size == 0 {
		return false, errors.New("链表中已经没有元素")
	}
	if node.ID == c.Head.ID {
		c.DeleteHead()
		return true, nil
	}
	if node.ID == c.Tail.ID {
		c.DeleteTail()
		return true, nil
	}
	var previousNode = *c.Head
	var nextNode = *c.Head.Next
	for nextNode.ID != c.Head.ID {
		if nextNode.ID == node.ID {
			break
		}
		previousNode = nextNode
		nextNode = *nextNode.Next
	}
	*previousNode.Next = *nextNode.Next
	c.Size--
	return true, nil
}

func (c *CircleLinkedList) Print() {
	node := c.Head
	for node.Next.ID != c.Head.ID {
		fmt.Print(node.ID)
		fmt.Print("->")
		node = node.Next
	}
	fmt.Print(node.ID)
	fmt.Print("->")
	fmt.Println(c.Head.ID)
}
