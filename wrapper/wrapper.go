package wrapper

import "strconv"

// Wrapper interface used to encapsulate all internal methods
type Wrapper interface {
	Atoi(s string) (int, error)
}

// wrapper implements Wrapper
type wrapper struct {}

// NewWrapper constructor to initialize Wrapper
func NewWrapper() Wrapper {
	return &wrapper{}
}

// Atoi wrapper of the internal Atoi method
func (w *wrapper) Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}
