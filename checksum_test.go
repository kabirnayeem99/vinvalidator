package main

import "testing"

func TestComputeVINCheckDigit_KnownExamples(t *testing.T) {
	tests := []struct {
		Name  string
		Value VIN
		Want  byte
	}{
		{"Honda Accord (check digit 3)", VIN("1HGCM82633A004352"), '3'},
		{"Example with X check digit", VIN("1M8GDM9AXKP042788"), 'X'},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			if len(tt.Value) != 17 {
				t.Fatalf("test VIN must be 17 chars, got %d: %q", len(tt.Value), tt.Value)
			}

			got := ComputeVINCheckDigit(tt.Value)
			if got != tt.Want {
				t.Fatalf("ComputeVINCheckDigit(%q) = %q; want %q", tt.Value, got, tt.Want)
			}
		})
	}
}

func TestComputeVINCheckDigit_MatchesNinthCharacter(t *testing.T) {
	tests := []struct {
		Name  string
		Value VIN
	}{
		{"Honda Accord", VIN("1HGCM82633A004352")},
		{"Example X", VIN("1M8GDM9AXKP042788")},
		{"Saturn (check digit 3)", VIN("5GZCZ43D13S812715")},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			want := tt.Value[8]
			got := ComputeVINCheckDigit(tt.Value)

			if got != want {
				t.Fatalf("ComputeVINCheckDigit(%q) = %q; want %q", tt.Value, got, want)
			}
		})
	}
}
