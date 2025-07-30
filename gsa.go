package signalk

import (
	"github.com/adrianmo/go-nmea"
)

type wrappedGSA struct {
	nmea.GSA
}

func NewGSA(s nmea.GSA) wrappedGSA {
	result := wrappedGSA{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedGSA) String() string {
	return w.GSA.String()
}

func (w wrappedGSA) Prefix() string {
	return w.GSA.Prefix()
}

func (w wrappedGSA) DataType() string {
	return w.GSA.DataType()
}

func (w wrappedGSA) TalkerID() string {
	return w.GSA.TalkerID()
}

// implement SignalK functions
func (w wrappedGSA) GetNumberOfSatellites() (int64, error) {
	return int64(len(w.SV)), nil
}

func (w wrappedGSA) GetFixType() (string, error) {
	return w.FixType, nil
}
