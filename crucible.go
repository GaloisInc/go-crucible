package crucible

func Assume(b bool, file string, msg string) {
	if !b {
		panic("Assume failure in file " + file + ": " + msg)
	}
}

func Assert(b bool, file string, msg string) {
	if !b {
		panic("Assert failure in file " + file + ": " + msg)
	}
}

func FreshInt() int {
	return 0
}

func FreshInt8() int8 {
	return 0
}

func FreshInt16() int16 {
	return 0
}

func FreshInt32() int32 {
	return 0
}

func FreshInt64() int64 {
	return 0
}

func FreshUint() uint {
	return 0
}

func FreshUint8() uint8 {
	return 0
}

func FreshUint16() uint16 {
	return 0
}

func FreshUint32() uint32 {
	return 0
}

func FreshUint64() uint64 {
	return 0
}

func FreshFloat32() float32 {
	return 0.0
}

func FreshFloat64() float64 {
	return 0.0
}

func FreshString() string {
	return ""
}

// func FreshComplex64() complex64 {
// 	return 0i
// }

// func FreshComplex128() complex128 {
// 	return 0i
// }
