package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	many := make([]int, 1001)
	for i := 0; i < len(many); i++ {
		many[i] = i
	}
	tests := []struct {
		name         string
		in           []int
		expectedOut  []int
		expectedSize int
	}{
		{
			name:         "test_empty_queue_is_empty",
			in:           []int{},
			expectedOut:  []int{},
			expectedSize: 0,
		},
		{
			name:         "test_push_one_elements",
			in:           []int{1},
			expectedOut:  []int{1},
			expectedSize: 1,
		},
		{
			name:         "test_push_few_elements",
			in:           []int{1, 2, 3, 4, 5},
			expectedOut:  []int{1, 2, 3, 4, 5},
			expectedSize: 5,
		},
		{
			name:         "test_push_many_elements",
			in:           many,
			expectedOut:  many,
			expectedSize: 1001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue(len(tt.in))
			for i := range tt.in {
				_ = q.Push(i)
			}
			assert.Equal(t, tt.expectedSize, q.Size())
			for i := range tt.in {
				o := q.Pop()
				assert.Equal(t, i, o)
			}
			assert.Equal(t, 0, q.Size())
		})
	}
}
