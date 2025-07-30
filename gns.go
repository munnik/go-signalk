package signalk

import (
	"fmt"
	"slices"

	"github.com/adrianmo/go-nmea"
)

type wrappedGNS struct {
	nmea.GNS
}

func NewGNS(s nmea.GNS) wrappedGNS {
	result := wrappedGNS{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedGNS) String() string {
	return w.GNS.String()
}

func (w wrappedGNS) Prefix() string {
	return w.GNS.Prefix()
}

func (w wrappedGNS) DataType() string {
	return w.GNS.DataType()
}

func (w wrappedGNS) TalkerID() string {
	return w.GNS.TalkerID()
}

// implement SignalK functions
func (w wrappedGNS) GetPosition3D() (float64, float64, float64, error) {
	if !slices.Contains(w.Mode, nmea.NoFixGNS) {
		return w.Latitude, w.Longitude, w.Altitude, nil
	}
	return 0, 0, 0, fmt.Errorf("value is unavailable")
}
