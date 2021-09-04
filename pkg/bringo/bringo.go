package bringo

import (
	"errors"
	"github.com/imroc/req"
)

const (
	DefaultBaseURL = "https://api.getbring.com/rest/v2/"
)

var (
	ErrNotInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidResponse       = errors.New("invalid response")
)

type Bringo struct {
	base string
}

type AuthBringo struct {
	auth *bringAuth
	Dog  *expireDog
}

///

func New() *Bringo {
	return NewWithBaseURL(DefaultBaseURL)
}

func NewWithBaseURL(base string) *Bringo {
	return &Bringo{
		base: base,
	}
}

func NewWithLogin(email, pass string) (a *AuthBringo, err error) {
	guest := New()
	a, err = guest.Login(email, pass)
	return
}

///

func (a *AuthBringo) Close() error {
	if a.Dog != nil {
		a.Dog.Stop()
	}
	return nil
}

///

func (b *Bringo) Login(email, password string) (a *AuthBringo, err error) {
	var resp *req.Resp
	if resp, err = req.Post(b.base+"bringauth", req.Param{
		"email":    email,
		"password": password,
	}); err != nil {
		return
	}
	// status code should be 200
	if resp.Response().StatusCode != 200 {
		if resp.Response().StatusCode == 401 {
			err = ErrNotInvalidCredentials
		} else {
			err = ErrInvalidResponse
		}
		return
	}
	a = new(AuthBringo)
	// parse json response to `AuthBringo` struct
	if err = resp.ToJSON(&a.auth); err != nil {
		return
	}
	// update `Expires` field
	if a.auth.ExpiresIn > 0 {
		a.Dog = newExpireDog(a.auth.ExpiresIn)
	}
	return
}
