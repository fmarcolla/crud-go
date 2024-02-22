package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(ud.GetPassword()), 6)
	ud.SetPassword(string(hash))
}

func (ud *userDomain) CheckPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(ud.GetPassword()))
	return err == nil
}
