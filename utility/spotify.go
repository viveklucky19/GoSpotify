package utility

import (
	"github.com/zmb3/spotify"
)

var (
	State        = "vivek_spotify"
	ClientId     string
	ClientSecret string
)

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func SpotifyAuthenticateURL() (url string) {
	//Create Auhenticator
	auth := spotify.NewAuthenticator(REDIRECT_URL, spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(ClientId, ClientSecret)
	//hit auth URL https://accounts.spotify.com/authorize
	url = auth.AuthURL(State)
	return
}
