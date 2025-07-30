package signalk

import (
	"github.com/adrianmo/go-nmea"
)

type wrappedGSV struct {
	nmea.GSV
}

func NewGSV(s nmea.GSV) wrappedGSV {
	result := wrappedGSV{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedGSV) String() string {
	return w.GSV.String()
}

func (w wrappedGSV) Prefix() string {
	return w.GSV.Prefix()
}

func (w wrappedGSV) DataType() string {
	return w.GSV.DataType()
}

func (w wrappedGSV) TalkerID() string {
	return w.GSV.TalkerID()
}

// implement SignalK functions
func (w wrappedGSV) GetNumberOfSatellites() (int64, error) {
	return w.NumberSVsInView, nil
}
