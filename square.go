package shape_interface

import "math"

type Square struct {
	side float64
}

func NewSquare(side float64) Square {
	if side <= 0 {
		panic("side length of square must be positive")
	}
	return Square{side: side}
}

func (sq Square) FindPerimeter() float64 {
	if sq.side <= 0 {
		panic("can't calculate perimeter when side is negative")
	} else if sq.side == 1 {
		return 4.0
	}
	return 4 * sq.side
}

func (sq Square) FindArea() float64 {
	if sq.side <= 0 {
		panic("can't calculate area when side is negative")
	} else if sq.side == 1 {
		return 1.0
	}
	return math.Pow(sq.side, 2)
}
