# Maidenhead Locator for golang

[![CircleCI](https://circleci.com/gh/logocomune/maidenhead/tree/main.svg?style=svg)](https://circleci.com/gh/logocomune/maidenhead/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/logocomune/maidenhead)](https://goreportcard.com/report/github.com/logocomune/maidenhead)
[![codecov](https://codecov.io/gh/logocomune/maidenhead/branch/master/graph/badge.svg)](https://codecov.io/gh/logocomune/maidenhead)

The [Maidenhead Locator System](https://en.wikipedia.org/wiki/Maidenhead_Locator_System)
(a.k.a. QTH Locator and IARU Locator) is a geocode system used by amateur radio operators
to succinctly describe their geographic coordinates.

This Golang library compresses and decompresses (latitude, longitude) coordinates to and
from a Maidenhead locator.

## Installation

```console
go get -u github.com/logocomune/maidenhead
```

## Usage

```golang
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
