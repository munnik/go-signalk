package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedVWR struct {
	nmea.VWR
}

func NewVWR(s nmea.VWR) wrappedVWR {
	result := wrappedVWR{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedVWR) String() string {
	return w.VWR.String()
}

func (w wrappedVWR) Prefix() string {
	return w.VWR.Prefix()
}

func (w wrappedVWR) DataType() string {
	return w.VWR.DataType()
}

func (w wrappedVWR) TalkerID() string {
	return w.VWR.TalkerID()
}

// implement SignalK functions
func (w wrappedVWR) GetRelativeWindDirection() (float64, error) {
	if w.MeasuredDirectionBow == nmea.Left {
		return -(unit.Angle(w.MeasuredAngle) * unit.Degree).Radians(), nil
	}
	if w.MeasuredDirectionBow == nmea.Right {
		return (unit.Angle(w.MeasuredAngle) * unit.Degree).Radians(), nil
	}

	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedVWR) GetWindSpeed() (float64, error) {
	if w.SpeedMPS > 0 {
		return w.SpeedMPS, nil
	}
	if w.SpeedKPH > 0 {
		return (unit.Speed(w.SpeedKPH) * unit.KilometersPerHour).MetersPerSecond(), nil
	}
	if w.SpeedKnots > 0 {
		return (unit.Speed(w.SpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
