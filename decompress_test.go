package maidenhead

import (
	"reflect"
	"testing"
)

func TestGridCenter(t *testing.T) {

	type args struct {
		locator string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		want1   float64
		wantErr bool
	}{
		{
			name: "Locator not even",
			args: args{
				locator: "JN5",
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Locator not empty",
			args: args{
				locator: "",
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Locator too big",
			args: args{
				locator: "JN53dk06MK00",
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Locator bad format #1",
			args: args{
				locator: "JNA0dk06MK",
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Locator bad format #2",
			args: args{
				locator: "JN0Adk06MK",
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Decode field locator",
			args: args{
				locator: "AA",
			},
			want:    0 + latFieldWidth/2 - latSouthPole,
			want1:   0 + lngFieldWidth/2 - lngEastwardGreenwich,
			wantErr: false,
		},
		{
			name: "Decode field and square locator",
			args: args{
				locator: "AA00",
			},
			want:    0 + latSquareWidth/2 - latSouthPole,
			want1:   0 + lngSquareWidth/2 - lngEastwardGreenwich,
			wantErr: false,
		},
		{
			name: "Decode field, square and subsquare locator",
			args: args{
				locator: "AA00aa",
			},
			want:    0 + latSubSquareWidth/2 - latSouthPole,
			want1:   0 + lngSubSquareWidth/2 - lngEastwardGreenwich,
			wantErr: false,
		},
		{
			name: "Decode field,square, subsquare and extended subsquare locator",
			args: args{
				locator: "AA00aa00",
			},
			want:    0 + latExtendedSquareWidth/2 - latSouthPole,
			want1:   0 + lngExtendedSquareWidth/2 - lngEastwardGreenwich,
			wantErr: false,
		},
		{
			name: "Decode field,square, subsquare, extended subsquare and extended sub locator locator",
			args: args{
				locator: "AA00aa00aa",
			},
			want:    0 + latSubExtendedSquareWidth/2 - latSouthPole,
			want1:   0 + lngSubExtendedSquareWidth/2 - lngEastwardGreenwich,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GridCenter(tt.args.locator)
			if (err != nil) != tt.wantErr {
				t.Errorf("GridCenter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GridCenter() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GridCenter() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	type args struct {
		locator string
	}
	var tests = []struct {
		name    string
		args    args
		want    SquareCoordinate
		wantErr bool
	}{
		{
			name: "Locator not even",
			args: args{
				locator: "JN5",
			},
			want: SquareCoordinate{},

			wantErr: true,
		},
		{
			name: "Locator not empty",
			args: args{
				locator: "",
			},
			want: SquareCoordinate{},

			wantErr: true,
		},
		{
			name: "Locator too big",
			args: args{
				locator: "JN53dk06MK00",
			},
			want: SquareCoordinate{},

			wantErr: true,
		},
		{
			name: "Decode field locator",
			args: args{
				locator: "JN",
			},
			want: SquareCoordinate{
				Center: Coordinate{
					Lat: 45.0,
					Lng: 10.0,
				},
				TopLeft: Coordinate{
					Lat: 50.0,
					Lng: 0.0,
				},
				TopRight: Coordinate{
					Lat: 50.0,
					Lng: 20.0,
				},
				BottomLeft: Coordinate{
					Lat: 40.0,
					Lng: 0.0,
				},
				BottomRight: Coordinate{
					Lat: 40.0,
					Lng: 20.0,
				},
			},

			wantErr: false,
		},
		{
			name: "Decode field and square locator",
			args: args{
				locator: "JN53",
			},
			want: SquareCoordinate{
				Center: Coordinate{
					Lat: 43.5,
					Lng: 11.0,
				},
				TopLeft: Coordinate{
					Lat: 44.0,
					Lng: 10.0,
				},
				TopRight: Coordinate{
					Lat: 44.0,
					Lng: 12.0,
				},
				BottomLeft: Coordinate{
					Lat: 43.0,
					Lng: 10.0,
				},
				BottomRight: Coordinate{
					Lat: 43.0,
					Lng: 12.0,
				},
			},

			wantErr: false,
		},
		{
			name: "Decode field, square and subsquare locator",
			args: args{
				locator: "JN53ce",
			},
			want: SquareCoordinate{
				Center: Coordinate{
					Lat: 43.1875,
					Lng: 10.208333333333332,
				},
				TopLeft: Coordinate{
					Lat: 43.208333333333336,
					Lng: 10.166666666666666,
				},
				TopRight: Coordinate{
					Lat: 43.208333333333336,
					Lng: 10.249999999999998,
				},
				BottomLeft: Coordinate{
					Lat: 43.166666666666664,
					Lng: 10.166666666666666,
				},
				BottomRight: Coordinate{
					Lat: 43.166666666666664,
					Lng: 10.249999999999998,
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Square(tt.args.locator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Square() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Square() got = %+v, want %v", got, tt.want)
			}
		})
	}
}
