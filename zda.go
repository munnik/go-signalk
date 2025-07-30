package signalk

import (
	"fmt"
	"time"

	"github.com/adrianmo/go-nmea"
)

type WrappedZDA struct {
	s nmea.ZDA
}

// implement nmea.Sentence functions
func (w WrappedZDA) String() string {
	return w.s.String()
}

func (w WrappedZDA) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedZDA) DataType() string {
	return w.s.DataType()
}

func (w WrappedZDA) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedZDA) GetDateTime() (string, error) {
	if w.s.Time.Valid {
		return time.Date(
			int(w.s.Year),
			time.Month(w.s.Month),
			int(w.s.Day),
			w.s.Time.Hour+int(w.s.OffsetHours),
			w.s.Time.Minute+int(w.s.OffsetMinutes),
			w.s.Time.Second,
			w.s.Time.Millisecond*1000000,
			time.UTC,
		).UTC().Format(time.RFC3339Nano), nil
	}
	return "", fmt.Errorf("value is unavailable")
}
