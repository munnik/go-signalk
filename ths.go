package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedTHS struct {
	s nmea.THS
}

// implement nmea.Sentence functions
func (w WrappedTHS) String() string {
	return w.s.String()
}

func (w WrappedTHS) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedTHS) DataType() string {
	return w.s.DataType()
}

func (w WrappedTHS) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedTHS) GetTrueHeading() (float64, error) {
	if w.s.Status != nmea.InvalidTHS {
		return (unit.Angle(w.s.Heading) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
