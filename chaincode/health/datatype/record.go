package datatype

type RecordType int

const (
	SleepAtNoon = iota
	SleepAtNight
)

type HealthRecord struct {
	Type  RecordType
	Start int64
	End   int64
}
