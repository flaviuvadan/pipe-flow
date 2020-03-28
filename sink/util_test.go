package sink

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		maps     []map[string][]float64
		expected map[string][]float64
	}{
		{
			name:     "test_returns_empty_map_on_nil_maps",
			maps:     nil,
			expected: map[string][]float64{},
		},
		{
			name:     "test_returns_empty_map_on_empty_maps",
			maps:     []map[string][]float64{},
			expected: map[string][]float64{},
		},
		{
			name: "test_returns_single_map_on_single_element_maps",
			maps: []map[string][]float64{
				{
					"a": {1},
				},
			},
			expected: map[string][]float64{
				"a": {1},
			},
		},
		{
			name: "test_returns_single_map_on_multi_elements_maps",
			maps: []map[string][]float64{
				{
					"a": {1, 2, 3},
				},
				{
					"b": {4, 5, 6},
				},
				{
					"c": {7, 8, 9},
				},
			},
			expected: map[string][]float64{
				"a": {1, 2, 3},
				"b": {4, 5, 6},
				"c": {7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Merge(tt.maps...)
			for k := range o {
				assert.ElementsMatch(t, o[k], tt.expected[k])
			}
		})
	}
}
