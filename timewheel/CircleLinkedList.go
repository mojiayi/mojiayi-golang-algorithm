package timewheel

import (
	"errors"
	"fmt"
)

type CircleLinkedList struct {
	Head *SimpleTaskNode
	Tail *SimpleTaskNode
	Size int
}

func (c *CircleLinkedList) AddToHead(node *SimpleTaskNode) {
	if c.Size == 0 {
		node.Next = node
		c.Tail = node
	} else {
		c.Tail.Next = node
		node.Next = c.Head
	}
	c.Head = node
	c.Size++
}

func (c *CircleLinkedList) AddToTail(node *SimpleTaskNode) {
	if c.Size == 0 {
		node.Next = node
		c.Tail = node
		c.Head = node
	} else {
		c.Tail.Next = node
		node.Next = c.Head
		c.Tail = node
	}
	c.Size++
}

func (c *CircleLinkedList) DeleteHead() (bool, error) {
	if c.Size > 1 {
		c.Head = c.Head.Next
		c.Tail.Next = c.Head
		c.Size--
		return true, nil
	} else if c.Size == 1 {
		c.Head = nil
		c.Tail = nil
		c.Size--
		return true, nil
	}
	return false, errors.New("链表中已经没有元素")
}

func (c *CircleLinkedList) DeleteTail() (bool, error) {
	if c.Size > 1 {
		node := c.Head
		for node.Next.ID != c.Tail.ID {
			node = node.Next
		}
		node.Next = c.Head
		c.Size--
		return true, nil
	} else if c.Size == 1 {
		c.Head = nil
		c.Tail = nil
		c.Size--
		return true, nil
	}
	return false, errors.New("链表中已经没有元素")
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
