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
