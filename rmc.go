package signalk

import (
	"fmt"
	"time"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedRMC struct {
	nmea.RMC
}

// implement nmea.Sentence functions
func (w wrappedRMC) String() string {
	return w.RMC.String()
}

func (w wrappedRMC) Prefix() string {
	return w.RMC.Prefix()
}

func (w wrappedRMC) DataType() string {
	return w.RMC.DataType()
}

func (w wrappedRMC) TalkerID() string {
	return w.RMC.TalkerID()
}

func NewRMC(s nmea.RMC) wrappedRMC {
	result := wrappedRMC{s}
	return result
}

// implement SignalK functions
func (w wrappedRMC) GetMagneticVariation() (float64, error) {
	if w.Validity == nmea.ValidRMC {
		return (unit.Angle(w.Variation) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedRMC) GetTrueCourseOverGround() (float64, error) {
	if w.Validity == nmea.ValidRMC {
		return (unit.Angle(w.Course) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedRMC) GetPosition2D() (float64, float64, error) {
	if w.Validity == nmea.ValidRMC {
		return w.Latitude, w.Longitude, nil
	}
	return 0, 0, fmt.Errorf("value is unavailable")
}

func (w wrappedRMC) GetSpeedOverGround() (float64, error) {
	if w.Validity == nmea.ValidRMC {
		return (unit.Speed(w.Speed) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedRMC) GetDateTime() (string, error) {
	if w.Validity == nmea.ValidRMC {
		if w.Date.Valid && w.Time.Valid {
			return time.Date(
				w.Date.YY,
				time.Month(w.Date.MM),
				w.Date.DD,
				w.Time.Hour,
				w.Time.Minute,
				w.Time.Second,
				w.Time.Millisecond*1000000,
				time.UTC,
			).UTC().Format(time.RFC3339Nano), nil
		}
	}
	return "", fmt.Errorf("value is unavailable")
}
