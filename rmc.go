package signalk

import (
	"fmt"
	"time"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedRMC struct {
	s nmea.RMC
}

// implement nmea.Sentence functions
func (w WrappedRMC) String() string {
	return w.s.String()
}

func (w WrappedRMC) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedRMC) DataType() string {
	return w.s.DataType()
}

func (w WrappedRMC) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedRMC) GetMagneticVariation() (float64, error) {
	if w.s.Validity == nmea.ValidRMC {
		return (unit.Angle(w.s.Variation) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedRMC) GetTrueCourseOverGround() (float64, error) {
	if w.s.Validity == nmea.ValidRMC {
		return (unit.Angle(w.s.Course) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedRMC) GetPosition2D() (float64, float64, error) {
	if w.s.Validity == nmea.ValidRMC {
		return w.s.Latitude, w.s.Longitude, nil
	}
	return 0, 0, fmt.Errorf("value is unavailable")
}

func (w WrappedRMC) GetSpeedOverGround() (float64, error) {
	if w.s.Validity == nmea.ValidRMC {
		return (unit.Speed(w.s.Speed) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedRMC) GetDateTime() (string, error) {
	if w.s.Validity == nmea.ValidRMC {
		if w.s.Date.Valid && w.s.Time.Valid {
			return time.Date(
				w.s.Date.YY,
				time.Month(w.s.Date.MM),
				w.s.Date.DD,
				w.s.Time.Hour,
				w.s.Time.Minute,
				w.s.Time.Second,
				w.s.Time.Millisecond*1000000,
				time.UTC,
			).UTC().Format(time.RFC3339Nano), nil
		}
	}
	return "", fmt.Errorf("value is unavailable")
}
