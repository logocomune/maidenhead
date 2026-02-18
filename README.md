# Maidenhead Locator for golang


![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/logocomune/maidenhead)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/logocomune/maidenhead/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/logocomune/maidenhead)](https://pkg.go.dev/github.com/logocomune/maidenhead)
[![codecov](https://codecov.io/gh/logocomune/maidenhead/graph/badge.svg?token=GGN3PHjyZV)](https://codecov.io/gh/logocomune/maidenhead)
[![Go Report Card](https://goreportcard.com/badge/github.com/logocomune/maidenhead)](https://goreportcard.com/report/github.com/logocomune/maidenhead)
[![CodeFactor](https://www.codefactor.io/repository/github/logocomune/maidenhead/badge)](https://www.codefactor.io/repository/github/logocomune/maidenhead)

The [Maidenhead Locator System](https://en.wikipedia.org/wiki/Maidenhead_Locator_System)
(a.k.a. QTH Locator and IARU Locator) is a geocode system used by amateur radio operators
to succinctly describe their geographic coordinates.

This Golang library converts latitude and longitude coordinates to and from Maidenhead locator strings.

## Installation

```console
go get -u github.com/logocomune/maidenhead
```

## API

### Locator

`Locator(lat, lng float64, precision int) (string, error)` converts latitude and longitude into a Maidenhead locator string with the given precision.

### GridCenter

`GridCenter(locator string) (lat, lng float64, err error)` returns the center coordinates of the given Maidenhead locator.

### Square

`Square(locator string) (SquareCoordinate, error)` returns the center and the four corner coordinates of the grid square.

### Types

```go
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
```

### Precision constants

| Constant | Value | Example |
|---|---|---|
| `FieldPrecision` | 2 | `JN` |
| `SquarePrecision` | 4 | `JN53` |
| `SubSquarePrecision` | 6 | `JN53dk` |
| `ExtendedSquarePrecision` | 8 | `JN53dk06` |
| `SubExtendedSquarePrecision` | 10 | `JN53dk06MK` |

## Usage

```go
package main

import (
	"fmt"
	"github.com/logocomune/maidenhead"
)

func main() {
	latitude := 43.723073
	longitude := 10.396637

	locator, _ := maidenhead.Locator(latitude, longitude, maidenhead.FieldPrecision)
	fmt.Println("Locator with field precision:", locator)
	locator, _ = maidenhead.Locator(latitude, longitude, maidenhead.SquarePrecision)
	fmt.Println("Locator with square precision:", locator)
	locator, err := maidenhead.Locator(latitude, longitude, maidenhead.SubSquarePrecision)
	fmt.Println("Locator with sub square precision:", locator, err)
	locator, _ = maidenhead.Locator(latitude, longitude, maidenhead.ExtendedSquarePrecision)
	fmt.Println("Locator with extended square precision:", locator)
	locator, _ = maidenhead.Locator(latitude, longitude, maidenhead.SubExtendedSquarePrecision)
	fmt.Println("Locator with sub extended square precision:", locator)

	lat, lng, _ := maidenhead.GridCenter("JN53er73OM")
	fmt.Printf("Grid center of %s is lat: %f and lng: %f\n", "JN53er73OM", lat, lng)

	square, _ := maidenhead.Square("JN53er73OM")
	fmt.Printf("Square coordinates of %s are %+v\n", "JN53er73OM", square)
}
```
