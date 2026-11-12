package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		randCharsetIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(fmt.Sprintf("Не удалось сгенерировать случайное число: %s", err.Error()))
		}
		b[i] = charset[randCharsetIndex.Int64()]
	}

	return string(b)
}

// RandomUUID генерирует случайный UUID v4, например для id несуществующего объявления
func RandomUUID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("Не удалось сгенерировать UUID: %s", err.Error()))
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
