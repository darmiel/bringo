package bringo

type bringAuth struct {
	UUID          string `json:"uuid"`
	PublicUUID    string `json:"publicUuid"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	PhotoPath     string `json:"photoPath"`
	BringListUUID string `json:"bringListUUID"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	TokenType     string `json:"token_type"`
	ExpiresIn     int    `json:"expires_in"`
}

type BringList struct {
	UUID  string `json:"listUuid"`
	Name  string `json:"name"`
	Theme string `json:"theme"`
}
