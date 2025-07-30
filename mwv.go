package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedMWV struct {
	s nmea.MWV
}

// implement nmea.Sentence functions
func (w WrappedMWV) String() string {
	return w.s.String()
}

func (w WrappedMWV) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedMWV) DataType() string {
	return w.s.DataType()
}

func (w WrappedMWV) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedMWV) GetTrueWindDirection() (float64, error) {
	if w.s.StatusValid && w.s.Reference == nmea.TheoreticalMWV {
		return (unit.Angle(w.s.WindAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedMWV) GetRelativeWindDirection() (float64, error) {
	if w.s.StatusValid && w.s.Reference == nmea.RelativeMWV {
		return (unit.Angle(w.s.WindAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedMWV) GetWindSpeed() (float64, error) {
	if w.s.StatusValid {
		switch w.s.WindSpeedUnit {
		case nmea.UnitKnotsMWV:
			return (unit.Speed(w.s.WindSpeed) * unit.Knot).MetersPerSecond(), nil
		case nmea.UnitSMilesHMWV:
			return (unit.Speed(w.s.WindSpeed) * unit.MilesPerHour).MetersPerSecond(), nil
		case nmea.UnitKMHMWV:
			return (unit.Speed(w.s.WindSpeed) * unit.KilometersPerHour).MetersPerSecond(), nil
		case nmea.UnitMSMWV:
			return w.s.WindSpeed, nil
		}
	}
	return 0, fmt.Errorf("value is unavailable")
}
