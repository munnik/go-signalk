package signalk

import (
	"fmt"
	"time"

	"github.com/adrianmo/go-nmea"
)

// MagneticCourseOverGround retrieves the magnetic course over ground from the sentence
type MagneticCourseOverGround interface {
	GetmagneticCourseOverGround() (float64, error)
}

// MagneticHeading retrieves the magnetic heading from the sentence
type MagneticHeading interface {
	GetMagneticHeading() (float64, error)
}

// MagneticVariation retrieves the magnetic variation from the sentence
type MagneticVariation interface {
	GetMagneticVariation() (float64, error)
}

// RateOfTurn retrieves the rate of turn from the sentence
type RateOfTurn interface {
	GetRateOfTurn() (float64, error)
}

// TrueCourseOverGround retrieves the true course over ground from the sentence
type TrueCourseOverGround interface {
	GetTrueCourseOverGround() (float64, error)
}

// TrueHeading retrieves the true heading from the sentence
type TrueHeading interface {
	GetTrueHeading() (float64, error)
}

// FixQuality retrieves the fix quality from the sentence
type FixQuality interface {
	GetFixQuality() (string, error)
}

// FixType retrieves the fix type from the sentence
type FixType interface {
	GetFixType() (string, error)
}

// NumberOfSatellites retrieves the number of satellites from the sentence
type NumberOfSatellites interface {
	GetNumberOfSatellites() (int64, error)
}

// Position2D retrieves the 2D position from the sentence
type Position2D interface {
	GetPosition2D() (float64, float64, error)
}

// Position3D retrieves the 3D position from the sentence
type Position3D interface {
	GetPosition3D() (float64, float64, float64, error)
}

// SpeedOverGround retrieves the speed over ground from the sentence
type SpeedOverGround interface {
	GetSpeedOverGround() (float64, error)
}

// SpeedThroughWater retrieves the speed through water from the sentence
type SpeedThroughWater interface {
	GetSpeedThroughWater() (float64, error)
	GetSpeedThroughWaterTransverse() (float64, error)
	GetSpeedThroughWaterLongitudinal() (float64, error)
}

// Log retrieves the distance through water from the sentence
type DistanceThroughWater interface {
	GetLog() (float64, error)
	GetTripLog() (float64, error)
}

// DepthBelowSurface retrieves the depth below surface from the sentence
type DepthBelowSurface interface {
	GetDepthBelowSurface() (float64, error)
}

// DepthBelowKeel retrieves the depth below keel from the sentence
type DepthBelowKeel interface {
	GetDepthBelowKeel() (float64, error)
}

// DepthBelowTransducer retrieves the depth below the transducer from the sentence
type DepthBelowTransducer interface {
	GetDepthBelowTransducer() (float64, error)
}

// WaterTemperature retrieves the water temperature from the sentence
type WaterTemperature interface {
	GetWaterTemperature() (float64, error)
}

// TrueWindDirection retrieves the true wind direction from the sentence
type TrueWindDirection interface {
	GetTrueWindDirection() (float64, error)
}

// MagneticWindDirection retrieves the magnetic wind direction from the sentence
type MagneticWindDirection interface {
	GetMagneticWindDirection() (float64, error)
}

// RelativeWindDirection retrieves the relative wind direction from the sentence
type RelativeWindDirection interface {
	GetRelativeWindDirection() (float64, error)
}

// WindSpeed retrieves the wind speed from the sentence
type WindSpeed interface {
	GetWindSpeed() (float64, error)
}

// OutsideTemperature retrieves the outside air temperature from the sentence
type OutsideTemperature interface {
	GetOutsideTemperature() (float64, error)
}

// DewPointTemperature retrieves the dew point temperature from the sentence
type DewPointTemperature interface {
	GetDewPointTemperature() (float64, error)
}

// Humidity retrieves the relative humidity from the sentence
type Humidity interface {
	GetHumidity() (float64, error)
}

