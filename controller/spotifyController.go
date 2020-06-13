package controller

import (
	"fmt"
	"go-sample/go-spotify/models"
	"go-sample/go-spotify/utility"
	"net/http"
	"os/exec"
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
	var err error
	returnData.Code = utility.CODE_400
	returnData.Message = utility.FAIL
	// get code query param from URL
	code := r.URL.Query().Get(utility.ConstCode)

	SpotifyTokenDetails, err = utility.SpotifyCallback(code)
	if err != nil {
		fmt.Println("Error in SpotifyCallBackController ", err)
		return
	}
	err = models.UpdateIfExistElseInsert(&SpotifyTokenDetails)
	if err != nil {
		fmt.Println("Error in UpdateIfExistElseInsert ", err)
		return
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

func SpotifySearchController(r *http.Request) (returnData utility.ReturnJson) {

	return
}
