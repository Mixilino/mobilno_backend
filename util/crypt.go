package util

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(oldPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(oldPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePasswords(oldPassword, newPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword)) == nil
}
