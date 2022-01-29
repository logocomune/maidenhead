package maidenhead

import (
	"errors"
	"strconv"
	"strings"
)

type Coordinate struct {
	Lat float64
	Lng float64
}
type SquareCoordinate struct {
	Center      Coordinate
	TopLeft     Coordinate
	TopRight    Coordinate
	BottomLeft  Coordinate
	BottomRight Coordinate
}

//GridCenter returns the center of the square grid
func GridCenter(locator string) (float64, float64, error) {
	locatorLength := len(locator)
	if locatorLength%2 != 0 {
		return 0, 0, errors.New("locator must be even")
	}

	if locatorLength == 0 {
		return 0, 0, errors.New("locator must not be empty")
	}

	if locatorLength > maxSteps*2 {
		return 0, 0, errors.New("locator must be less than 10 characters")
	}

	locator = strings.ToUpper(locator)

	lng := -lngEastwardGreenwich
	lat := -latSouthPole
	step := 0

	for step < maxSteps {
		start := step * 2
		end := start + 1

		eS := end + 1
		if eS > len(locator) {
			break
		}

		tmpLat, tmpLng, err := subLocatorDecode(locator[start:end], locator[end:eS], step)

		if err != nil {
			return 0, 0, err
		}

		lat += tmpLat
		lng += tmpLng
		step++
	}
	//decrement step
	step--

	lat += latDivider[step] / 2.0
	lng += lngDivider[step] / 2.0

	return lat, lng, nil
}

func subLocatorDecode(subLocatorLng, subLocatorLat string, step int) (float64, float64, error) {
	var lat, lng float64
	if (step % 2) == 0 {
		lng = float64(alphaMap[subLocatorLng])
		lat = float64(alphaMap[subLocatorLat])
	} else {
		iValue, err := strconv.Atoi(subLocatorLng)
		if err != nil {
			return 0, 0, errors.New("invalid locator format")
		}
		lng = float64(iValue)
		iValue, err = strconv.Atoi(subLocatorLat)
		if err != nil {
			return 0, 0, errors.New("invalid locator format")
		}
		lat = float64(iValue)
	}

	lng *= lngDivider[step]
	lat *= latDivider[step]

	return lat, lng, nil
}

//Square returns the coordinates of the vertices of the square grid
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
