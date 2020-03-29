// source package is responsible for holding the logic of the start state of the plumbing system
package source

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/flaviuvadan/pipe-flow/pipe"
)

// ColIndex defines the position of the column names in the slice given by the csv package after reading a CSV file
const ColIndex = 0

// Source represents the beginning state of a pipeline
type Source struct {
	description string                // description of the source
	filename    string                // filename to the CSV file to be read by the source, in the current working directory
	data        map[string][]float64  // mapping of CSV column titles to the column data
	pipes       map[string]*pipe.Pipe // mapping of CSV column titles to the pipes that will operate on the columns
}

// New returns a new instance of a Source
func NewSource(dsc, file string, pps map[string]*pipe.Pipe) (*Source, error) {
	s := &Source{
		description: dsc,
		filename:    file,
		pipes:       pps,
	}
	err := s.read()
	return s, err
}

// read reads in the CSV formatted file passed as filename to the Source initializer
func (s *Source) read() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get the current working directory")
	}
	f, err := os.Open(path.Join(cwd, s.filename))
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to open the file located at: %s", s.filename)
	}
	defer f.Close()

	r := csv.NewReader(f)
	content, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read the content of the file located at: %s", s.filename)
	}

	if len(content) == 0 {
		return fmt.Errorf("empty file provided")
	}

	cols := content[ColIndex]
	s.data = map[string][]float64{}
	for i, c := range cols {
		colData := make([]float64, len(content)-1)
		for j, r := range content[1:] {
			v, err := strconv.ParseFloat(r[i], 64)
			if err != nil {
				return fmt.Errorf("failed to parse row value to float64: %v", r[i])
			}
			colData[j] = v
		}
		s.data[c] = colData
	}
	return nil
}
