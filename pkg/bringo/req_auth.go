package bringo

import (
	"fmt"
	"github.com/imroc/req"
)

var DefaultHeaders = req.Header{
	"X-BRING-API-KEY":       "cof4Nc6D8saplXjE3h3HXqHH8m7VU2i1Gs0g85Sp",
	"X-BRING-CLIENT":        "webApp",
	"X-BRING-CLIENT-SOURCE": "webApp",
	"X-BRING-COUNTRY":       "DE",
}

func (a *AuthBringo) createHeaders() req.Header {
	h := DefaultHeaders
	h["H-BRING-USER-UUID"] = a.Auth.UUID
	h["Authorization"] = fmt.Sprintf("%s %s", a.Auth.TokenType, a.Auth.AccessToken)
	return h
}

type loadListsModel struct {
	Lists []*BringList
}

func (a *AuthBringo) LoadLists() (lists []*BringList, err error) {
	url := fmt.Sprintf("%s/bringusers/%s/lists", a.Base, a.Auth.UUID)
	var resp *req.Resp
	if resp, err = req.Get(url, a.createHeaders()); err != nil {
		return
	}
	var model loadListsModel
	if err = resp.ToJSON(&model); err != nil {
		return
	}
	return model.Lists, nil
}
