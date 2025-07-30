package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedDBT struct {
	nmea.DBT
}

func NewDBT(s nmea.DBT) wrappedDBT {
	result := wrappedDBT{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedDBT) String() string {
	return w.DBT.String()
}

func (w wrappedDBT) Prefix() string {
	return w.DBT.Prefix()
}

func (w wrappedDBT) DataType() string {
	return w.DBT.DataType()
}

func (w wrappedDBT) TalkerID() string {
	return w.DBT.TalkerID()
}

// implement SignalK functions
func (w wrappedDBT) GetDepthBelowTransducer() (float64, error) {
	if v := w.DepthMeters; v > 0 {
		return v, nil
	}
	if v := w.DepthFeet; v > 0 {
		return (unit.Length(v) * unit.Foot).Meters(), nil
	}
	if v := w.DepthFathoms; v > 0 {
		return (unit.Length(v) * unit.Fathom).Meters(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
