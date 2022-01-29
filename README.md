# Maidenhead Locator for golang
[![Build Status](https://app.travis-ci.com/logocomune/maidenhead.svg?branch=master)](https://app.travis-ci.com/logocomune/maidenhead)
[![Go Report Card](https://goreportcard.com/badge/github.com/logocomune/maidenhead)](https://goreportcard.com/report/github.com/logocomune/maidenhead)
[![codecov](https://codecov.io/gh/logocomune/maidenhead/branch/master/graph/badge.svg)](https://codecov.io/gh/logocomune/maidenhead)

The Maidenhead Locator System (a.k.a. QTH Locator and IARU Locator) is a geocode system used by amateur radio operators to succinctly describe their geographic coordinates.

This golang library compress and decompress latitude and longitude coordinates into Maidenhead locator 

## Installation

`go get -u github.com/logocomune/maidenhead`

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

	locator, _ := maidenhead.Locator(latitude, longitude, maidenhead.FIELD_PRECISION)
	fmt.Println("Locator with field precision:", locator)
	locator, _ = maidenhead.Locator(latitude, longitude, maidenhead.SQUARE_PRECSION)
	fmt.Println("Locator with square precision:", locator)
	locator, err := maidenhead.Locator(latitude, longitude, maidenhead.SUB_SQUARE_PRECISION)
	fmt.Println("Locator with sub square precision:", locator, err)
	locator, _ = maidenhead.Locator(latitude, longitude, maidenhead.EXTENDED_SQUARE_PRECSION)
	fmt.Println("Locator with extended square precision:", locator)
	locator, _ = maidenhead.Locator(latitude, longitude, maidenhead.SUB_EXTENDED_SQUARE_PRECISION)
	fmt.Println("Locator with sub extended square precision:", locator)

	lat, lng, _ := maidenhead.GridCenter("JN53er73OM")
	fmt.Printf("Grid center of %s is lat: %f and lng: %f\n", "JN53er73OM", lat, lng)

	square, _ := maidenhead.Square("JN53er73OM")
	fmt.Printf("Square coordinates of %s are %+v\n", "JN53er73OM", square)

}
```
