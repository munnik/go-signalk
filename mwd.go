package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedMWD struct {
	nmea.MWD
}

func NewMWD(s nmea.MWD) wrappedMWD {
	result := wrappedMWD{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedMWD) String() string {
	return w.MWD.String()
}

func (w wrappedMWD) Prefix() string {
	return w.MWD.Prefix()
}

func (w wrappedMWD) DataType() string {
	return w.MWD.DataType()
}

func (w wrappedMWD) TalkerID() string {
	return w.MWD.TalkerID()
}

// implement SignalK functions
func (w wrappedMWD) GetTrueWindDirection() (float64, error) {
	return (unit.Angle(w.WindDirectionTrue) * unit.Degree).Radians(), nil
}

func (w wrappedMWD) GetMagneticWindDirection() (float64, error) {
	return (unit.Angle(w.WindDirectionMagnetic) * unit.Degree).Radians(), nil
}

func (w wrappedMWD) GetWindSpeed() (float64, error) {
	if w.WindSpeedMeters > 0 {
		return w.WindSpeedMeters, nil
	}
	return (unit.Speed(w.WindSpeedKnots) * unit.Knot).MetersPerSecond(), nil
}
