package security

import "golang.org/x/crypto/bcrypt"

func Hash(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	if err != nil {
		return ""
	}
	return string(hash)
}

func Compare(hash, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
