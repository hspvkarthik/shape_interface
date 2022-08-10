package shape_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("check initialization of rectangle", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle(1.0, 2.0))
	})

	t.Run("check initialization of rectangle with positive sides", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewRectangle(1.0, 2.0)
		})
	})

	t.Run("check initialization of rectangle with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(-1.0, 2.0)
		})
	})

	t.Run("check initialization of rectangle with negative width", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(1.0, -2.0)
		})
	})
}

func TestFindPerimeter(t *testing.T) {

	t.Run("check perimeter when length is equal to 1", func(t *testing.T) {
		assert.Equal(t, 10.0, NewRectangle(1, 4.0).FindPerimeter())
	})

	t.Run("check perimeter when width is equal to 1", func(t *testing.T) {
		assert.Equal(t, 3.0, NewRectangle(0.5, 1).FindPerimeter())
	})

	t.Run("check perimeter when length is equal to width", func(t *testing.T) {
		assert.Equal(t, 6.4, NewRectangle(1.6, 1.6).FindPerimeter())
	})

	t.Run("check perimeter when length and width are not equal", func(t *testing.T) {
		assert.Equal(t, 5.6, NewRectangle(2.0, 0.8).FindPerimeter())
	})
}

func TestFindArea(t *testing.T) {
	t.Run("check area when length is equal to 1", func(t *testing.T) {
		assert.Equal(t, 5.6, NewRectangle(1, 5.6).FindArea())
	})

	t.Run("check area when width is equal to 1", func(t *testing.T) {
		assert.Equal(t, 6.4, NewRectangle(6.4, 1).FindArea())
	})

	t.Run("check area when length is equal to width", func(t *testing.T) {
		assert.Equal(t, 6.25, NewRectangle(2.5, 2.5).FindArea())
	})

	t.Run("check area when length and width are not equal", func(t *testing.T) {
		assert.Equal(t, 12.8, NewRectangle(2.0, 6.4).FindArea())
	})
}
