package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	{
		lhs := Point{28, 42}
		rhs := Point{12, 23}

		add := lhs.Add(rhs)
		assert.Equal(t, 40, add.X)
		assert.Equal(t, 65, add.Y)

		sub := lhs.Sub(rhs)
		assert.Equal(t, 16, sub.X)
		assert.Equal(t, 19, sub.Y)
	}
	{
		p := Point{28, 42}
		p.AddAssign(Point{3, 5})
		p.SubAssign(Point{8, 19})
		assert.Equal(t, 23, p.X)
		assert.Equal(t, 28, p.Y)
	}
}
