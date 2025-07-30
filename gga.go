package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

type wrappedGGA struct {
	nmea.GGA
}

func NewGGA(s nmea.GGA) wrappedGGA {
	result := wrappedGGA{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedGGA) String() string {
	return w.GGA.String()
}

func (w wrappedGGA) Prefix() string {
	return w.GGA.Prefix()
}

func (w wrappedGGA) DataType() string {
	return w.GGA.DataType()
}

func (w wrappedGGA) TalkerID() string {
	return w.GGA.TalkerID()
}

// implement SignalK functions
func (w wrappedGGA) GetNumberOfSatellites() (int64, error) {
	if w.FixQuality != nmea.Invalid {
		return w.NumSatellites, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedGGA) GetPosition3D() (float64, float64, float64, error) {
	if w.FixQuality != nmea.Invalid {
		return w.Latitude, w.Longitude, w.Altitude, nil
	}
	return 0, 0, 0, fmt.Errorf("value is unavailable")
}

func (w wrappedGGA) GetFixQuality() (string, error) {
	return w.FixQuality, nil
}
