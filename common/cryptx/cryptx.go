package cryptx

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	encodedKey := base64.StdEncoding.EncodeToString(dk)
	return encodedKey, err
}

func PasswordVerify(salt, password, hash string) bool {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	encodedKey := base64.StdEncoding.EncodeToString(dk)
	return encodedKey == hash
}
