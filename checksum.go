package main

var vinTranslit = func() [256]uint8 {
	var t [256]uint8
	// A=1..H=8
	t['A'], t['B'], t['C'], t['D'], t['E'], t['F'], t['G'], t['H'] = 1, 2, 3, 4, 5, 6, 7, 8
	// J=1..N=5
	t['J'], t['K'], t['L'], t['M'], t['N'] = 1, 2, 3, 4, 5
	// P=7
	t['P'] = 7
	// R=9
	t['R'] = 9
	// S=2..Z=9 (skipping I,O,Q)
	t['S'], t['T'], t['U'], t['V'], t['W'], t['X'], t['Y'], t['Z'] = 2, 3, 4, 5, 6, 7, 8, 9
	return t
}()

var vinWeights = [17]uint8{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

func ComputeVINCheckDigit(v VIN) byte {
	sum := 0
	for i := 0; i < 17; i++ {
		ch := v[i]
		val := int(0)

		if ch >= '0' && ch <= '9' {
			val = int(ch - '0')
		} else {
			val = int(vinTranslit[ch])
		}

		sum += val * int(vinWeights[i])
	}

	mod := sum % 11
	if mod == 10 {
		return 'X'
	}
	return byte('0' + mod)
}
