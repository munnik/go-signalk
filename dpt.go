package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

type WrappedDPT struct {
	s nmea.DPT
}

// implement nmea.Sentence functions
func (w WrappedDPT) String() string {
	return w.s.String()
}

func (w WrappedDPT) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedDPT) DataType() string {
	return w.s.DataType()
}

func (w WrappedDPT) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedDPT) GetDepthBelowTransducer() (float64, error) {
	if v := w.s.Depth; v > 0 {
		return v, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedDPT) GetDepthBelowKeel() (float64, error) {
	if v := w.s.Depth; v > 0 {
		return v, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedDPT) GetDepthBelowSurface() (float64, error) {
	if v := w.s.Depth; v > 0 {
		return v + w.s.Offset, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
