package generator

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits   = 10
	sequenceBits = 12
	workerMax    = -1 ^ (-1 << workerBits)
	sequenceMask = -1 ^ (-1 << sequenceBits)
	timeShift    = workerBits + sequenceBits
	workerShift  = sequenceBits
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	sequence  int64
	startTime int64
}

func NewSnowflake(workerId int64, startTime int64) (*Snowflake, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("worker ID excess of quantity")
	}

	return &Snowflake{
		timestamp: 0,
		workerId:  workerId,
		sequence:  0,
		startTime: startTime,
	}, nil
}

func (s *Snowflake) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	id := ((now - s.startTime) << timeShift) |
		(s.workerId << workerShift) |
		s.sequence

	return id, nil
}

func (s *Snowflake) Parse(id int64) (map[string]int64, error) {
	timestamp := (id >> timeShift) + s.startTime
	workerId := (id >> workerShift) & workerMax
	sequence := id & sequenceMask

	return map[string]int64{
		"timestamp": timestamp,
		"workerId":  workerId,
		"sequence":  sequence,
	}, nil
}
