// structure package is responsible for holding the whole pipeline system (plumbing) and coordinating actions of
// starting and stopping
package structure

// Structure represents the state of the data processing system
type Structure struct {
}

// New returns a new instance of a Structure
func New() *Structure {
	return &Structure{}
}
