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

///

type getListMetasModel struct {
	Lists []*BringListMeta
}

func (a *AuthBringo) GetListMetas() (lists []*BringListMeta, err error) {
	url := fmt.Sprintf("%s/bringusers/%s/lists", a.Base, a.Auth.UUID)
	var resp *req.Resp
	if resp, err = req.Get(url, a.createHeaders()); err != nil {
		return
	}
	var model getListMetasModel
	if err = resp.ToJSON(&model); err != nil {
		return
	}
	return model.Lists, nil
}

func (a *AuthBringo) GetLists() (lists []*BringListExpensive, err error) {
	var meta []*BringListMeta
	if meta, err = a.GetListMetas(); err != nil {
		return
	}
	lists = make([]*BringListExpensive, len(meta))
	for i, m := range meta {
		var list *BringList
		if list, err = a.GetListByMeta(m); err != nil {
			return
		}
		lists[i] = &BringListExpensive{
			BringListMeta: m,
			BringList:     list,
		}
	}
	return
}

///

func (a *AuthBringo) GetListByMeta(meta *BringListMeta) (list *BringList, err error) {
	return a.GetList(meta.UUID)
}

func (a *AuthBringo) GetList(listUUID string) (list *BringList, err error) {
	url := fmt.Sprintf("%s/bringlists/%s", a.Base, listUUID)
	var resp *req.Resp
	if resp, err = req.Get(url, a.createHeaders()); err != nil {
		return
	}
	err = resp.ToJSON(&list)
	return
}
