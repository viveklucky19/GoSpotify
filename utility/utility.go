package utility

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"encoding/base64"

	"encoding/json"

	"github.com/spf13/viper"
)

const (
	POST_METHOD = "POST"
)

//end points
const (
	AUTHORIZE_END_POINT = "/authorize"
	CALLBACK_END_POINT  = "/callback"
)
const (
	COLON           = ":"
	SLASH           = "/"
	EQUALTO         = "="
	AMPERSAND       = "&"
	UrlSlashReplace = "%2F"
	UrlColonReplace = "%3A"
)
const (
	ConstSpotify       = "spotify"
	ConstClientId      = "client_id"
	ConstClientSecret  = "client_secret"
	ConstCode          = "code"
	ConstFormEncoded   = "application/x-www-form-urlencoded"
	ConstAuthorization = "Authorization"
	BasicAuthType      = "Basic "
	GRANT_TYPE         = "grant_type"
	AUTHORIZATION_CODE = "authorization_code"
	REDIRECT_URI       = "redirect_uri"
	RedirectURI        = "http://localhost:8080/callback"
	ConstContentType   = "Content-Type"
	Token_URL          = "https://accounts.spotify.com/api/token"
)

var (
	State        = "vivek_spotify"
	ClientId     string
	ClientSecret string
)

//SetUpViper... set up viper
func SetUpViper(configPath, configName string) {

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error in read env file ", err)
	}

}

func SetAndGetHeaders(contentType, authType, auth string) map[string]string {
	headers := make(map[string]string)
	headers[ConstContentType] = contentType
	headers[ConstAuthorization] = authType + auth
	return headers
}

func GetBase64EncodedValue(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

//SendRequest.. Common Func to send hhtp request
func SendRequest(method, url string, reqdata interface{}, headers map[string]string) ([]byte, error) {

	//create payload and client
	payload := strings.NewReader(reqdata.(string))
	fmt.Println("Request Body :   ", reqdata.(string))
	client := &http.Client{}

	//create request
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println("Request error: ", err)
		return nil, err
	}

	//add headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	//log request
	log.Printf("request: %+v ", *req)

	//make call
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Response error: ", err)
		return nil, err
	}

	//convert to byte array
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Response Body error: ", err)
		return nil, err
	}

	res.Body.Close()
	//log response body
	fmt.Println("Response body : ", string(body))
	return body, err
}

//ReturnResponse... common function to set response
func ReturnResponse(w http.ResponseWriter, data interface{}) {
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
