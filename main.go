package main

import (
	"fmt"
	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/flaviuvadan/pipe-flow/sink"
	"github.com/flaviuvadan/pipe-flow/source"
	"github.com/flaviuvadan/pipe-flow/structure"
)

func main() {
	pipeA := pipe.NewPipe("column_a_pipe", func(v float64) (float64, error) {
		return v + 1, nil
	})
	pipeB := pipe.NewPipe("column_b_pipe", func(v float64) (float64, error) {
		return v + 1, nil
	})
	pipeC := pipe.NewPipe("column_c_pipe", func(v float64) (float64, error) {
		return v + 1, nil
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
	stc := structure.NewStructure("structure_for_test_data_pipeline", src, snk)
	for k, v := range pipes {
		if err := stc.Register(v); err != nil {
			panic(fmt.Sprintf("failed to register pipe for column %v", k))
		}
	}
	if err := stc.Flow(); err != nil {
		panic("structure failed to flow")
	}
}
