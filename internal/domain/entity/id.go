package entity

// ID represents a distributed unique identifier
type ID struct {
	Value  int64
	Source string // 标识ID的生成来源（snowflake/segment）
}
