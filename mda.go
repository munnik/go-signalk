package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type WrappedMDA struct {
	s nmea.MDA
}

// implement nmea.Sentence functions
func (w WrappedMDA) String() string {
	return w.s.String()
}

func (w WrappedMDA) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedMDA) DataType() string {
	return w.s.DataType()
}

func (w WrappedMDA) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedMDA) GetTrueWindDirection() (float64, error) {
	return (unit.Angle(w.s.WindDirectionTrue) * unit.Degree).Radians(), nil
}

func (w WrappedMDA) GetMagneticWindDirection() (float64, error) {
	return (unit.Angle(w.s.WindDirectionMagnetic) * unit.Degree).Radians(), nil
}

func (w WrappedMDA) GetWindSpeed() (float64, error) {
	if w.s.WindSpeedMeters > 0 {
		return w.s.WindSpeedMeters, nil
	}
	return (unit.Speed(w.s.WindSpeedKnots) * unit.Knot).MetersPerSecond(), nil
}

func (w WrappedMDA) GetOutsideTemperature() (float64, error) {
	if w.s.AirTempValid {
		return unit.FromCelsius(w.s.AirTemp).Kelvin(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedMDA) GetWaterTemperature() (float64, error) {
	if w.s.WaterTempValid {
		return unit.FromCelsius(w.s.WaterTemp).Kelvin(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedMDA) GetDewPointTemperature() (float64, error) {
	if w.s.DewPointValid {
		return unit.FromCelsius(w.s.DewPoint).Kelvin(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedMDA) GetOutsidePressure() (float64, error) {
	if w.s.PressureBar > 0 {
		return (unit.Pressure(w.s.PressureBar) * unit.Bar).Pascals(), nil
	}
	if w.s.PressureInch > 0 {
		return (unit.Pressure(w.s.PressureInch) * unit.InchOfMercury).Pascals(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w WrappedMDA) GetHumidity() (float64, error) {
	return w.s.RelativeHum / 100.0, nil
}
