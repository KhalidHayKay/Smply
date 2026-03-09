package utils

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(num int64) string {
	if num == 0 {
		return "0"
	}

	base := int64(len(charset))
	result := ""

	for num > 0 {
		rem := num % base
		result = string(charset[rem]) + result
		num = num / base
	}

	return result
}
