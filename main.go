package main

import (
	"fmt"

	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/flaviuvadan/pipe-flow/sink"
	"github.com/flaviuvadan/pipe-flow/source"
	"github.com/flaviuvadan/pipe-flow/structure"
)

func main() {
	pipeA := pipe.NewSingleOpsPipe("column_a_pipe", []func(v float64) (float64, error){
		func(v float64) (float64, error) {
			return v + 1, nil
		},
	})
	pipeB := pipe.NewSingleOpsPipe("column_b_pipe", []func(v float64) (float64, error){
		func(v float64) (float64, error) {
			return v + 1, nil
		},
	})
	pipeC := pipe.NewSingleOpsPipe("column_c_pipe", []func(v float64) (float64, error){
		func(v float64) (float64, error) {
			return v + 1, nil
		},
	})
	pipes := map[string]*pipe.Pipe{
		"a": pipeA,
		"b": pipeB,
		"c": pipeC,
	}
	src, err := source.NewSource("source_of_test_data", "test.csv", pipes)
	if err != nil {
		panic("failed to create source for a, b, c pipes")
	}
	snk, err := sink.NewSink("out.csv", []*pipe.Pipe{pipeA, pipeB, pipeC})
	if err != nil {
		panic("failed to create sink for a, b, c pipes")
	}
	stc := structure.NewStructure("structure_for_test_data_pipeline")
	if err := stc.Register(src); err != nil {
		panic("failed to add source to structure")
	}
	if err := stc.Register(snk); err != nil {
		panic("failed to add sink to structure")
	}
	if d, err := stc.Flow(); err != nil {
		panic("structure failed to flow")
	} else {
		fmt.Printf("Pipe structure done in: %v\n", d)
	}
}
