package points

import "fmt"

type Point3D struct {
	X, Y, Z float64
}

func (p *Point3D) Dimensions() int {
	return 3
}

func (p *Point3D) Dimension(i int) float64 {
	switch i {
	case 0:
		return p.X
	case 1:
		return p.Y
	case 2:
		return p.Z
	default: //should return an error?
		return 0
	}
}

func (p *Point3D) String() string {
	return fmt.Sprintf("{%.2f, %.2f, %.2f}", p.X, p.Y, p.Z)
}
