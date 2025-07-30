package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedMWD struct {
	s nmea.MWD
}

// implement nmea.Sentence functions
func (w WrappedMWD) String() string {
	return w.s.String()
}

func (w WrappedMWD) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedMWD) DataType() string {
	return w.s.DataType()
}

func (w WrappedMWD) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedMWD) GetTrueWindDirection() (float64, error) {
	return (unit.Angle(w.s.WindDirectionTrue) * unit.Degree).Radians(), nil
}

func (w WrappedMWD) GetMagneticWindDirection() (float64, error) {
	return (unit.Angle(w.s.WindDirectionMagnetic) * unit.Degree).Radians(), nil
}

func (w WrappedMWD) GetWindSpeed() (float64, error) {
	if w.s.WindSpeedMeters > 0 {
		return w.s.WindSpeedMeters, nil
	}
	return (unit.Speed(w.s.WindSpeedKnots) * unit.Knot).MetersPerSecond(), nil
}
