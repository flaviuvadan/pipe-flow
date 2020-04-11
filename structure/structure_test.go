package structure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flaviuvadan/pipe-flow/sink"
	"github.com/flaviuvadan/pipe-flow/source"
)

func TestStructure_Register(t *testing.T) {
	testSource, _ := source.NewSource("test source", "test_file.csv", nil)
	testSink, _ := sink.NewSink("test_result.csv", nil)
	s := NewStructure("test")
	tests := []struct {
		name        string
		toRegister  interface{}
		expectedErr error
	}{
		{
			name:        "test_registers_source",
			toRegister:  testSource,
			expectedErr: nil,
		},
		{
			name:        "test_registers_sink",
			toRegister:  testSink,
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
