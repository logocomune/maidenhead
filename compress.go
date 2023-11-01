package maidenhead

//http://www.w8bh.net/grid_squares.pdf
import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Locator returns the Maidenhead locator for the given latitude and longitude and requested precision.
func Locator(lat, lng float64, precision int) (string, error) {
	if math.Abs(lat) >= 90 {
		return "", errors.New("invalid latitude allowed values are between -90 and 90")
	}

	if math.Abs(lng) > 180 {
		return "", errors.New("invalid longitude allowed values are between -180 and 180")
	}

	if precision%2 != 0 {
		return "", errors.New("grid size must be even ")
	}

	if precision/2 > maxSteps {
		return "", errors.New("grid size must be less than 5")
	}

	lat += latSouthPole
	lng += lngEastwardGreenwich

	step := 0

	var locator, tmpLoc string
	for step < maxSteps {
		tmpLoc, lat, lng = subLocator(lat, lng, step)
		if step == 2 {
			tmpLoc = strings.ToLower(tmpLoc)
		}

		locator += tmpLoc

		if step == (precision/2 - 1) {
			break
		}
		step++
	}

	return locator, nil
}

func subLocator(lat, lng float64, step int) (string, float64, float64) {
	latTmp := math.Trunc(lat / latDivider[step])
	lngTmp := math.Trunc(lng / lngDivider[step])

	locator := strconv.Itoa(int(lngTmp)) + strconv.Itoa(int(latTmp))

	if (step % 2) == 0 {
		locator = alphabet[int(lngTmp)] + alphabet[int(latTmp)]
	}

	lat -= latTmp * latDivider[step]
	lng -= lngTmp * lngDivider[step]

	return locator, lat, lng
}
