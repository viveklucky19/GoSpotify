package controller

import (
	"encoding/json"
	"fmt"
	"go-sample/go-spotify/models"
	"go-sample/go-spotify/utility"
	"net/http"
	"os/exec"
	"strings"

	"github.com/spf13/cast"
)

//SpotifyAuthorizeController... Authorization controller
func AuthorizeController() (returnData utility.ReturnJson) {

	returnData.Code = utility.CODE_200
	returnData.Message = utility.SUCCESS
	url := utility.SpotifyAuthenticateURL()
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	//open this URL in browser
	err := exec.Command("xdg-open", url).Start()

	if err != nil {
		fmt.Println("Error in exec command  = ", err)
		returnData.Code = utility.CODE_400
		returnData.Message = utility.FAIL
	}
	return
}

//SpotifyCallBackController...spotify callback controller
func CallbackController(r *http.Request) (returnData utility.ReturnJson) {

	var SpotifyTokenDetails models.Spotify
	returnData.Code = utility.CODE_400
	returnData.Message = utility.FAIL
	//get code query param from URL
	code := r.URL.Query().Get(utility.ConstCode)
	// fmt.Println("Code = ", code)

	auth := utility.GetBase64EncodedValue(utility.ClientId + utility.COLON + utility.ClientSecret)
	headers := utility.SetAndGetHeaders(utility.ConstFormEncoded, utility.BasicAuthType, auth)

	uri := strings.ReplaceAll(utility.REDIRECT_URL, utility.SLASH, utility.UrlSlashReplace)
	uri = strings.ReplaceAll(uri, utility.COLON, utility.UrlColonReplace)

	payload := utility.ConstGrantType + utility.EQUALTO + utility.ConstAuthorizationCode + utility.AMPERSAND + utility.ConstCode + utility.EQUALTO + code + utility.AMPERSAND + utility.ConstRedirectUri + utility.EQUALTO + uri
	// fmt.Println("payload = ", payload)

	//hit get toket POST request https://accounts.spotify.com/api/token
	resp, err := utility.SendRequest(utility.POST_METHOD, utility.TOKEN_URL, payload, headers)
	if err != nil {
		fmt.Println("Error in SpotifyCallBackController ", err)
		return
	}

	resMap := cast.ToStringMap(string(resp))
	if _, ok := resMap["error"]; !ok {
		//unmarshall it to Spotify model and save it to db

		var SpotifyData utility.SpotifyToken

		err = json.Unmarshal(resp, &SpotifyData)

		if err != nil {
			fmt.Println("Error in Unmarshalling ", err)
			return
		}

		fmt.Printf("data =%+v ", SpotifyData)

		SpotifyTokenDetails.ClientId = utility.ClientId
		SpotifyTokenDetails.AccessToken = SpotifyData.AccessToken
		SpotifyTokenDetails.ExpiresIn = SpotifyData.ExpiresIn
		SpotifyTokenDetails.RefreshToken = SpotifyData.RefreshToken
		SpotifyTokenDetails.Scope = SpotifyData.Scope
		SpotifyTokenDetails.TokenType = SpotifyData.TokenType
		err = models.UpdateIfExistElseInsert(&SpotifyTokenDetails)
		if err != nil {
			fmt.Println("Error in UpdateIfExistElseInsert ", err)
			return
		}
	}

	returnData.Code = utility.CODE_200
	returnData.Message = utility.SUCCESS
	returnData.Data = SpotifyTokenDetails

	return

}

func GetAccessTokenController() (returnData utility.ReturnJson) {

	v, err := models.GetSpotifyByClientId(utility.ClientId)
	if err != nil {
		returnData.Code = utility.CODE_400
		returnData.Message = utility.FAIL
	}

	returnData.Code = utility.CODE_200
	returnData.Message = utility.SUCCESS
	returnData.Data = v

	return
}
