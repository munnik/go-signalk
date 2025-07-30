package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedROT struct {
	nmea.ROT
}

func NewROT(s nmea.ROT) wrappedROT {
	result := wrappedROT{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedROT) String() string {
	return w.ROT.String()
}

func (w wrappedROT) Prefix() string {
	return w.ROT.Prefix()
}

func (w wrappedROT) DataType() string {
	return w.ROT.DataType()
}

func (w wrappedROT) TalkerID() string {
	return w.ROT.TalkerID()
}

// implement SignalK functions
func (w wrappedROT) GetRateOfTurn() (float64, error) {
	if w.Valid {
		return (unit.Angle(w.RateOfTurn) * unit.Degree).Radians() / unit.Minute.Seconds(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
