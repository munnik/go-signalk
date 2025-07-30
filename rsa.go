package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedRSA struct {
	nmea.RSA
}

func NewRSA(s nmea.RSA) wrappedRSA {
	result := wrappedRSA{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedRSA) String() string {
	return w.RSA.String()
}

func (w wrappedRSA) Prefix() string {
	return w.RSA.Prefix()
}

func (w wrappedRSA) DataType() string {
	return w.RSA.DataType()
}

func (w wrappedRSA) TalkerID() string {
	return w.RSA.TalkerID()
}

// implement SignalK functions
func (w wrappedRSA) GetRudderAngle() (float64, error) {
	if _, err := w.GetRudderAnglePortside(); err == nil {
		return 0, fmt.Errorf("not a single rudder system, use the specific functions for the startboard and portside rudder")
	}
	return w.GetRudderAngleStarboard()
}

// GetRudderAngleStarboard retrieves the rudder angle of the startboard rudder from the sentence
func (w wrappedRSA) GetRudderAngleStarboard() (float64, error) {
	if w.StarboardRudderAngleStatus == nmea.StatusValid {
		return (unit.Angle(w.StarboardRudderAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

// GetRudderAnglePortside retrieves the rudder angle of the portside rudder from the sentence
func (w wrappedRSA) GetRudderAnglePortside() (float64, error) {
	if w.PortRudderAngleStatus == nmea.StatusValid {
		return (unit.Angle(w.PortRudderAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
