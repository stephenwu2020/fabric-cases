package datatype

const (
	SleepAtNoon = iota
	SleepAtNight
)

type HealthRecord struct {
	Type  int
	Start int64
	End   int64
}
