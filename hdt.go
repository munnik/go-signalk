package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedHDT struct {
	nmea.HDT
}

func NewHDT(s nmea.HDT) wrappedHDT {
	result := wrappedHDT{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedHDT) String() string {
	return w.HDT.String()
}

func (w wrappedHDT) Prefix() string {
	return w.HDT.Prefix()
}

func (w wrappedHDT) DataType() string {
	return w.HDT.DataType()
}

func (w wrappedHDT) TalkerID() string {
	return w.HDT.TalkerID()
}

// implement SignalK functions
func (w wrappedHDT) GetTrueHeading() (float64, error) {
	if w.True {
		return (unit.Angle(w.Heading) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
