package signalk

import (
	"fmt"
	"time"

	"github.com/adrianmo/go-nmea"
)

type wrappedZDA struct {
	nmea.ZDA
}

func NewZDA(s nmea.ZDA) wrappedZDA {
	result := wrappedZDA{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedZDA) String() string {
	return w.ZDA.String()
}

func (w wrappedZDA) Prefix() string {
	return w.ZDA.Prefix()
}

func (w wrappedZDA) DataType() string {
	return w.ZDA.DataType()
}

func (w wrappedZDA) TalkerID() string {
	return w.ZDA.TalkerID()
}

// implement SignalK functions
func (w wrappedZDA) GetDateTime() (string, error) {
	if w.Time.Valid {
		return time.Date(
			int(w.Year),
			time.Month(w.Month),
			int(w.Day),
			w.Time.Hour+int(w.OffsetHours),
			w.Time.Minute+int(w.OffsetMinutes),
			w.Time.Second,
			w.Time.Millisecond*1000000,
			time.UTC,
		).UTC().Format(time.RFC3339Nano), nil
	}
	return "", fmt.Errorf("value is unavailable")
}
