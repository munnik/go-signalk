package signalk

import (
	"fmt"
	"math"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedVBW struct {
	nmea.VBW
}

func NewVBW(s nmea.VBW) wrappedVBW {
	result := wrappedVBW{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedVBW) String() string {
	return w.VBW.String()
}

func (w wrappedVBW) Prefix() string {
	return w.VBW.Prefix()
}

func (w wrappedVBW) DataType() string {
	return w.VBW.DataType()
}

func (w wrappedVBW) TalkerID() string {
	return w.VBW.TalkerID()
}

// implement SignalK functions
func (w wrappedVBW) GetSpeedThroughWater() (float64, error) {
	if w.WaterSpeedStatusValid {
		speed := w.LongitudinalWaterSpeedKnots / math.Abs(w.LongitudinalWaterSpeedKnots) * math.Sqrt(math.Pow(w.LongitudinalWaterSpeedKnots, 2)+math.Pow(w.TransverseWaterSpeedKnots, 2))
		return (unit.Speed(speed) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedVBW) GetSpeedThroughWaterTransverse() (float64, error) {
	if w.WaterSpeedStatusValid {
		return (unit.Speed(w.TransverseWaterSpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedVBW) GetSpeedThroughWaterLongitudinal() (float64, error) {
	if w.WaterSpeedStatusValid {
		return (unit.Speed(w.LongitudinalWaterSpeedKnots) * unit.Knot).MetersPerSecond(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
