package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedMWV struct {
	nmea.MWV
}

func NewMWV(s nmea.MWV) wrappedMWV {
	result := wrappedMWV{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedMWV) String() string {
	return w.MWV.String()
}

func (w wrappedMWV) Prefix() string {
	return w.MWV.Prefix()
}

func (w wrappedMWV) DataType() string {
	return w.MWV.DataType()
}

func (w wrappedMWV) TalkerID() string {
	return w.MWV.TalkerID()
}

// implement SignalK functions
func (w wrappedMWV) GetTrueWindDirection() (float64, error) {
	if w.StatusValid && w.Reference == nmea.TheoreticalMWV {
		return (unit.Angle(w.WindAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedMWV) GetRelativeWindDirection() (float64, error) {
	if w.StatusValid && w.Reference == nmea.RelativeMWV {
		return (unit.Angle(w.WindAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedMWV) GetWindSpeed() (float64, error) {
	if w.StatusValid {
		switch w.WindSpeedUnit {
		case nmea.UnitKnotsMWV:
			return (unit.Speed(w.WindSpeed) * unit.Knot).MetersPerSecond(), nil
		case nmea.UnitSMilesHMWV:
			return (unit.Speed(w.WindSpeed) * unit.MilesPerHour).MetersPerSecond(), nil
		case nmea.UnitKMHMWV:
			return (unit.Speed(w.WindSpeed) * unit.KilometersPerHour).MetersPerSecond(), nil
		case nmea.UnitMSMWV:
			return w.WindSpeed, nil
		}
	}
	return 0, fmt.Errorf("value is unavailable")
}
