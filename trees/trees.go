package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/* implementation of Tree

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	nodeWalk(t, ch)
	close(ch)
}

func nodeWalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		nodeWalk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		nodeWalk(t.Right, ch)
	}	
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for t1v := range ch1 {
		t2v := <-ch2

		if (t1v != t2v) { return false }	
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("Check trees for sameness")

	fmt.Printf("tree.New(1), tree.New(1): %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(2), tree.New(1): %v\n", Same(tree.New(2), tree.New(1)))
}
