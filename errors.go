package main

import "fmt"

type VINLengthError struct {
	Got  int
	Want int
}

func (e VINLengthError) Error() string {
	switch {
	case e.Got < e.Want:
		return fmt.Sprintf("VIN too short: got %d, want %d.", e.Got, e.Want)
	case e.Got > e.Want:
		return fmt.Sprintf("VIN too long: got %d, want %d.", e.Got, e.Want)
	default:
		return fmt.Sprintf("VIN length mismatch: got %d, want %d", e.Got, e.Want)
	}
}

type VINWrongCharacterError struct {
	Char  byte
	Index uint32
}

func (e VINWrongCharacterError) Error() string {
	return fmt.Sprintf("invalid VIN character %q at position %d", e.Char, e.Index)
}

type VINChecksumError struct {
	Expected byte
	Actual   byte
}

func (e VINChecksumError) Error() string {
	return fmt.Sprintf("invalid VIN checksum: expected %q, got %q", e.Expected, e.Actual)
}
