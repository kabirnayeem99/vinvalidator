package main

import "testing"

func TestComputeVINCheckDigit_KnownExamples(t *testing.T) {

	tests := []struct {
		Name  string
		Value VIN
		Want  byte
	}{
		{"Honda Accord (check digit 3)", VIN("1HGCM82633A004352"), '3'},
		{"Toyota (check digit 0)", VIN("JH4KA9659MCX12345"), '9'},
		{"Example with X check digit", VIN("1M8GDM9AXKP042788"), 'X'},
	}

	for _, tt := range tests {
		t.Run(
			tt.Name,
			func(t *testing.T) {
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
	
}
