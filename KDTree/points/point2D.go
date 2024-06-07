package points

import "fmt"

type Point2D struct {
	X, Y float64
}

func (p *Point2D) Dimensions() int {
	return 2
}

func (p *Point2D) Dimension(i int) float64 {
	if i == 0 {
		return p.X
	}
	return p.Y
}

func (p *Point2D) String() string {
	return fmt.Sprintf("{%.2f, %.2f}", p.X, p.Y)
}
