package sink

import (
	"fmt"
	"os"
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
				pipe.NewPipe("", nil, nil),
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
				pipe.NewPipe("", nil, nil),
				pipe.NewPipe("", nil, nil),
				pipe.NewPipe("", nil, nil),
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
			s.Collect()
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

func TestSink_Dump(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		pipes       []*pipe.Pipe
		pipesOut    []map[string][]float64
		expectedErr error
	}{
		{
			name:     "test_creates_empty_CSV_called_results",
			filename: "",
			pipes: []*pipe.Pipe{
				pipe.NewPipe("", nil, nil),
			},
			pipesOut:    []map[string][]float64{{}},
			expectedErr: nil,
		},
		{
			name:     "test_creates_empty_CSV_called_filename",
			filename: "test.csv",
			pipes: []*pipe.Pipe{
				pipe.NewPipe("", nil, nil),
			},
			pipesOut:    []map[string][]float64{{}},
			expectedErr: nil,
		},
		{
			name:     "test_creates_CSV_with_expected_content",
			filename: "test_1.csv",
			pipes: []*pipe.Pipe{
				pipe.NewPipe("", nil, nil),
			},
			pipesOut: []map[string][]float64{
				{
					"a": {1, 2, 3},
				},
			},
		},
		{
			name:     "test_creates_CSV_with_expected_extended_content",
			filename: "test_2.csv",
			pipes: []*pipe.Pipe{
				pipe.NewPipe("", nil, nil),
				pipe.NewPipe("", nil, nil),
				pipe.NewPipe("", nil, nil),
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, p := range tt.pipes {
				p.SetOutput(tt.pipesOut[i])
			}
			s, _ := NewSink(tt.filename, tt.pipes)
			err := s.Dump()
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
			// TODO: this can benefit from making the source.read function a util one
			// TODO: so other packages can read CSV files and check for equality between
			// TODO: their maps and the maps built from reading the CSV
			if tt.filename == "" {
				if err := os.Remove("results.csv"); err != nil {
					panic(fmt.Errorf("could not remove results.csv for tests teardown"))
				}
			} else {
				if err := os.Remove(tt.filename); err != nil {
					panic(fmt.Errorf("could not remove %v for tests teardown", tt.filename))
				}
			}
		})
	}
}
