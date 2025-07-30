package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedVHW struct {
	s nmea.VHW
}

// implement nmea.Sentence functions
func (w WrappedVHW) String() string {
	return w.s.String()
}

func (w WrappedVHW) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedVHW) DataType() string {
	return w.s.DataType()
}

func (w WrappedVHW) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedVHW) GetMagneticHeading() (float64, error) {
	return (unit.Angle(w.s.MagneticHeading) * unit.Degree).Radians(), nil
}

func (w WrappedVHW) GetTrueHeading() (float64, error) {
	return (unit.Angle(w.s.TrueHeading) * unit.Degree).Radians(), nil
}

func (w WrappedVHW) GetSpeedThroughWater() (float64, error) {
	if w.s.SpeedThroughWaterKPH == 0 && w.s.SpeedThroughWaterKnots > 0 {
		return (unit.Speed(w.s.SpeedThroughWaterKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return (unit.Speed(w.s.SpeedThroughWaterKPH) * unit.KilometersPerHour).MetersPerSecond(), nil
}
