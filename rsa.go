package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedRSA struct {
	s nmea.RSA
}

// implement nmea.Sentence functions
func (w WrappedRSA) String() string {
	return w.s.String()
}

func (w WrappedRSA) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedRSA) DataType() string {
	return w.s.DataType()
}

func (w WrappedRSA) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedRSA) GetRudderAngle() (float64, error) {
	if _, err := w.GetRudderAnglePortside(); err == nil {
		return 0, fmt.Errorf("not a single rudder system, use the specific functions for the startboard and portside rudder")
	}
	return w.GetRudderAngleStarboard()
}

// GetRudderAngleStarboard retrieves the rudder angle of the startboard rudder from the sentence
func (w WrappedRSA) GetRudderAngleStarboard() (float64, error) {
	if w.s.StarboardRudderAngleStatus == nmea.StatusValid {
		return (unit.Angle(w.s.StarboardRudderAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

// GetRudderAnglePortside retrieves the rudder angle of the portside rudder from the sentence
func (w WrappedRSA) GetRudderAnglePortside() (float64, error) {
	if w.s.PortRudderAngleStatus == nmea.StatusValid {
		return (unit.Angle(w.s.PortRudderAngle) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
