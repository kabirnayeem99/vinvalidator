package main

import (
	"math/rand"
	"time"
)

const vinCharset = "ABCDEFGHJKLMNPRSTUVWXYZ0123456789"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateMockVIN() VIN {
	var vin [17]byte

	for i := 0; i < 17; i++ {
		vin[i] = vinCharset[rng.Intn(len(vinCharset))]
	}

	vin[8] = '0'

	check := ComputeVINCheckDigit(VIN(vin[:]))
	vin[8] = check

	return VIN(vin[:])
}
