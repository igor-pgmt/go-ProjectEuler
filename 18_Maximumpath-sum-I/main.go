// https://projecteuler.net/problem=18
// https://projecteuler.net/problem=67
// Maximum path sum
// Problem 18 & problem 67
package main

import (
	"fmt"
	"math"
)

type node struct {
	id, num, sum int
	left, right  *node
	child        int
}

type Tree []*node

var tree *Tree

func (t *Tree) insert(n *node) {
	*t = append(*t, n)
}

func (t *Tree) find(id int) *node {
	return (*t)[id]
}

func getParentIDs(id int) []int {
	line := int(math.Ceil((-1 + math.Sqrt(float64(1+8*(id+1)))) / 2))
	rightElement := line*(line+1)/2 - 1
	leftElement := rightElement - (line - 1)

	switch id {
	case leftElement:
		return []int{-1, id - (line - 1)}
	case rightElement:
		return []int{id - line, -1}
	default:
		return []int{id - line, id - (line - 1)}
	}
}

func createTree(numbers []int) *Tree {
	lines := int(math.Ceil((-1 + math.Sqrt(float64(1+8*len(numbers)))) / 2))
	tree = &Tree{}
	tree.insert(&node{0, numbers[0], numbers[0], nil, nil, -1})

	for i := 2; i <= lines; i++ {
		elements := i * (i + 1) / 2
		for j := elements - i; j < elements; j++ {
			newNode := &node{j, numbers[j], numbers[j], nil, nil, -1}
			tree.parentsInsert(newNode)
			tree.insert(newNode)
		}
	}
	return tree
}

func (tree Tree) parentsInsert(n *node) {
	parents := getParentIDs(n.id)
	if parents[0] != -1 {
		parent := tree.find(parents[0])
		parent.right = n
	}
	if parents[1] != -1 {
		parent := tree.find(parents[1])
		parent.left = n
	}
}

func (tree Tree) checkLines() {
	lines := int(math.Ceil((-1 + math.Sqrt(float64(1+8*len(tree)))) / 2))
	for line := lines; line > 0; line-- {
		tree.checkLine(line)
		tree.sumLines(line - 1)
	}
}

func (tree Tree) checkLine(line int) {
	elements := line * (line + 1) / 2
	for i := elements - line; i < elements-2; i++ {
		if tree[i].sum < tree[i-1].sum && tree[i].sum < tree[i+1].sum {
		}
	}
}

func (tree Tree) sumLines(line int) {
	elements := line * (line + 1) / 2
	for i := elements - line; i < elements; i++ {
		child, sum := getMaxNode(tree[i])
		tree[i].child = child
		tree[i].sum += sum
	}
}

func (tree Tree) showTree() {
	for _, v := range tree {
		fmt.Printf("%v,%v,%v,%v \n", v.id, v.num, v.sum, v.child)
	}
}

func (tree Tree) showPath() {
	var ans []int
	ans = append(ans, tree[0].id+1)
	currentNode := tree[0]
	line := int(math.Ceil((-1 + math.Sqrt(float64(1+8*len(tree)))) / 2))
	for i := 1; i < line; i++ {
		nodeID, _ := getMaxNode(currentNode)
		ans = append(ans, nodeID)
		currentNode = tree[nodeID]
	}
	fmt.Println("Path:", ans)
}

func getMaxNode(n *node) (int, int) {
	switch {
	case n.left.sum > n.right.sum:
		return n.left.id, n.left.sum
	default:
		return n.right.id, n.right.sum
	}
}

func main() {
	tree := createTree(numbers2)
	tree.checkLines()
	tree.showPath()
	tree.showTree()
	fmt.Println((*tree)[0].sum)
}
