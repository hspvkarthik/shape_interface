package shape_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t testing.T) {
	t.Run("check initialization of square with positive length", func(t *testing.T) {
		assert.IsTypef(t, Square{}, NewSquare(4.0))
	})

	t.Run("check initialization of square with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-3.6)
		})
	})
}
