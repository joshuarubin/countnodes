package node

// Node represents a node of a binary tree
type Node interface {
	Left() Node
	Right() Node
}
