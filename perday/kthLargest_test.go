package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	var demo = &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	actual := kthLargest(demo, 1)
	assert.Equal(t, 4, actual)
}
