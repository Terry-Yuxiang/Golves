package idgen

import (
    "context"
    "testing"
    "time"
    
    "goloves/internal/domain/generator"
)

func TestService_RegisterGenerator(t *testing.T) {
    svc := NewService()

    // 创建一个有效的生成器
    sf, _ := generator.NewSnowflake(1, time.Now().UnixMilli())

    tests := []struct {
        name    string
        typ     GeneratorType
        gen     generator.Generator
        wantErr bool
    }{
        {
            name:    "valid generator",
            typ:     TypeSnowflake,
            gen:     sf,
            wantErr: false,
        },
        {
            name:    "nil generator",
            typ:     TypeSnowflake,
            gen:     nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := svc.RegisterGenerator(tt.typ, tt.gen)
            if (err != nil) != tt.wantErr {
                t.Errorf("RegisterGenerator() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestService_GenerateID(t *testing.T) {
    svc := NewService()
    sf, _ := generator.NewSnowflake(1, time.Now().UnixMilli())
    svc.RegisterGenerator(TypeSnowflake, sf)

    tests := []struct {
        name    string
        typ     GeneratorType
        wantErr bool
    }{
        {
            name:    "valid generator type",
            typ:     TypeSnowflake,
            wantErr: false,
        },
        {
            name:    "invalid generator type",
            typ:     "invalid",
            wantErr: true,
        },
    }

    ctx := context.Background()
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            id, err := svc.GenerateID(ctx, tt.typ)
            if (err != nil) != tt.wantErr {
                t.Errorf("GenerateID() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && id == nil {
                t.Error("GenerateID() returned nil but wanted valid ID")
            }
        })
    }
}

func TestService_ParseID(t *testing.T) {
    svc := NewService()
    sf, _ := generator.NewSnowflake(1, time.Now().UnixMilli())
    svc.RegisterGenerator(TypeSnowflake, sf)

    ctx := context.Background()
    id, _ := svc.GenerateID(ctx, TypeSnowflake)

    tests := []struct {
        name    string
        typ     GeneratorType
        id      int64
        wantErr bool
    }{
        {
            name:    "valid ID",
            typ:     TypeSnowflake,
            id:      id.Value,
            wantErr: false,
        },
        {
            name:    "invalid generator type",
            typ:     "invalid",
            id:      id.Value,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            components, err := svc.ParseID(ctx, tt.typ, tt.id)
            if (err != nil) != tt.wantErr {
                t.Errorf("ParseID() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && components == nil {
                t.Error("ParseID() returned nil but wanted valid components")
            }
        })
    }
} 