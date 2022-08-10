package shape_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Run("check initialization of square with positive length", func(t *testing.T) {
		assert.IsType(t, Square{}, NewSquare(4.0))
	})

	t.Run("check initialization of square with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-3.6)
		})
	})
}

func TestSquarePerimeter(t *testing.T) {
	t.Run("check perimeter of square when length is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-4.0).FindPerimeter()
		})
	})

	t.Run("check perimeter of square when length is 1", func(t *testing.T) {
		assert.Equal(t, 4.0, NewSquare(1.0).FindPerimeter())
	})

	t.Run("check perimeter of square value", func(t *testing.T) {
		assert.Equal(t, 12.8, NewSquare(3.2).FindPerimeter())
	})

	t.Run("check perimeter of square is positive", func(t *testing.T) {
		assert.Less(t, 0.0, NewSquare(5.6).FindPerimeter())
	})
}

func TestSquareArea(t *testing.T) {
	t.Run("check area of square when length is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-4.0).FindArea()
		})
	})

	t.Run("check area of square when length is 1", func(t *testing.T) {
		assert.Equal(t, 1.0, NewSquare(1).FindArea())
	})

	t.Run("check area of square value", func(t *testing.T) {
		assert.Equal(t, 6.25, NewSquare(2.5).FindArea())
	})

	t.Run("check area of square is positive", func(t *testing.T) {
		assert.Less(t, 0.0, NewSquare(5.6).FindArea())
	})
}
