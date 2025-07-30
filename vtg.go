package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedVTG struct {
	nmea.VTG
}

func NewVTG(s nmea.VTG) wrappedVTG {
	result := wrappedVTG{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedVTG) String() string {
	return w.VTG.String()
}

func (w wrappedVTG) Prefix() string {
	return w.VTG.Prefix()
}

func (w wrappedVTG) DataType() string {
	return w.VTG.DataType()
}

func (w wrappedVTG) TalkerID() string {
	return w.VTG.TalkerID()
}

// implement SignalK functions
func (w wrappedVTG) GetTrueCourseOverGround() (float64, error) {
	return (unit.Angle(w.TrueTrack) * unit.Degree).Radians(), nil
}

func (w wrappedVTG) GetMagneticCourseOverGround() (float64, error) {
	return (unit.Angle(w.MagneticTrack) * unit.Degree).Radians(), nil
}

func (w wrappedVTG) GetSpeedOverGround() (float64, error) {
	if w.GroundSpeedKPH == 0 && w.GroundSpeedKnots > 0 {
		return (unit.Speed(w.GroundSpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return (unit.Speed(w.GroundSpeedKPH) * unit.KilometersPerHour).MetersPerSecond(), nil
}
