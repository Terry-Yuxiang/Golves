package generator

// Generator defines the interface for ID generation
type Generator interface {
	NextID() (int64, error)
	Parse(id int64) (map[string]int64, error)
}
