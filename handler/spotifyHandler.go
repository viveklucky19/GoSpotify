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
		utility.ReturnResponse(w, controller.SpotifyAuthorizeController())
	case strings.Contains(r.URL.Path, utility.CALLBACK_END_POINT):
		fmt.Println("CALLBACK_END_POINT handler")
		utility.ReturnResponse(w, controller.SpotifyController(r))

	}

}
