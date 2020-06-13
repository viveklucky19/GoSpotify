package main

import (
	"log"
	"net/http"

	"go-sample/go-spotify/handler"
	"go-sample/go-spotify/utility"

	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
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
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/vivek?charset=utf8")
	orm.Debug = true
}

func main() {

	//URL handling
	http.HandleFunc(utility.AUTHORIZE_END_POINT, handler.SpotifyHandler)
	http.HandleFunc(utility.CALLBACK_END_POINT, handler.SpotifyHandler)
	http.HandleFunc(utility.GET_ACCESS_TOKEN_END_POINT, handler.SpotifyHandler)
	http.HandleFunc(utility.SPOTIFY_SEARCH, handler.SpotifyHandler)

	//server
	log.Fatal(http.ListenAndServe(Port, nil))

}
