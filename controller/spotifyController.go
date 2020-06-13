package controller

import (
	"encoding/json"
	"fmt"
	"go-sample/go-spotify/models"
	"go-sample/go-spotify/utility"
	"io/ioutil"
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

//GetAccessTokenController.. Fetch Access Token Details
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

//SpotifySearchController... Spotify Search
func SpotifySearchController(r *http.Request) (returnData utility.ReturnJson) {

	returnData.Code = utility.CODE_400
	returnData.Message = utility.FAIL

	SpotifySearchQueryParms := utility.SpotifySearchQuery{}

	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &SpotifySearchQueryParms)
	if err != nil {
		returnData.Data = "Invalid Request Data"
		fmt.Println("Error in Unmarshalling ", err)
		return
	}

	//fetch access token from DB
	v, err := models.GetSpotifyByClientId(utility.ClientId)
	if err != nil {
		returnData.Data = "Error Fetching Access Token"
		fmt.Println("Error in fetching Access Token from DB ", err)
		return
	}

	if v == nil {
		returnData.Data = "No Access Token in DB"
		fmt.Println("NO Access Token from DB ")
		return
	}

	res, err := utility.SpotifySearch(v.AccessToken, SpotifySearchQueryParms)
	if err != nil {
		returnData.Data = "Error In Spotify Search"
		fmt.Println("Error In Spotify Search ", err)
		return
	}

	returnData.Code = utility.CODE_200
	returnData.Message = utility.SUCCESS
	returnData.Data = res

	return
}
