package main

type VIN string

const ValidVINLength = 17

var validVinCharacterTable = func() [256]bool {
	var t [256]bool
	for _, c := range "ABCDEFGHJKLMNPRSTUVWXYZ0123456789" {
		t[byte(c)] = true
	}
	return t
}()

func ValidateVIN(v VIN) error {
	vinLength := len(string(v))

	if vinLength != ValidVINLength {
		return VINLengthError{Got: vinLength, Want: ValidVINLength}
	}

	for i := range vinLength {
		if !validVinCharacterTable[v[i]] {
			return VINWrongCharacterError{Char: v[i], Index: uint32(i)}
		}
	}

	expected := ComputeVINCheckDigit(v)
	actual := v[8]
	if actual != expected {
		return VINChecksumError{Expected: expected, Actual: actual}
	}

	return nil
}
