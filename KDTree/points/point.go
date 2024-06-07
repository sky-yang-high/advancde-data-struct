package points

import "fmt"

type Point struct {
	Coordinates []float64
	Data        interface{}
}

func NewPoint(coordinates []float64, data interface{}) *Point {
	return &Point{coordinates, data}
}

func (p *Point) Dimensions() int {
	return len(p.Coordinates)
}

func (p *Point) Dimension(i int) float64 {
	//should check if i is within range?
	return p.Coordinates[i]
}

func (p *Point) String() string {
	return fmt.Sprintf("{%v %v}", p.Coordinates, p.Data)
}
