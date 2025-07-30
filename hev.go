package signalk

import (
	"github.com/adrianmo/go-nmea"
)

type HEV struct {
	nmea.BaseSentence
	Heave float64 // Heave in meters
}

type wrappedHEV struct {
	HEV
}

func NewHEV(s HEV) wrappedHEV {
	result := wrappedHEV{s}
	return result
}

func init() {
	nmea.RegisterParser("HEV", func(s nmea.BaseSentence) (nmea.Sentence, error) {
		p := nmea.NewParser(s)
		return HEV{
			BaseSentence: s,
			Heave:        p.Float64(0, "heave"),
		}, p.Err()
	})
}

// implement nmea.Sentence functions
func (w wrappedHEV) String() string {
	return w.HEV.String()
}

func (w wrappedHEV) Prefix() string {
	return w.HEV.Prefix()
}

func (w wrappedHEV) DataType() string {
	return w.HEV.DataType()
}

func (w wrappedHEV) TalkerID() string {
	return w.HEV.TalkerID()
}

// implement SignalK functions
func (w wrappedHEV) GetHeave() (float64, error) {
	return w.Heave, nil
}
