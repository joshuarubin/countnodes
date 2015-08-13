package main

import (
	"container/list"
	"crypto/rand"
	"errors"
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

	if numNodes == 0 {
		return nil
	}

	numNodes--
	tree = &testNode{}
	nodes := list.New()
	nodes.PushBack(tree)

	for nodes.Front() != nil {
		cur := nodes.Front().Value.(*testNode)
		nodes.Remove(nodes.Front())

		for _, i := range node.Sides {
			if numNodes == 0 {
				break
			}

			n := &testNode{}
			nodes.PushBack(n)

			switch i {
			case node.LeftSide:
				cur.left = n
			case node.RightSide:
				cur.right = n
			default:
				return errors.New("invalid side")
			}

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
