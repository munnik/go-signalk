package signalk

import (
	"fmt"
	"slices"

	"github.com/adrianmo/go-nmea"
)

type WrappedGNS struct {
	s nmea.GNS
}

// implement nmea.Sentence functions
func (w WrappedGNS) String() string {
	return w.s.String()
}

func (w WrappedGNS) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedGNS) DataType() string {
	return w.s.DataType()
}

func (w WrappedGNS) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedGNS) GetPosition3D() (float64, float64, float64, error) {
	if !slices.Contains(w.s.Mode, nmea.NoFixGNS) {
		return w.s.Latitude, w.s.Longitude, w.s.Altitude, nil
	}
	return 0, 0, 0, fmt.Errorf("value is unavailable")
}
