package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeFromArray(t *testing.T) {
	mapData := [][]int{
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
		{2, 2, 2, 3, 4},
		{2, 2, 2, 3, 3},
	}

	expectedQuadtree := Quadtree{width: 5, height: 5, root: &node{
		topLeftX:       0,
		topLeftY:       0,
		width:          5,
		height:         5,
		content:        0,
		topLeftNode:    &node{content: 0},
		topRightNode:   &node{content: 1},
		bottomLeftNode: &node{content: 2},
		bottomRightNode: &node{
			topLeftX:        3,
			topLeftY:        3,
			width:           2,
			height:          2,
			content:         3,
			topLeftNode:     &node{content: 3},
			topRightNode:    &node{content: 4},
			bottomLeftNode:  &node{content: 3},
			bottomRightNode: &node{content: 3},
		},
	}}

	result := MakeFromArray(mapData)

	if !reflect.DeepEqual(result, expectedQuadtree) {
		t.Errorf("Le quadtree obtenu ne correspond pas au résultat attendu. Attendu : %v, Récupéré : %v", 1, 1)
		printNode(expectedQuadtree.root, 0)
		printNode(result.root, 0)
	}

}

func printNode(n *node, depth int) {
	if n == nil {
		return
	}

	indent := ""
	for i := 0; i < depth; i++ {
		indent += "\t"
	}

	fmt.Printf("%sContent: %d, X: %d, Y: %d, Width: %d, Height: %d\n", indent, n.content, n.topLeftX, n.topLeftY, n.width, n.height)

	printNode(n.topLeftNode, depth+1)
	printNode(n.topRightNode, depth+1)
	printNode(n.bottomLeftNode, depth+1)
	printNode(n.bottomRightNode, depth+1)
}
