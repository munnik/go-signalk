package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

type WrappedGGA struct {
	s nmea.GGA
}

// implement nmea.Sentence functions
func (w WrappedGGA) String() string {
	return w.s.String()
}

func (w WrappedGGA) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedGGA) DataType() string {
	return w.s.DataType()
}

func (w WrappedGGA) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedGGA) GetNumberOfSatellites() (int64, error) {
	if w.s.FixQuality != nmea.Invalid {
		return w.s.NumSatellites, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedGGA) GetPosition3D() (float64, float64, float64, error) {
	if w.s.FixQuality != nmea.Invalid {
		return w.s.Latitude, w.s.Longitude, w.s.Altitude, nil
	}
	return 0, 0, 0, fmt.Errorf("value is unavailable")
}

func (w WrappedGGA) GetFixQuality() (string, error) {
	return w.s.FixQuality, nil
}
