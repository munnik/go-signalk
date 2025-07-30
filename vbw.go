package signalk

import (
	"fmt"
	"math"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedVBW struct {
	s nmea.VBW
}

// implement nmea.Sentence functions
func (w WrappedVBW) String() string {
	return w.s.String()
}

func (w WrappedVBW) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedVBW) DataType() string {
	return w.s.DataType()
}

func (w WrappedVBW) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedVBW) GetSpeedThroughWater() (float64, error) {
	if w.s.WaterSpeedStatusValid {
		speed := w.s.LongitudinalWaterSpeedKnots / math.Abs(w.s.LongitudinalWaterSpeedKnots) * math.Sqrt(math.Pow(w.s.LongitudinalWaterSpeedKnots, 2)+math.Pow(w.s.TransverseWaterSpeedKnots, 2))
		return (unit.Speed(speed) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedVBW) GetSpeedThroughWaterTransverse() (float64, error) {
	if w.s.WaterSpeedStatusValid {
		return (unit.Speed(w.s.TransverseWaterSpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedVBW) GetSpeedThroughWaterLongitudinal() (float64, error) {
	if w.s.WaterSpeedStatusValid {
		return (unit.Speed(w.s.LongitudinalWaterSpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