// Heave retrieves the heave from the sentence
type Heave interface {
	GetHeave() (float64, error)
}

// DateTime retrieves the date and time in RFC3339Nano format
type DateTime interface {
	GetDateTime() (string, error)
}

// CallSign retrieves the call sign of the vessel from the sentence
type CallSign interface {
	GetCallSign() (string, error)
}

// ENINumber retrieves the ENI number of the vessel from the sentence
type ENINumber interface {
	// https://en.wikipedia.org/wiki/ENI_number
	GetENINumber() (string, error)
}

// IMONumber retrieves the IMO number of the vessel from the sentence
type IMONumber interface {
	GetIMONumber() (string, error)
}

// MMSI retrieves the MMSI of the vessel from the sentence
type MMSI interface {
	GetMMSI() (string, error)
}

// NavigationStatus retrieves the navigation status from the sentence
type NavigationStatus interface {
	GetNavigationStatus() (string, error)
}

// VesselLength retrieves the length of the vessel from the sentence
type VesselLength interface {
	GetVesselLength() (float64, error)
}

// VesselBeam retrieves the beam of the vessel from the sentence
type VesselBeam interface {
	GetVesselBeam() (float64, error)
}

// VesselName retrieves the name of the vessel from the sentence
type VesselName interface {
	GetVesselName() (string, error)
}

// VesselType retrieves the type of the vessel from the sentence
type VesselType interface {
	GetVesselType() (string, error)
}

// Destination retriFspeedtreves the destination of the vessel from the sentence
type Destination interface {
	GetDestination() (string, error)
}

// ETA retrieves the ETA of the vessel from the sentence
type ETA interface {
	GetETA() (time.Time, error)
}

// RudderAngle retrieves the rudder angle from the sentence
type RudderAngle interface {
	GetRudderAngle() (float64, error)
	GetRudderAngleStarboard() (float64, error)
	GetRudderAnglePortside() (float64, error)
}

// Alarm retrieves alarm information from the sentence
type Alarm interface {
	IsActive() (bool, error)
	IsUnacknowledged() (bool, error)
	GetDescription() (string, error)
}

func Parse(raw string) (nmea.Sentence, error) {
	s, error := nmea.Parse(raw)
	if error != nil {
		return nil, error
	}

	switch s := s.(type) {
	case nmea.DBS:
		return NewDBS(s), nil
	case nmea.DBT:
		return NewDBT(s), nil
	case nmea.DPT:
		return NewDPT(s), nil
	case nmea.GGA:
		return NewGGA(s), nil
	case nmea.GLL:
		return NewGLL(s), nil
	case nmea.GNS:
		return NewGNS(s), nil
	case nmea.GSA:
		return NewGSA(s), nil
	case nmea.GSV:
		return NewGSV(s), nil
	case nmea.HDT:
		return NewHDT(s), nil
	case HEV:
		return NewHEV(s), nil
	case nmea.MDA:
		return NewMDA(s), nil
	case nmea.MWD:
		return NewMWD(s), nil
	case nmea.MWV:
		return NewMWV(s), nil
	case nmea.RMC:
		return NewRMC(s), nil
	case nmea.ROT:
		return NewROT(s), nil
	case nmea.RSA:
		return NewRSA(s), nil
	case nmea.THS:
		return NewTHS(s), nil
	case nmea.VBW:
		return NewVBW(s), nil
	case nmea.VDMVDO:
		return NewVDMVDO(s), nil
	case nmea.VHW:
		return NewVHW(s), nil
	case nmea.VLW:
		return NewVLW(s), nil
	case nmea.VTG:
		return NewVTG(s), nil
	case nmea.VWR:
		return NewVWR(s), nil
	case nmea.ZDA:
		return NewZDA(s), nil
	}
	return s, fmt.Errorf("could not convert to SignalK sentence, %s is not supported", s.DataType())
}
