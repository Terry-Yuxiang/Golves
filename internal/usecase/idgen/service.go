package idgen

import (
	"context"
	"errors"
	"sync"

	"goloves/internal/domain/entity"
	"goloves/internal/domain/generator"
)

// Service errors
var (
	ErrInvalidGenerator  = errors.New("invalid generator type")
	ErrGeneratorNotFound = errors.New("generator not found")
)

// GeneratorType defines the type of ID generator
type GeneratorType string

const (
	TypeSnowflake GeneratorType = "snowflake"
	TypeSegment   GeneratorType = "segment"
)

// Service represents the ID generation service
type Service struct {
	mu         sync.RWMutex
	generators map[GeneratorType]generator.Generator
}

// NewService creates a new ID generation service
func NewService() *Service {
	return &Service{
		generators: make(map[GeneratorType]generator.Generator),
	}
}

// RegisterGenerator registers a new generator
func (s *Service) RegisterGenerator(typ GeneratorType, gen generator.Generator) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if gen == nil {
		return ErrInvalidGenerator
	}

	s.generators[typ] = gen
	return nil
}

// GenerateID generates a new ID using the specified generator
func (s *Service) GenerateID(ctx context.Context, typ GeneratorType) (*entity.ID, error) {
	s.mu.RLock()
	gen, exists := s.generators[typ]
	s.mu.RUnlock()

	if !exists {
		return nil, ErrGeneratorNotFound
	}

	id, err := gen.NextID()
	if err != nil {
		return nil, err
	}

	return &entity.ID{
		Value:  id,
		Source: string(typ),
	}, nil
}

// ParseID parses an ID to get its components
func (s *Service) ParseID(ctx context.Context, typ GeneratorType, id int64) (map[string]int64, error) {
	s.mu.RLock()
	gen, exists := s.generators[typ]
	s.mu.RUnlock()

	if !exists {
		return nil, ErrGeneratorNotFound
	}

	return gen.Parse(id)
}
