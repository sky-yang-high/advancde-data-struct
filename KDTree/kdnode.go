package kdtree

import "fmt"

type node struct {
	Point
	left  *node
	right *node
}

func printNodeAndChildren(n *node) string {
	if n == nil {
		return ""
	}
	if n != nil && (n.left != nil || n.right != nil) {
		return fmt.Sprintf("{%s %s %s}", printNodeAndChildren(n.left), n.String(), printNodeAndChildren(n.right))
	}
	return n.String()
}

func (n *node) String() string {
	return n.Point.String()
}

func (n *node) points() []Point {
	var points []Point
	if n.left != nil {
		points = append(points, n.left.points()...)
	}
	points = append(points, n.Point)
	if n.right != nil {
		points = append(points, n.right.points()...)
	}
	return points
}

func (n *node) insert(p Point, axis int) {

}

func (n *node) remove(p Point) {

}
