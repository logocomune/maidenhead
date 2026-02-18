package maidenhead

import (
	"errors"
)

// Coordinate represents a geographic coordinate.
type Coordinate struct {
	Lat float64
	Lng float64
}

// SquareCoordinate represents the center and corner coordinates of a grid square.
type SquareCoordinate struct {
	Center      Coordinate
	TopLeft     Coordinate
	TopRight    Coordinate
	BottomLeft  Coordinate
	BottomRight Coordinate
}

// GridCenter returns the center coordinates of the given Maidenhead locator.
func GridCenter(locator string) (float64, float64, error) {
	locatorLength := len(locator)

	if locatorLength == 0 {
		return 0, 0, errors.New("locator must not be empty")
	}

	if locatorLength%2 != 0 {
		return 0, 0, errors.New("locator length must be even")
	}

	if locatorLength > maxSteps*2 {
		return 0, 0, errors.New("locator must be at most 10 characters")
	}

	lng := -lngEastwardGreenwich
	lat := -latSouthPole
	lastStep := 0

	for step := 0; step < maxSteps; step++ {
		start := step * 2
		if start+2 > locatorLength {
			break
		}

		dLat, dLng, err := subLocatorDecode(locator[start], locator[start+1], step)
		if err != nil {
			return 0, 0, err
		}

		lat += dLat
		lng += dLng
		lastStep = step
	}

	lat += latDivider[lastStep] / 2.0
	lng += lngDivider[lastStep] / 2.0

	return lat, lng, nil
}

func subLocatorDecode(lngChar, latChar byte, step int) (float64, float64, error) {
	var lat, lng float64
	if step%2 == 0 {
		// Normalize to uppercase
		lngChar &= 0xDF
		latChar &= 0xDF

		lngIdx := int(lngChar - 'A')
		latIdx := int(latChar - 'A')
		if lngIdx < 0 || lngIdx >= alphabetSize || latIdx < 0 || latIdx >= alphabetSize {
			return 0, 0, errors.New("invalid locator format")
		}
		lng = float64(lngIdx)
		lat = float64(latIdx)
	} else {
		lngIdx := int(lngChar - '0')
		latIdx := int(latChar - '0')
		if lngIdx < 0 || lngIdx > 9 || latIdx < 0 || latIdx > 9 {
			return 0, 0, errors.New("invalid locator format")
		}
		lng = float64(lngIdx)
		lat = float64(latIdx)
	}

	lng *= lngDivider[step]
	lat *= latDivider[step]

	return lat, lng, nil
}

// Square returns the center and vertex coordinates of the given Maidenhead locator grid square.
func Square(locator string) (SquareCoordinate, error) {
	centerLat, centerLng, err := GridCenter(locator)
	if err != nil {
		return SquareCoordinate{}, err
	}

	pos := len(locator)/2 - 1

	return SquareCoordinate{
		Center: Coordinate{
			Lat: centerLat,
			Lng: centerLng,
		},
		TopLeft: Coordinate{
			Lat: centerLat + latDivider[pos]/2,
			Lng: centerLng - lngDivider[pos]/2,
		},
		TopRight: Coordinate{
			Lat: centerLat + latDivider[pos]/2,
			Lng: centerLng + lngDivider[pos]/2,
		},
		BottomLeft: Coordinate{
			Lat: centerLat - latDivider[pos]/2,
			Lng: centerLng - lngDivider[pos]/2,
		},
		BottomRight: Coordinate{
			Lat: centerLat - latDivider[pos]/2,
			Lng: centerLng + lngDivider[pos]/2,
		},
	}, nil
}
