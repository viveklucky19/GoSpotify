package handler

import (
	"fmt"
	"go-sample/go-spotify/controller"
	"go-sample/go-spotify/utility"
	"net/http"
	"strings"
)

//SpotifyHandler... Handler Function for Spotify
func SpotifyHandler(w http.ResponseWriter, r *http.Request) {

	switch true {
	case strings.Contains(r.URL.Path, utility.AUTHORIZE_END_POINT):
		fmt.Println("AUTHORIZE_END_POINT handler")
		utility.ReturnResponse(w, controller.AuthorizeController())
	case strings.Contains(r.URL.Path, utility.CALLBACK_END_POINT):
		fmt.Println("CALLBACK_END_POINT handler")
		utility.ReturnResponse(w, controller.CallbackController(r))
	case strings.Contains(r.URL.Path, utility.GET_ACCESS_TOKEN_END_POINT):
		fmt.Println("GET_ACCESS_TOKEN_END_POINT handler")
		utility.ReturnResponse(w, controller.GetAccessTokenController())
	case strings.Contains(r.URL.Path, utility.SPOTIFY_SEARCH):
		fmt.Println("SPOTIFY_SEARCH handler")
		utility.ReturnResponse(w, controller.SpotifySearchController(r))

	default:
		fmt.Println("Inavlid URL : ", r.URL)
	}

}
