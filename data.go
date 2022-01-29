package maidenhead

const (
	latSouthPole         = 90.0
	lngEastwardGreenwich = 180.0

	lngFieldWidth             = 20.0
	lngSquareWidth            = 2.0
	lngSubSquareWidth         = 1 / 12.0
	lngExtendedSquareWidth    = 1 / 120.0
	lngSubExtendedSquareWidth = 1 / 2880.0

	latFieldWidth             = lngFieldWidth / 2
	latSquareWidth            = lngSquareWidth / 2
	latSubSquareWidth         = lngSubSquareWidth / 2
	latExtendedSquareWidth    = lngExtendedSquareWidth / 2
	latSubExtendedSquareWidth = lngSubExtendedSquareWidth / 2

	maxSteps = 5

	FieldPrecision             = 2
	SquarePrecision            = 4
	SubSquarePrecision         = 6
	ExtendedSquarePrecision    = 8
	SubExtendedSquarePrecision = 10
)

var latDivider = [5]float64{latFieldWidth, latSquareWidth, latSubSquareWidth, latExtendedSquareWidth, latSubExtendedSquareWidth}

var lngDivider = [5]float64{lngFieldWidth, lngSquareWidth, lngSubSquareWidth, lngExtendedSquareWidth, lngSubExtendedSquareWidth}

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X"}

var alphaMap map[string]int

func init() {
	alphaMap = make(map[string]int, len(alphaMap))
	for i, v := range alphabet {
		alphaMap[v] = i
	}
}
