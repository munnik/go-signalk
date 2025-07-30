package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

type wrappedGLL struct {
	nmea.GLL
}

func NewGLL(s nmea.GLL) wrappedGLL {
	result := wrappedGLL{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedGLL) String() string {
	return w.GLL.String()
}

func (w wrappedGLL) Prefix() string {
	return w.GLL.Prefix()
}

func (w wrappedGLL) DataType() string {
	return w.GLL.DataType()
}

func (w wrappedGLL) TalkerID() string {
	return w.GLL.TalkerID()
}

// implement SignalK functions
func (w wrappedGLL) GetPosition2D() (float64, float64, error) {
	if w.Validity != nmea.ValidGLL {
		return w.Latitude, w.Longitude, nil
	}
	return 0, 0, fmt.Errorf("value is unavailable")
}
