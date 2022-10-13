package timewheel

type SimpleTaskNode struct {
	ID   int
	Next *SimpleTaskNode
}
