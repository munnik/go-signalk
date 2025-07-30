package signalk

import (
	"github.com/adrianmo/go-nmea"
)

type WrappedGSA struct {
	s nmea.GSA
}

// implement nmea.Sentence functions
func (w WrappedGSA) String() string {
	return w.s.String()
}

func (w WrappedGSA) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedGSA) DataType() string {
	return w.s.DataType()
}

func (w WrappedGSA) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedGSA) GetNumberOfSatellites() (int64, error) {
	return int64(len(w.s.SV)), nil
}

func (w WrappedGSA) GetFixType() (string, error) {
	return w.s.FixType, nil
}
