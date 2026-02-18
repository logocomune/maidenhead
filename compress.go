package maidenhead

// http://www.w8bh.net/grid_squares.pdf
import (
	"errors"
	"math"
	"strings"
)

// Locator returns the Maidenhead locator for the given latitude and longitude and requested precision.
func Locator(lat, lng float64, precision int) (string, error) {
	if math.Abs(lat) >= 90 {
		return "", errors.New("invalid latitude: allowed values are between -90 and 90")
	}

	if math.Abs(lng) > 180 {
		return "", errors.New("invalid longitude: allowed values are between -180 and 180")
	}

	if precision <= 0 {
		return "", errors.New("precision must be greater than zero")
	}

	if precision%2 != 0 {
		return "", errors.New("precision must be even")
	}

	if precision/2 > maxSteps {
		return "", errors.New("precision must be less or equal to 10")
	}

	lat += latSouthPole
	lng += lngEastwardGreenwich

	var buf strings.Builder
	buf.Grow(precision)

	steps := precision / 2
	for step := 0; step < steps; step++ {
		latIdx := int(math.Trunc(lat / latDivider[step]))
		lngIdx := int(math.Trunc(lng / lngDivider[step]))

		if step%2 == 0 {
			lngChar := alphabet[lngIdx]
			latChar := alphabet[latIdx]
			if step == 2 {
				lngChar |= 0x20
				latChar |= 0x20
			}
			buf.WriteByte(lngChar)
			buf.WriteByte(latChar)
		} else {
			buf.WriteByte(digits[lngIdx])
			buf.WriteByte(digits[latIdx])
		}

		lat -= float64(latIdx) * latDivider[step]
		lng -= float64(lngIdx) * lngDivider[step]
	}

	return buf.String(), nil
}
