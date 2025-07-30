package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedDBT struct {
	s nmea.DBT
}

// implement nmea.Sentence functions
func (w WrappedDBT) String() string {
	return w.s.String()
}

func (w WrappedDBT) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedDBT) DataType() string {
	return w.s.DataType()
}

func (w WrappedDBT) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedDBT) GetDepthBelowTransducer() (float64, error) {
	if v := w.s.DepthMeters; v > 0 {
		return v, nil
	}
	if v := w.s.DepthFeet; v > 0 {
		return (unit.Length(v) * unit.Foot).Meters(), nil
	}
	if v := w.s.DepthFathoms; v > 0 {
		return (unit.Length(v) * unit.Fathom).Meters(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
