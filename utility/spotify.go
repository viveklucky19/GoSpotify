package utility

import (
	"encoding/json"
	"fmt"
	"go-sample/go-spotify/models"
	"strings"

	"github.com/spf13/cast"
	"github.com/zmb3/spotify"
)

var (
	State        = "vivek_spotify"
	ClientId     string
	ClientSecret string
)

const (
	TOKEN_URL  = "https://accounts.spotify.com/api/token"
	SEARCH_URL = "https://api.spotify.com/v1/search"
)

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type SpotifySearchQuery struct {
	Query           string `json:"query"`
	Type            string `json:"type"`
	Market          string `json:"market"`
	Limit           string `json:"limit"`
	Offset          string `json:"offset"`
	IncludeExternal string `json:"include_external"`
}

func SpotifyAuthenticateURL() (url string) {
	//Create Auhenticator
	auth := spotify.NewAuthenticator(REDIRECT_URL, spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(ClientId, ClientSecret)
	//hit auth URL https://accounts.spotify.com/authorize
	url = auth.AuthURL(State)
	return
}

//hit get token POST request https://accounts.spotify.com/api/token
func SpotifyCallback(code string) (SpotifyTokenDetails models.Spotify, err error) {

	auth := GetBase64EncodedValue(ClientId + COLON + ClientSecret)
	headers := SetAndGetHeaders(ConstFormEncoded, BasicAuthType, auth)

	uri := strings.ReplaceAll(REDIRECT_URL, SLASH, UrlSlashReplace)
	uri = strings.ReplaceAll(uri, COLON, UrlColonReplace)

	payload := ConstGrantType + EQUALTO + ConstAuthorizationCode + AMPERSAND + ConstCode + EQUALTO + code + AMPERSAND + ConstRedirectUri + EQUALTO + uri
	// fmt.Println("payload = ", payload)

	//hit get toket POST request https://accounts.spotify.com/api/token
	resp, err := SendRequest(POST_METHOD, TOKEN_URL, payload, headers)
	if err != nil {
		fmt.Println("Error in SpotifyCallback ", err)
		return
	}

	resMap := cast.ToStringMap(string(resp))
	if _, ok := resMap["error"]; !ok {
		//unmarshall it to Spotify struct
		var SpotifyData SpotifyToken
		err = json.Unmarshal(resp, &SpotifyData)
		if err != nil {
			fmt.Println("Error in Unmarshalling ", err)
			return
		}
		fmt.Printf("data =%+v ", SpotifyData)

		SpotifyTokenDetails.ClientId = ClientId
		SpotifyTokenDetails.AccessToken = SpotifyData.AccessToken
		SpotifyTokenDetails.ExpiresIn = SpotifyData.ExpiresIn
		SpotifyTokenDetails.RefreshToken = SpotifyData.RefreshToken
		SpotifyTokenDetails.Scope = SpotifyData.Scope
		SpotifyTokenDetails.TokenType = SpotifyData.TokenType

	}
	return
}

//https://api.spotify.com/v1/search
func SpotifySearch(accessToken string, queryData SpotifySearchQuery) (response map[string]interface{}, err error) {

	headers := SetAndGetHeaders("", BearerAuthType, accessToken)

	mapValues := map[string]string{}
	mapValues["q"] = queryData.Query
	mapValues["type"] = queryData.Type
	mapValues["limit"] = queryData.Limit
	mapValues["offset"] = queryData.Offset
	mapValues["market"] = queryData.Market
	mapValues["include_external"] = queryData.IncludeExternal

	Url := SEARCH_URL + QUESTION_MARK

	for k, v := range mapValues {

		if v != "" {
			Url = Url + k + EQUALTO + v + AMPERSAND

		}
	}

	fmt.Println("URL =", Url)

	resp, err := SendRequest(GET_METHOD, Url, nil, headers)
	if err != nil {
		fmt.Println("Error in SpotifyCallback ", err)
		return
	}

	response = cast.ToStringMap(string(resp))

	return response, nil

}
