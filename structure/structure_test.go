package structure

import (
	"fmt"
	"testing"

	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/stretchr/testify/assert"
)

func TestStructure_Register(t *testing.T) {
	s := NewStructure("test")
	tests := []struct {
		name        string
		toRegister  interface{}
		expectedErr error
	}{
		{
			name:        "test_performs_type_assertion_on_pipe",
			toRegister:  pipe.NewPipe("", nil),
			expectedErr: nil,
		},
		{
			name:        "test_falls_on_default_error",
			toRegister:  "",
			expectedErr: fmt.Errorf("provided interface cannot be cast to any known type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Register(tt.toRegister)
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
