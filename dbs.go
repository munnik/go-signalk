package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedDBS struct {
	s nmea.DBS
}

// implement nmea.Sentence functions
func (w WrappedDBS) String() string {
	return w.s.String()
}

func (w WrappedDBS) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedDBS) DataType() string {
	return w.s.DataType()
}

func (w WrappedDBS) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedDBS) GetDepthBelowSurface() (float64, error) {
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
