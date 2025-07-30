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
	case nmea.RMC:
		return WrappedRMC{s: s}, nil
		// case nmea.TypeAAM:
		// 	return newAAM(s)
		// case nmea.TypeACK:
		// 	return newACK(s)
		// case nmea.TypeACN:
		// 	return newACN(s)
		// case nmea.TypeALA:
		// 	return newALA(s)
		// case nmea.TypeALC:
		// 	return newALC(s)
		// case nmea.TypeALF:
		// 	return newALF(s)
		// case nmea.TypeALR:
		// 	return newALR(s)
		// case nmea.TypeAPB:
		// 	return newAPB(s)
		// case nmea.TypeARC:
		// 	return newARC(s)
		// case nmea.TypeBEC:
		// 	return newBEC(s)
		// case nmea.TypeBOD:
		// 	return newBOD(s)
		// case nmea.TypeBWC:
		// 	return newBWC(s)
		// case nmea.TypeBWR:
		// 	return newBWR(s)
		// case nmea.TypeBWW:
		// 	return newBWW(s)
		// case nmea.TypeDOR:
		// 	return newDOR(s)
		// case nmea.TypeDSC:
		// 	return newDSC(s)
		// case nmea.TypeDSE:
		// 	return newDSE(s)
		// case nmea.TypeDTM:
		// 	return newDTM(s)
		// case nmea.TypeEVE:
		// 	return newEVE(s)
		// case nmea.TypeFIR:
		// 	return newFIR(s)
		// case nmea.TypeGGA:
		// 	return newGGA(s)
		// case nmea.TypeGSA:
		// 	return newGSA(s)
		// case nmea.TypeGLL:
		// 	return newGLL(s)
		// case nmea.TypeVTG:
		// 	return newVTG(s)
		// case nmea.TypeZDA:
		// 	return newZDA(s)
		// case nmea.TypePGN:
		// 	return newPGN(s)
		// case nmea.TypePCDIN:
		// 	return newPCDIN(s)
		// case nmea.TypePGRME:
		// 	return newPGRME(s)
		// case nmea.TypePGRMT:
		// 	return newPGRMT(s)
		// case nmea.TypePHTRO:
		// 	return newPHTRO(s)
		// case nmea.TypePMTK001:
		// 	return newPMTK001(s)
		// case nmea.TypePRDID:
		// 	return newPRDID(s)
		// case nmea.TypePSKPDPT:
		// 	return newPSKPDPT(s)
		// case nmea.TypePSONCMS:
		// 	return newPSONCMS(s)
		// case nmea.TypeQuery:
		// 	return newQuery(s)
		// case nmea.TypeGSV:
		// 	return newGSV(s)
		// case nmea.TypeHBT:
		// 	return newHBT(s)
		// case nmea.TypeHDG:
		// 	return newHDG(s)
		// case nmea.TypeHDT:
		// 	return newHDT(s)
		// case nmea.TypeHDM:
		// 	return newHDM(s)
		// case nmea.TypeHSC:
		// 	return newHSC(s)
		// case nmea.TypeGNS:
		// 	return newGNS(s)
		// case nmea.TypeTHS:
		// 	return newTHS(s)
		// case nmea.TypeTLB:
		// 	return newTLB(s)
		// case nmea.TypeTLL:
		// 	return newTLL(s)
		// case nmea.TypeTTM:
		// 	return newTTM(s)
		// case nmea.TypeTXT:
		// 	return newTXT(s)
		// case nmea.TypeWPL:
		// 	return newWPL(s)
		// case nmea.TypeRMB:
		// 	return newRMB(s)
		// case nmea.TypeRPM:
		// 	return newRPM(s)
		// case nmea.TypeRSA:
		// 	return newRSA(s)
		// case nmea.TypeRSD:
		// 	return newRSD(s)
		// case nmea.TypeRTE:
		// 	return newRTE(s)
		// case nmea.TypeROT:
		// 	return newROT(s)
		// case nmea.TypeVBW:
		// 	return newVBW(s)
		// case nmea.TypeVDR:
		// 	return newVDR(s)
		// case nmea.TypeVHW:
		// 	return newVHW(s)
		// case nmea.TypeVSD:
		// 	return newVSD(s)
		// case nmea.TypeVPW:
		// 	return newVPW(s)
		// case nmea.TypeVLW:
		// 	return newVLW(s)
		// case nmea.TypeVWR:
		// 	return newVWR(s)
		// case nmea.TypeVWT:
		// 	return newVWT(s)
		// case nmea.TypeDPT:
		// 	return newDPT(s)
		// case nmea.TypeDBT:
		// 	return newDBT(s)
		// case nmea.TypeDBK:
		// 	return newDBK(s)
		// case nmea.TypeDBS:
		// 	return newDBS(s)
		// case nmea.TypeMDA:
		// 	return newMDA(s)
		// case nmea.TypeMTA:
		// 	return newMTA(s)
		// case nmea.TypeMTW:
		// 	return newMTW(s)
		// case nmea.TypeMWD:
		// 	return newMWD(s)
		// case nmea.TypeMWV:
		// 	return newMWV(s)
		// case nmea.TypeOSD:
		// 	return newOSD(s)
		// case nmea.TypeXDR:
		// 	return newXDR(s)
		// case nmea.TypeXTE:
		// 	return newXTE(s)
		// case nmea.TypePKLID:
		// 	return newPKLID(s)
		// case nmea.TypePKNID:
		// 	return newPKNID(s)
		// case nmea.TypePKLSH:
		// 	return newPKLSH(s)
		// case nmea.TypePKNSH:
		// 	return newPKNSH(s)
		// case nmea.TypePKLDS:
		// 	return newPKLDS(s)
		// case nmea.TypePKNDS:
		// 	return newPKNDS(s)
		// case nmea.TypePKWDWPL:
		// 	return newPKWDWPL(s)
		// case nmea.TypeABM:
		// 	return newABM(s)
		// case nmea.TypeBBM:
		// 	return newBBM(s)
		// case nmea.TypeTTD:
		// 	return newTTD(s)
		// case nmea.TypeVDM, TypeVDO:
		// 	return newVDMVDO(s)
	}
	return s, fmt.Errorf("could not convert to SignalK sentence, %s is not supported", s.DataType())
}
