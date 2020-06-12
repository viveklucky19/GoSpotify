package main

import (
	"log"
	"net/http"

	"go-sample/go-spotify/handler"
	"go-sample/go-spotify/utility"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

const (
	Port       = ":8080"
	ConfigPath = "."
	ConfigName = "env"
)

func init() {
	utility.SetUpViper(ConfigPath, ConfigName)
	utility.ClientId = cast.ToString(viper.Get(utility.ConstSpotify + "." + utility.ConstClientId))
	utility.ClientSecret = cast.ToString(viper.Get(utility.ConstSpotify + "." + utility.ConstClientSecret))
}

func main() {

	http.HandleFunc(utility.AUTHORIZE_END_POINT, handler.SpotifyHandler)
	http.HandleFunc(utility.CALLBACK_END_POINT, handler.SpotifyHandler)

	log.Fatal(http.ListenAndServe(Port, nil))

}
