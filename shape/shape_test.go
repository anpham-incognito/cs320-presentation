package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	var len float32 = 10
	s := new(Square)
	s.SetLength(len)
	assert.Equal(t, float32(50), s.GetPerimeter())
}
