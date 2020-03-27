package source

import (
	"fmt"
	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSource(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		description string
		path        string
		pipes       map[string]*pipe.Pipe
		expectedErr error
	}{
		{
			name:        "test_source_reads_in_empty_csv",
			description: "",
			path:        "test_1.csv",
			pipes:       nil,
			expectedErr: nil,
		},
		{
			name:        "test_source_reads_in_csv_with_single_row",
			description: "",
			path:        "test_2.csv",
			pipes:       nil,
			expectedErr: nil,
		},
		{
			name:        "test_source_reads_in_csv_with_multiple_rows",
			description: "",
			path:        "test_3.csv",
			pipes:       nil,
			expectedErr: nil,
		},
		{
			name:        "test_returns_err_when_file_not_found",
			description: "",
			path:        "bad",
			pipes:       nil,
			expectedErr: fmt.Errorf("failed to open the file located at: bad"),
		},
		{
			name:        "test_source_returns_err_when_csv_not_valid",
			description: "",
			path:        "test_4.csv",
			pipes:       nil,
			expectedErr: fmt.Errorf("failed to read the content of the file located at: %s", "test_4.csv"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewSource(tt.description, tt.path, tt.pipes)
			assert.Equal(t, s.description, tt.description)
			assert.Equal(t, s.path, tt.path)
			assert.Equal(t, s.pipes, tt.pipes)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
