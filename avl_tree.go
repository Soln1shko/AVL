package main

import "fmt"

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func (n *Node) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *Node) UpdateHeight() {
	leftHeight := 0
	rightHeight := 0

	if n.Left != nil {
		leftHeight = n.Left.GetHeight()
	}

	if n.Right != nil {
		rightHeight = n.Right.GetHeight()
	}

	n.Height = max(leftHeight, rightHeight) + 1
}

func (n *Node) BalanceFactor() int {
	if n == nil {
		return 0
	}

	leftHeight := 0
	rightHeight := 0

	if n.Left != nil {
		leftHeight = n.Left.GetHeight()
	}

	if n.Right != nil {
		rightHeight = n.Right.GetHeight()
	}

	return leftHeight - rightHeight
}

func (t *AVLTree) RotateRight(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.UpdateHeight()
	x.UpdateHeight()

	return x
}

func (t *AVLTree) RotateLeft(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.UpdateHeight()
	y.UpdateHeight()

	return y
}

func (t *AVLTree) Insert(key int) {
	t.Root = t.insert(t.Root, key)
}

func (t *AVLTree) insert(node *Node, key int) *Node {
	if node == nil {
		return &Node{Key: key, Height: 1}
	}

	if key < node.Key {
		node.Left = t.insert(node.Left, key)
	} else if key > node.Key {
		node.Right = t.insert(node.Right, key)
	} else {
		return node
	}

	node.UpdateHeight()
	balance := node.BalanceFactor()

	if balance > 1 && key < node.Left.Key {
		return t.RotateRight(node)
	}

	if balance < -1 && key > node.Right.Key {
		return t.RotateLeft(node)
	}

	if balance > 1 && key > node.Left.Key {
		node.Left = t.RotateLeft(node.Left)
		return t.RotateRight(node)
	}

	if balance < -1 && key < node.Right.Key {
		node.Right = t.RotateRight(node.Right)
		return t.RotateLeft(node)
	}

	return node
}

func (t *AVLTree) PrintInOrder() {
	printInOrder(t.Root)
	fmt.Println()
}

func printInOrder(node *Node) {
	if node != nil {
		printInOrder(node.Left)
		fmt.Printf("%d ", node.Key)
		printInOrder(node.Right)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tree := &AVLTree{}

	values := []int{10, 20, 30, 40, 50, 25}
	for _, val := range values {
		tree.Insert(val)
		fmt.Printf("Вставлено %d: ", val)
		tree.PrintInOrder()
	}
}


