package controller

import (
	"fmt"
	"go-sample/go-spotify/utility"
	"net/http"
	"os/exec"
	"strings"

	"github.com/spf13/cast"
	"github.com/zmb3/spotify"
)

func SpotifyAuthorizeController() string {

	auth := spotify.NewAuthenticator(utility.RedirectURI, spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(utility.ClientId, utility.ClientSecret)

	url := auth.AuthURL(utility.State)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	exec.Command("xdg-open", url).Start()
	return "Success"

}

func SpotifyController(r *http.Request) map[string]interface{} {

	code := r.URL.Query().Get(utility.ConstCode)

	fmt.Println("Code = ", code)

	auth := utility.GetBase64EncodedValue(utility.ClientId + utility.COLON + utility.ClientSecret)
	headers := utility.SetAndGetHeaders(utility.ConstFormEncoded, utility.BasicAuthType, auth)

	uri := strings.ReplaceAll(utility.RedirectURI, utility.SLASH, utility.UrlSlashReplace)
	uri = strings.ReplaceAll(uri, utility.COLON, utility.UrlColonReplace)

	payload := utility.GRANT_TYPE + utility.EQUALTO + utility.AUTHORIZATION_CODE + utility.AMPERSAND + utility.ConstCode + utility.EQUALTO + code + utility.AMPERSAND + utility.REDIRECT_URI + utility.EQUALTO + uri

	fmt.Println("payload = ", payload)

	resp, err := utility.SendRequest(utility.POST_METHOD, utility.Token_URL, payload, headers)

	if err != nil {
		fmt.Println("Error in SpotifyController ", err)
		return nil
	}

	resMap := cast.ToStringMap(string(resp))

	return resMap

}
