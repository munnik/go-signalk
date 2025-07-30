package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedVTG struct {
	s nmea.VTG
}

// implement nmea.Sentence functions
func (w WrappedVTG) String() string {
	return w.s.String()
}

func (w WrappedVTG) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedVTG) DataType() string {
	return w.s.DataType()
}

func (w WrappedVTG) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedVTG) GetTrueCourseOverGround() (float64, error) {
	return (unit.Angle(w.s.TrueTrack) * unit.Degree).Radians(), nil
}

func (w WrappedVTG) GetMagneticCourseOverGround() (float64, error) {
	return (unit.Angle(w.s.MagneticTrack) * unit.Degree).Radians(), nil
}

func (w WrappedVTG) GetSpeedOverGround() (float64, error) {
	if w.s.GroundSpeedKPH == 0 && w.s.GroundSpeedKnots > 0 {
		return (unit.Speed(w.s.GroundSpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return (unit.Speed(w.s.GroundSpeedKPH) * unit.KilometersPerHour).MetersPerSecond(), nil
}
