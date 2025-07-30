package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedTHS struct {
	nmea.THS
}

func NewTHS(s nmea.THS) wrappedTHS {
	result := wrappedTHS{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedTHS) String() string {
	return w.THS.String()
}

func (w wrappedTHS) Prefix() string {
	return w.THS.Prefix()
}

func (w wrappedTHS) DataType() string {
	return w.THS.DataType()
}

func (w wrappedTHS) TalkerID() string {
	return w.THS.TalkerID()
}

// implement SignalK functions
func (w wrappedTHS) GetTrueHeading() (float64, error) {
	if w.Status != nmea.InvalidTHS {
		return (unit.Angle(w.Heading) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
