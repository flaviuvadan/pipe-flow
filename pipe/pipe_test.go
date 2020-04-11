package pipe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSingleOpsPipe_Flow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		pipeOps     []func(float64) (float64, error)
		pipeIn      map[string][]float64
		pipeOut     map[string][]float64
		expectedErr error
	}{
		{
			name: "test_cannot_flow_nil_through_op",
			pipeOps: []func(v float64) (float64, error){
				func(v float64) (float64, error) {
					return v + 1, nil
				},
			},
			pipeIn:      nil,
			pipeOut:     nil,
			expectedErr: fmt.Errorf("cannot flow nil input through specified singleOps"),
		},
		{
			name:    "test_cannot_flow_input_through_nil_op",
			pipeOps: nil,
			pipeIn: map[string][]float64{
				"a": {1.0},
			},
			pipeOut:     nil,
			expectedErr: fmt.Errorf("cannot flow input through nil singleOps"),
		},
		{
			name: "test_flows_single_val_through_op",
			pipeOps: []func(v float64) (float64, error){
				func(v float64) (float64, error) {
					return v + 1, nil
				},
			},
			pipeIn: map[string][]float64{
				"a": {1.0},
			},
			pipeOut: map[string][]float64{
				"a": {2.0},
			},
			expectedErr: nil,
		},
		{
			name: "test_flows_multi_vals_through_op",
			pipeOps: []func(v float64) (float64, error){
				func(v float64) (float64, error) {
					return v + 1, nil
				},
			},
			pipeIn: map[string][]float64{
				"a": {1.0, 2.0, 3.0},
				"b": {4.0, 5.0, 6.0},
			},
			pipeOut: map[string][]float64{
				"a": {2.0, 3.0, 4.0},
				"b": {5.0, 6.0, 7.0},
			},
			expectedErr: nil,
		},
		{
			name: "test_throws_err_on_op_err",
			pipeOps: []func(v float64) (float64, error){
				func(v float64) (float64, error) {
					return 0, fmt.Errorf("failed")
				},
			},
			pipeIn: map[string][]float64{
				"a": {1.0},
			},
			pipeOut: map[string][]float64{
				"a": {0.0},
			},
			expectedErr: fmt.Errorf("failed to apply op to val 1 on row 0 with op msg: failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewSingleOpsPipe(tt.name, tt.pipeOps)
			p.SetInput(tt.pipeIn)
			err := p.Flow()
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
			o := p.GetOutput()
			for k := range o {
				assert.ElementsMatch(t, tt.pipeOut[k], o[k])
			}
		})
	}
}

func TestNewAggregateOpPipe_Flow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		pipeOp      func([]float64) (float64, error)
		pipeIn      map[string][]float64
		pipeOut     map[string][]float64
		expectedErr error
	}{
		{
			name:   "test_returns_nil_on_empty_op",
			pipeOp: nil,
			pipeIn: map[string][]float64{
				"a": {1, 1, 1},
			},
			pipeOut:     nil,
			expectedErr: nil,
		},
		{
			name: "test_return_aggregate_data",
			pipeOp: func(values []float64) (float64, error) {
				agg := 0.0
				for _, v := range values {
					agg += v
				}
				return agg, nil
			},
			pipeIn: map[string][]float64{
				"a": {1, 1, 1},
			},
			pipeOut: map[string][]float64{
				"a": {3},
			},
			expectedErr: nil,
		},
		{
			name: "test_returns_err_on_aggregate_func_error",
			pipeOp: func(values []float64) (float64, error) {
				return 0, fmt.Errorf("test error")
			},
			pipeIn: map[string][]float64{
				"a": {1, 1, 1},
			},
			expectedErr: fmt.Errorf("failed to perform aggregate op on col (a), err: test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewAggregateOpPipe(tt.name, tt.pipeOp)
			p.SetInput(tt.pipeIn)
			err := p.Flow()
			fmt.Printf("err: %v\n", err)
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				o := p.GetOutput()
				for k := range o {
					assert.ElementsMatch(t, tt.pipeOut[k], o[k])
				}
			}
		})
	}
}
