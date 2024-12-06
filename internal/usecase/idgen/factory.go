package idgen

import (
    "time"
    "goloves/internal/domain/generator"
)

// Config represents the configuration for ID generators
type Config struct {
    Snowflake struct {
        WorkerID  int64
        StartTime time.Time
    }
    // 可以在这里添加 Segment 配置
}

// NewServiceWithConfig creates a new service with the provided configuration
func NewServiceWithConfig(cfg Config) (*Service, error) {
    service := NewService()

    // Initialize Snowflake generator
    snowflake, err := generator.NewSnowflake(cfg.Snowflake.WorkerID, cfg.Snowflake.StartTime.UnixMilli())
    if err != nil {
        return nil, err
    }
    
    err = service.RegisterGenerator(TypeSnowflake, snowflake)
    if err != nil {
        return nil, err
    }

    // TODO: Initialize Segment generator when implemented
    
    return service, nil
} 