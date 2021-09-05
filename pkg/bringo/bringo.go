package bringo

import (
	"errors"
	"github.com/imroc/req"
)

const (
	DefaultBaseURL = "https://api.getbring.com/rest/v2/"
)

///

func createHeaders(h req.Header) (resp req.Header) {
	resp = DefaultHeaders
	for k, v := range h {
		resp[k] = v
	}
	return
}

///

var (
	ErrNotInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidResponse       = errors.New("invalid response")
)

type Bringo struct {
	Base string `json:"base"`
}

type AuthBringo struct {
	*Bringo `json:"bringo"`
	Auth    *bringAuth `json:"auth"`
}

///

func New() *Bringo {
	return NewWithBaseURL(DefaultBaseURL)
}

func NewWithBaseURL(base string) *Bringo {
	return &Bringo{
		Base: base,
	}
}

func NewWithLogin(email, pass string) (a *AuthBringo, err error) {
	guest := New()
	a, err = guest.Login(email, pass)
	return
}
