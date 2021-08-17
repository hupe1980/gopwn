package tubes

// Bytes takes type interface{} and converts it to []byte, if it can't convert
// to []byte it will panic
func Bytes(t interface{}) []byte {
	switch x := t.(type) {
	case string:
		return []byte(x)
	case []byte:
		return x
	case byte:
		return []byte{x}
	case rune:
		return []byte(string(x))
	default:
		panic("failed to convert t to type []byte")
	}
}
