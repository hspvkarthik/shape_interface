package shape_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShapeInterface(t *testing.T) {
	t.Run("check square implements shape", func(t *testing.T) {
		assert.Implements(t, new(Shape), NewSquare(3.0))
	})

	t.Run("check rectangle implements shape", func(t *testing.T) {
		assert.Implements(t, new(Shape), NewRectangle(5.0, 6.0))
	})
}
