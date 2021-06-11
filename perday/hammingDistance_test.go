package perday

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	var demos = []struct {
		x      int
		y      int
		except int
	}{
		{1, 4, 2},
		{3, 1, 1},
	}

	for _, demo := range demos {
		actual := hammingDistance(demo.x, demo.y)
		assert.Equal(t, demo.except, actual)
	}
}
