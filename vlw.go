package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedVLW struct {
	s nmea.VLW
}

// implement nmea.Sentence functions
func (w WrappedVLW) String() string {
	return w.s.String()
}

func (w WrappedVLW) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedVLW) DataType() string {
	return w.s.DataType()
}

func (w WrappedVLW) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedVLW) GetLog() (float64, error) {
	return (unit.Length(w.s.TotalInWater) * unit.NauticalMile).Meters(), nil
}

func (w WrappedVLW) GetTripLog() (float64, error) {
	return (unit.Length(w.s.SinceResetInWater) * unit.NauticalMile).Meters(), nil
}
