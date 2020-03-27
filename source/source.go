// source package is responsible for holding the logic of the start state of the plumbing system
package source

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"

	"github.com/flaviuvadan/pipe-flow/pipe"
	"strconv"
)

// COL_INDEX defines the position of the column names in the slice given by the csv package after reading a CSV file
const COL_INDEX = 0

// Source represents the beginning state of a pipeline
type Source struct {
	description string                // description of the source
	path        string                // path to the CSV file to be read by the source
	data        map[string][]float64  // mapping of CSV column titles to the column data
	pipes       map[string]*pipe.Pipe // mapping of CSV column titles to the pipes that will operate on the columns
}

// New returns a new instance of a Source
func NewSource(dsc, pth string, pps map[string]*pipe.Pipe) (*Source, error) {
	s := &Source{
		description: dsc,
		path:        pth,
		pipes:       pps,
	}
	err := s.read()
	return s, err
}

// read reads in the CSV formatted file passed as path to the Source initializer
func (s *Source) read() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get the current working directory")
	}
	f, err := os.Open(path.Join(cwd, s.path))
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to open the file located at: %s", s.path)
	}
	defer f.Close()

	r := csv.NewReader(f)
	content, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read the content of the file located at: %s", s.path)
	}

	if len(content) == 0 {
		return fmt.Errorf("empty file provided")
	}

	cols := content[COL_INDEX]
	s.data = map[string][]float64{}
	for _, c := range cols {
		colData := make([]float64, len(content)-1)
		for i, r := range content[1:] {
			v, err := strconv.ParseFloat(r[i], 64)
			if err != nil {
				return fmt.Errorf("failed to parse row value to float64: %v", r[i])
			}
			colData[i] = v
		}
		s.data[c] = colData
	}
	return nil
}
