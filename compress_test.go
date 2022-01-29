package maidenhead

import "testing"

func TestLocator(t *testing.T) {

	testLat := 43.443534
	testLng := 10.254315
	testLocator := "JN53dk06MK"

	type args struct {
		lat      float64
		lng      float64
		gridSize int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Bad latitude 01",
			args: args{
				lat:      90,
				lng:      0,
				gridSize: SquarePrecision,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Bad latitude 02",
			args: args{
				lat:      91,
				lng:      0,
				gridSize: SquarePrecision,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Bad latitude 03",
			args: args{
				lat:      -91,
				lng:      0,
				gridSize: SquarePrecision,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Bad longitude 01",
			args: args{
				lat:      0,
				lng:      180.1,
				gridSize: SquarePrecision,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Bad longitude 02",
			args: args{
				lat:      0,
				lng:      -180.1,
				gridSize: SquarePrecision,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Grid size not even",
			args: args{
				lat:      0,
				lng:      0,
				gridSize: 5,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Grid size bigger than 10",
			args: args{
				lat:      0,
				lng:      0,
				gridSize: 12,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Grid size 2",
			args: args{
				lat:      testLat,
				lng:      testLng,
				gridSize: FieldPrecision,
			},
			want:    testLocator[0:FieldPrecision],
			wantErr: false,
		},

		{
			name: "Grid size 4",
			args: args{
				lat:      testLat,
				lng:      testLng,
				gridSize: SquarePrecision,
			},
			want:    testLocator[0:SquarePrecision],
			wantErr: false,
		},
		{
			name: "Grid size 6",
			args: args{
				lat:      testLat,
				lng:      testLng,
				gridSize: SubSquarePrecision,
			},
			want:    testLocator[0:SubSquarePrecision],
			wantErr: false,
		},
		{
			name: "Grid size 8",
			args: args{
				lat:      testLat,
				lng:      testLng,
				gridSize: ExtendedSquarePrecision,
			},
			want:    testLocator[0:ExtendedSquarePrecision],
			wantErr: false,
		},
		{
			name: "Grid size 10",
			args: args{
				lat:      testLat,
				lng:      testLng,
				gridSize: SubExtendedSquarePrecision,
			},
			want:    testLocator[0:SubExtendedSquarePrecision],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Locator(tt.args.lat, tt.args.lng, tt.args.gridSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Locator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Locator() got = %v, want %v", got, tt.want)
			}
		})
	}
}
