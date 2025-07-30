package signalk

import (
	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedVLW struct {
	nmea.VLW
}

func NewVLW(s nmea.VLW) wrappedVLW {
	result := wrappedVLW{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedVLW) String() string {
	return w.VLW.String()
}

func (w wrappedVLW) Prefix() string {
	return w.VLW.Prefix()
}

func (w wrappedVLW) DataType() string {
	return w.VLW.DataType()
}

func (w wrappedVLW) TalkerID() string {
	return w.VLW.TalkerID()
}

// implement SignalK functions
func (w wrappedVLW) GetLog() (float64, error) {
	return (unit.Length(w.TotalInWater) * unit.NauticalMile).Meters(), nil
}

func (w wrappedVLW) GetTripLog() (float64, error) {
	return (unit.Length(w.SinceResetInWater) * unit.NauticalMile).Meters(), nil
}
