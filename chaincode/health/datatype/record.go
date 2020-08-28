package datatype

import "time"

type RecordType int

const (
	SleepAtNoon = iota
	SleepAtNight
)

type HealthRecord struct {
	Type  RecordType
	Start time.Time
	End   time.Time
}
