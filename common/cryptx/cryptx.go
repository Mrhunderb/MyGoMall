package cryptx

import "golang.org/x/crypto/scrypt"

func PasswordEncrypt(salt, password string) (string, error) {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	return string(dk), nil
}

func PasswordVerify(salt, password, hash string) bool {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	return string(dk) == hash
}
