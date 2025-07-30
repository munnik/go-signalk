package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedVWR struct {
	s nmea.VWR
}

// implement nmea.Sentence functions
func (w WrappedVWR) String() string {
	return w.s.String()
}

func (w WrappedVWR) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedVWR) DataType() string {
	return w.s.DataType()
}

func (w WrappedVWR) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedVWR) GetRelativeWindDirection() (float64, error) {
	if w.s.MeasuredDirectionBow == nmea.Left {
		return -(unit.Angle(w.s.MeasuredAngle) * unit.Degree).Radians(), nil
	}
	if w.s.MeasuredDirectionBow == nmea.Right {
		return (unit.Angle(w.s.MeasuredAngle) * unit.Degree).Radians(), nil
	}

	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedVWR) GetWindSpeed() (float64, error) {
	if w.s.SpeedMPS > 0 {
		return w.s.SpeedMPS, nil
	}
	if w.s.SpeedKPH > 0 {
		return (unit.Speed(w.s.SpeedKPH) * unit.KilometersPerHour).MetersPerSecond(), nil
	}
	if w.s.SpeedKnots > 0 {
		return (unit.Speed(w.s.SpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
