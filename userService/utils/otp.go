package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() (string, error) {
	seed := "0123456789"
	otpLength := 4
	byteSlice := make([]byte, 4)

	for i := 0; i < otpLength; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}
