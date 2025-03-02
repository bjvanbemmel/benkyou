package utils

var AccessToken string

func NewAccessToken() string {
	AccessToken = Hash(RandomString())
	return AccessToken
}
