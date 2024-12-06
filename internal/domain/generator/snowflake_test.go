package generator

import (
	"testing"
	"time"
)

func TestNewSnowflake(t *testing.T) {
	tests := []struct {
		name      string
		workerId  int64
		startTime int64
		wantErr   bool
	}{
		{
			name:      "valid worker id",
			workerId:  1,
			startTime: time.Now().UnixMilli(),
			wantErr:   false,
		},
		{
			name:      "invalid worker id - negative",
			workerId:  -1,
			startTime: time.Now().UnixMilli(),
			wantErr:   true,
		},
		{
			name:      "invalid worker id - too large",
			workerId:  1024, // 超过 workerMax
			startTime: time.Now().UnixMilli(),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf, err := NewSnowflake(tt.workerId, tt.startTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSnowflake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && sf == nil {
				t.Error("NewSnowflake() returned nil but wanted valid Snowflake")
			}
		})
	}
}

func TestSnowflake_NextID(t *testing.T) {
	startTime := time.Now().UnixMilli()
	sf, err := NewSnowflake(1, startTime)
	if err != nil {
		t.Fatalf("Failed to create Snowflake: %v", err)
	}

	// 生成多个ID并验证它们的唯一性
	idMap := make(map[int64]bool)
	for i := 0; i < 1000; i++ {
		id, err := sf.NextID()
		if err != nil {
			t.Errorf("NextID() error = %v", err)
			continue
		}

		if idMap[id] {
			t.Errorf("Duplicate ID generated: %d", id)
		}
		idMap[id] = true
	}
}

func TestSnowflake_Parse(t *testing.T) {
	startTime := time.Now().UnixMilli()
	workerId := int64(1)
	sf, _ := NewSnowflake(workerId, startTime)

	id, err := sf.NextID()
	if err != nil {
		t.Fatalf("Failed to generate ID: %v", err)
	}

	components, err := sf.Parse(id)
	if err != nil {
		t.Fatalf("Failed to parse ID: %v", err)
	}

	// 验证解析出的组件
	if components["workerId"] != workerId {
		t.Errorf("Parse() workerId = %v, want %v", components["workerId"], workerId)
	}

	timestamp := components["timestamp"]
	if timestamp < startTime {
		t.Errorf("Parse() timestamp = %v, should be >= %v", timestamp, startTime)
	}

	sequence := components["sequence"]
	if sequence < 0 || sequence > sequenceMask {
		t.Errorf("Parse() sequence = %v, should be between 0 and %v", sequence, sequenceMask)
	}
}
