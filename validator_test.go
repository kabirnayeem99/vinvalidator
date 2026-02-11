package main

import (
	"errors"
	"testing"
)

func TestValidateVIN_Length(t *testing.T) {
	tests := []struct {
		name    string
		vin     VIN
		wantErr error
	}{
		{
			name:    "too short",
			vin:     "1HGCM82633A00435",
			wantErr: VINLengthError{Got: 16, Want: ValidVINLength},
		},
		{
			name:    "valid length",
			vin:     "1HGCM82633A004352",
			wantErr: nil,
		},
		{
			name:    "too long",
			vin:     "1HGCM82633A0043525",
			wantErr: VINLengthError{Got: 18, Want: ValidVINLength},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ValidateVIN(tt.vin)
			want := tt.wantErr

			if got != want {
				t.Fatalf("ValidateVIN(%s) = %v; want = %v", tt.vin, got, want)
			}

		})
	}
}

func TestValidateVIN_WrongCharacter(t *testing.T) {
	tests := []struct {
		name string
		vin  VIN
		want *VINWrongCharacterError // nil => want no error
	}{
		{
			name: "rejects I",
			vin:  "IHGCM82633A004352",
			want: &VINWrongCharacterError{Char: 'I', Index: 0},
		},
		{
			name: "rejects O",
			vin:  "1HGCM82633AO04352",
			want: &VINWrongCharacterError{Char: 'O', Index: 11},
		},
		{
			name: "rejects Q",
			vin:  "1HGCM82633A0Q4352",
			want: &VINWrongCharacterError{Char: 'Q', Index: 12},
		},
		{
			name: "rejects non-alphanumeric",
			vin:  "1HGCM82633A00435?",
			want: &VINWrongCharacterError{Char: '?', Index: 16},
		},
		{
			name: "accepts valid VIN characters",
			vin:  "1HGCM82633A004352",
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateVIN(tt.vin)

			if tt.want == nil {
				if err != nil {
					t.Fatalf("ValidateVIN(%q) error = %v; want nil", tt.vin, err)
				}
				return
			}

			var got VINWrongCharacterError
			if !errors.As(err, &got) {
				t.Fatalf("ValidateVIN(%q) error = %T (%v); want VINWrongCharacterError", tt.vin, err, err)
			}

			if got.Char != tt.want.Char || got.Index != tt.want.Index {
				t.Fatalf(
					"ValidateVIN(%q) error = {Char:%q Index:%d}; want {Char:%q Index:%d}",
					tt.vin, got.Char, got.Index,
					tt.want.Char, tt.want.Index,
				)
			}
		})
	}
}


func TestValidateVIN_Checksum(t *testing.T) {
	tests := []struct {
		name string
		vin  VIN
		want *VINChecksumError // nil => want no error
	}{
		{
			name: "valid checksum",
			vin:  "1HGCM82633A004352",
			want: nil,
		},
		{
			name: "wrong checksum",
			vin:  "1HGCM82653A004352",
			want: &VINChecksumError{Expected: '3', Actual: '5'},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateVIN(tt.vin)

			if tt.want == nil {
				if err != nil {
					t.Fatalf("ValidateVIN(%q) error = %v; want nil", tt.vin, err)
				}
				return
			}

			var got VINChecksumError
			if !errors.As(err, &got) {
				t.Fatalf("ValidateVIN(%q) error = %T (%v); want VINChecksumError", tt.vin, err, err)
			}

			if got.Expected != tt.want.Expected || got.Actual != tt.want.Actual {
				t.Fatalf(
					"ValidateVIN(%q) checksum error = {Expected:%q Actual:%q}; want {Expected:%q Actual:%q}",
					tt.vin, got.Expected, got.Actual,
					tt.want.Expected, tt.want.Actual,
				)
			}
		})
	}
}

