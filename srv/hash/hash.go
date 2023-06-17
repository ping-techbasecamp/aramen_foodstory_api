package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(plain string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.MinCost)
	return string(hash)
}

func Compare(plain string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
