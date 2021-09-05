package bringo

import "github.com/imroc/req"

func (b *Bringo) Login(email, password string) (a *AuthBringo, err error) {
	var resp *req.Resp
	if resp, err = req.Post(b.Base+"bringauth", req.Param{
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
	a = &AuthBringo{Bringo: b}
	// parse json response to `AuthBringo` struct
	if err = resp.ToJSON(&a.Auth); err != nil {
		return
	}
	return
}
