package sink

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flaviuvadan/pipe-flow/pipe"
)

func TestNewSink(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		pipes       []*pipe.Pipe
		expectedErr error
	}{
		{
			name:        "test_errors_on_nil_pipes",
			filename:    "",
			pipes:       nil,
			expectedErr: fmt.Errorf("cannot create a sink without pipes"),
		},
		{
			name:        "test_errors_on_no_pipes",
			filename:    "",
			pipes:       []*pipe.Pipe{},
			expectedErr: fmt.Errorf("cannot create a sink without pipes"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSink(tt.filename, tt.pipes)
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestSink_Collect(t *testing.T) {
	tests := []struct {
		name             string
		filename         string
		pipes            []*pipe.Pipe
		pipesOut         []map[string][]float64
		expectedSinkData map[string][]float64
	}{
		{
			name:     "test_collect_gets_single_data",
			filename: "",
			pipes: []*pipe.Pipe{
				pipe.NewPipe(nil),
			},
			pipesOut: []map[string][]float64{
				{
					"a": {1, 2, 3},
				},
			},
			expectedSinkData: map[string][]float64{
				"a": {1, 2, 3},
			},
		},
		{
			name:     "test_collect_gets_multi_data",
			filename: "",
			pipes: []*pipe.Pipe{
				pipe.NewPipe(nil),
				pipe.NewPipe(nil),
				pipe.NewPipe(nil),
			},
			pipesOut: []map[string][]float64{
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
			expectedSinkData: map[string][]float64{
				"a": {1, 2, 3},
				"b": {4, 5, 6},
				"c": {7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewSink(tt.filename, tt.pipes)
			for i, p := range s.pipes {
				p.SetOutput(tt.pipesOut[i])
			}
			s.Collect()
			for _, m := range tt.pipesOut {
				for k := range m {
					assert.EqualValues(t, m[k], tt.expectedSinkData[k])
				}
			}
		})
	}
}
