package signalk

import (
	"github.com/adrianmo/go-nmea"
)

type WrappedGSV struct {
	s nmea.GSV
}

// implement nmea.Sentence functions
func (w WrappedGSV) String() string {
	return w.s.String()
}

func (w WrappedGSV) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedGSV) DataType() string {
	return w.s.DataType()
}

func (w WrappedGSV) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedGSV) GetNumberOfSatellites() (int64, error) {
	return w.s.NumberSVsInView, nil
}
