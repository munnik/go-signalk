package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

type WrappedGLL struct {
	s nmea.GLL
}

// implement nmea.Sentence functions
func (w WrappedGLL) String() string {
	return w.s.String()
}

func (w WrappedGLL) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedGLL) DataType() string {
	return w.s.DataType()
}

func (w WrappedGLL) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedGLL) GetPosition2D() (float64, float64, error) {
	if w.s.Validity != nmea.ValidGLL {
		return w.s.Latitude, w.s.Longitude, nil
	}
	return 0, 0, fmt.Errorf("value is unavailable")
}
