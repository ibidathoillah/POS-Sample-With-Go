package types

import (
	"crypto/md5"
	"fmt"

	"github.com/go-chi/jwtauth"
)

type Password string
type Encrypted Password

var tokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

func (p *Password) Validate(password []byte) bool {

	if fmt.Sprintf("%s", *p) != fmt.Sprintf("%x", md5.Sum(password)) {
		return false
	}

	return true
}

func (p *Password) Encrypt() error {

	data := []byte(*p)
	*p = Password(fmt.Sprintf("%x", md5.Sum(data)))

	return nil
}
