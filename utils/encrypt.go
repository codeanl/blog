package utils

import "golang.org/x/crypto/bcrypt"

func BcryptHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}

func BcryptCheck(plian, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plian))
	return err == nil
}
