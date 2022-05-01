package types

import (
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type Password string
type Encrypted Password

var tokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

func (p *Password) Validate(password []byte) bool {

	byteHash := []byte(*p)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		return false
	}

	return true
}

func (p *Password) Encrypt() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(*p), 13)
	if err != nil {
		return err
	}

	*p = Password(hash)

	return err
}
