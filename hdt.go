package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedHDT struct {
	s nmea.HDT
}

// implement nmea.Sentence functions
func (w WrappedHDT) String() string {
	return w.s.String()
}

func (w WrappedHDT) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedHDT) DataType() string {
	return w.s.DataType()
}

func (w WrappedHDT) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedHDT) GetTrueHeading() (float64, error) {
	if w.s.True {
		return (unit.Angle(w.s.Heading) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
