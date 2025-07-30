package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedVHW struct {
	nmea.VHW
}

func NewVHW(s nmea.VHW) wrappedVHW {
	result := wrappedVHW{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedVHW) String() string {
	return w.VHW.String()
}

func (w wrappedVHW) Prefix() string {
	return w.VHW.Prefix()
}

func (w wrappedVHW) DataType() string {
	return w.VHW.DataType()
}

func (w wrappedVHW) TalkerID() string {
	return w.VHW.TalkerID()
}

// implement SignalK functions
func (w wrappedVHW) GetMagneticHeading() (float64, error) {
	return (unit.Angle(w.MagneticHeading) * unit.Degree).Radians(), nil
}

func (w wrappedVHW) GetTrueHeading() (float64, error) {
	return (unit.Angle(w.TrueHeading) * unit.Degree).Radians(), nil
}

func (w wrappedVHW) GetSpeedThroughWater() (float64, error) {
	if w.SpeedThroughWaterKPH == 0 && w.SpeedThroughWaterKnots > 0 {
		return (unit.Speed(w.SpeedThroughWaterKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return (unit.Speed(w.SpeedThroughWaterKPH) * unit.KilometersPerHour).MetersPerSecond(), nil
}
