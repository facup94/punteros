package nullable

import (
	"time"
)

type Time struct {
	Value    time.Time
	HasValue bool
}

func FromTime(t time.Time) Time {
	return Time{
		Value:    t,
		HasValue: true,
	}
}

func NilTime() Time {
	return Time{
		Value:    time.Time{},
		HasValue: false,
	}
}

func (t Time) MarshalJSON() ([]byte, error) {
	return marshall(t.HasValue, t.Value)
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	return unmarshal(data, &t.HasValue, &t.Value)
}

func (t Time) String() string {
	return toString(t.HasValue, t.Value)
}
