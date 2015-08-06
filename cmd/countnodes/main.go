package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"

	"github.com/codegangsta/cli"
	"github.com/joshuarubin/countnodes"
	"github.com/joshuarubin/countnodes/node"
)

var (
	app  = cli.NewApp()
	tree *testNode
)

type testNode struct {
	left  *testNode
	right *testNode
}

func (n *testNode) Left() node.Node {
	if n.left == nil {
		return nil
	}
	return n.left
}

func (n *testNode) Right() node.Node {
	if n.right == nil {
		return nil
	}
	return n.right
}

func init() {
	app.Name = "countnodes"
	app.Version = "1.0.0"
	app.Usage = "countnodes"
	app.Authors = []cli.Author{
		{Name: "Joshua Rubin", Email: "joshua@rubixconsulting.com"},
	}
	app.Before = before
	app.Action = run
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "maxnodes, m",
			Usage: "maximum number of nodes in the generated tree",
			Value: 1024,
		},
	}
}

func before(c *cli.Context) error {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(c.Int("maxnodes"))))
	if err != nil {
		log.Fatal(err)
	}
	numNodes := num.Uint64()
	log.Printf("generating binary tree with %d nodes\n", numNodes)

	numNodes--
	tree = &testNode{}
	nodes := []*testNode{tree}

	for len(nodes) > 0 {
		var cur *testNode

		// shift off the first node
		cur, nodes = nodes[0], nodes[1:]

		if numNodes > 0 {
			n := &testNode{}
			nodes = append(nodes, n)
			cur.left = n
			numNodes--
		}

		if numNodes > 0 {
			n := &testNode{}
			nodes = append(nodes, n)
			cur.right = n
			numNodes--
		}
	}

	return nil
}

func run(c *cli.Context) {
	log.Println("counted", countnodes.CountNodes(tree))
}

func main() {
	app.Run(os.Args)
}
