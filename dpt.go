package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

type wrappedDPT struct {
	nmea.DPT
}

func NewDPT(s nmea.DPT) wrappedDPT {
	result := wrappedDPT{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedDPT) String() string {
	return w.DPT.String()
}

func (w wrappedDPT) Prefix() string {
	return w.DPT.Prefix()
}

func (w wrappedDPT) DataType() string {
	return w.DPT.DataType()
}

func (w wrappedDPT) TalkerID() string {
	return w.DPT.TalkerID()
}

// implement SignalK functions
func (w wrappedDPT) GetDepthBelowTransducer() (float64, error) {
	if v := w.Depth; v > 0 {
		return v, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedDPT) GetDepthBelowKeel() (float64, error) {
	if v := w.Depth; v > 0 {
		return v, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedDPT) GetDepthBelowSurface() (float64, error) {
	if v := w.Depth; v > 0 {
		return v + w.Offset, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
