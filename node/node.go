package node

// Side represents the left or right side child of a binary tree node
type Side int

const (
	// LeftSide represents the left child of a binary tree node
	LeftSide Side = iota
	// RightSide represents the right child of a binary tree node
	RightSide
)

var (
	// Sides is an array of all possible node sides
	Sides = []Side{LeftSide, RightSide}
)

// Node represents a node of a binary tree
type Node interface {
	Left() Node
	Right() Node
}
