package signalk

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type wrappedMDA struct {
	nmea.MDA
}

func NewMDA(s nmea.MDA) wrappedMDA {
	result := wrappedMDA{s}
	return result
}

// implement nmea.Sentence functions
func (w wrappedMDA) String() string {
	return w.MDA.String()
}

func (w wrappedMDA) Prefix() string {
	return w.MDA.Prefix()
}

func (w wrappedMDA) DataType() string {
	return w.MDA.DataType()
}

func (w wrappedMDA) TalkerID() string {
	return w.MDA.TalkerID()
}

// implement SignalK functions
func (w wrappedMDA) GetTrueWindDirection() (float64, error) {
	return (unit.Angle(w.WindDirectionTrue) * unit.Degree).Radians(), nil
}

func (w wrappedMDA) GetMagneticWindDirection() (float64, error) {
	return (unit.Angle(w.WindDirectionMagnetic) * unit.Degree).Radians(), nil
}

func (w wrappedMDA) GetWindSpeed() (float64, error) {
	if w.WindSpeedMeters > 0 {
		return w.WindSpeedMeters, nil
	}
	return (unit.Speed(w.WindSpeedKnots) * unit.Knot).MetersPerSecond(), nil
}

func (w wrappedMDA) GetOutsideTemperature() (float64, error) {
	if w.AirTempValid {
		return unit.FromCelsius(w.AirTemp).Kelvin(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedMDA) GetWaterTemperature() (float64, error) {
	if w.WaterTempValid {
		return unit.FromCelsius(w.WaterTemp).Kelvin(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedMDA) GetDewPointTemperature() (float64, error) {
	if w.DewPointValid {
		return unit.FromCelsius(w.DewPoint).Kelvin(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedMDA) GetOutsidePressure() (float64, error) {
	if w.PressureBar > 0 {
		return (unit.Pressure(w.PressureBar) * unit.Bar).Pascals(), nil
	}
	if w.PressureInch > 0 {
		return (unit.Pressure(w.PressureInch) * unit.InchOfMercury).Pascals(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

func (w wrappedMDA) GetHumidity() (float64, error) {
	return w.RelativeHum / 100.0, nil
}
