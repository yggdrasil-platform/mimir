package model

type Token struct {
	AccessToken string `json:"accessToken" redis:"access_token"`
	ExpiresIn int `json:"expiresIn" redis:"expires_in"`
  TokenType string `json:"tokenType" redis:"token_type"`
	Id int `json:"id"`
}
