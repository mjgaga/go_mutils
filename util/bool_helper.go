package util

func ConvertBoolToUint8(b bool) uint8 {
	var bitSetVar uint8
	if b {
		bitSetVar = 1
	}
	return bitSetVar
}

func ConvertUint8ToBool(i uint8) bool {
	var b bool
	if string(i) == "1" {
		b = true
	}
	return b
}
