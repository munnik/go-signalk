package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedROT struct {
	s nmea.ROT
}

// implement nmea.Sentence functions
func (w WrappedROT) String() string {
	return w.s.String()
}

func (w WrappedROT) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedROT) DataType() string {
	return w.s.DataType()
}

func (w WrappedROT) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedROT) GetRateOfTurn() (float64, error) {
	if w.s.Valid {
		return (unit.Angle(w.s.RateOfTurn) * unit.Degree).Radians() / unit.Minute.Seconds(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
