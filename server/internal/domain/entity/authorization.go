package entity

type Authorization struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
