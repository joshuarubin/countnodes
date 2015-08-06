package countnodes

import (
	"github.com/joshuarubin/countnodes/node"
)

// CountNodes returns the number of nodes in the binary tree with root node
func CountNodes(root node.Node) uint32 {
	ret := uint32(0)
	return countNodes(root, &ret)
}

func countNodes(root node.Node, num *uint32) uint32 {
	if root == nil {
		return *num
	}

	(*num)++

	countNodes(root.Left(), num)
	countNodes(root.Right(), num)

	return *num
}
