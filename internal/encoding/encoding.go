package encoding

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-"

func Encode(number int64) string {
	var code []byte
	for number > 0 {
		mod := number % int64(len(chars))
		number /= int64(len(chars))
		code = append(code, chars[mod])
	}
	return string(code)
}
