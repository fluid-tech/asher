package api

// Base signature for the generators
// TODO: Add Build() to Generator in the later iterations
type Generator interface {
	String() string
}
