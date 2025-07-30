package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedDBS struct {
	nmea.DBS
}

func NewDBS(s nmea.DBS) wrappedDBS {
	result := wrappedDBS{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedDBS) String() string {
	return w.DBS.String()
}

func (w wrappedDBS) Prefix() string {
	return w.DBS.Prefix()
}

func (w wrappedDBS) DataType() string {
	return w.DBS.DataType()
}

func (w wrappedDBS) TalkerID() string {
	return w.DBS.TalkerID()
}

// implement SignalK functions
func (w wrappedDBS) GetDepthBelowSurface() (float64, error) {
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
