package kdtree

import "fmt"

type Point interface {
	Dimension() float64   // returns the dimension of the point
	Dimensions(i int) int // returns the i-th dimension of the point
	String() string       // returns a string representation of the point
}

type KDTree struct {
	root *node
}

func NewKDTree(points []Point) *KDTree {
	return &KDTree{
		root: newKDTree(points, 0),
	}
}

func newKDTree(points []Point, axis int) *node {

}

func (t *KDTree) String() string {
	return fmt.Sprintf("[%s]", printNodeAndChildren(t.root))
}

func (t *KDTree) Insert(p Point) {

}

func (t *KDTree) Remove(p Point) {

}

// Balance rebalanced the tree
func (t *KDTree) Balance() {
	t.root = newKDTree()
}

func (t *KDTree) Points() []Point {
	if t.root == nil {
		return nil
	}
	return t.root.points()
}
