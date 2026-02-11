package main

import "testing"

func TestValidateVIN_WrongLength(t *testing.T) {
	tests := []struct {
		Name  string
		Value VIN
		err   error
	}{
		{
			Name:  "Smaller VIN Length should give error",
			Value: "1HGCM82633A00435",
			err:   VINLengthError{Got: 16, Want: ValidVINLength},
		},
		{
			Name:  "Proper VIN Length should not give any error",
			Value: "1HGCM82633A004352",
			err:   nil,
		},
		{
			Name:  "Larger VIN Length should give error",
			Value: "1HGCM82633A0043525",
			err:   VINLengthError{Got: 18, Want: ValidVINLength},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			g := ValidateVIN(tt.Value)
			w := tt.err

			if g != w {
				t.Fatalf("ValidateVIN(%s) = %v; want = %v", tt.Value, g, w)
			}

		})
	}
}

func TestValidateVIN_WrongCharacter(t *testing.T) {
	tests := []struct {
		Name  string
		Value VIN
		err   error
	}{
		{
			Name:  "I should give error",
			Value: "IHGCM82633A004352",
			err:   VINWrongCharacterError{Char: 'I', Index: 0},
		},

		{
			Name:  "I, O, Q should give error",
			Value: "1HGCM82633AO04352",
			err:   VINWrongCharacterError{Char: 'O', Index: 11},
		},

		{
			Name:  "Q should give error",
			Value: "1HGCM82633A0Q4352",
			err:   VINWrongCharacterError{Char: 'Q', Index: 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			w := tt.err
			g := ValidateVIN(tt.Value)

			if g != w {
				t.Fatalf("ValidateVIN(%s) = %v; want = %v", tt.Value, g, w)
			}

		})
	}

}
