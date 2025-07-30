package signalk

import (
	"github.com/adrianmo/go-nmea"
)

type HEV struct {
	nmea.BaseSentence
	Heave float64 // Heave in meters
}

type WrappedHEV struct {
	s HEV
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
func (w WrappedHEV) String() string {
	return w.s.String()
}

func (w WrappedHEV) Prefix() string {
	return w.s.Prefix()
}

func (w WrappedHEV) DataType() string {
	return w.s.DataType()
}

func (w WrappedHEV) TalkerID() string {
	return w.s.TalkerID()
}

// implement SignalK functions
func (w WrappedHEV) GetHeave() (float64, error) {
	return w.s.Heave, nil
}
