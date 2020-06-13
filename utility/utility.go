package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"encoding/base64"

	"encoding/json"

	"github.com/spf13/viper"
)

type ReturnJson struct {
	Code    string
	Message string
	Data    interface{}
}

//SetUpViper... set up viper
func SetUpViper(configPath, configName string) {

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error in read env file ", err)
	}

}

//SetAndGetHeaders... set headers and return it
func SetAndGetHeaders(contentType, authType, auth string) map[string]string {
	headers := make(map[string]string)
	if contentType != "" {
		headers[ConstContentType] = contentType
	}
	headers[ConstAuthorization] = authType + auth
	return headers
}

//GetBase64EncodedValue...get base 64 encoded value
func GetBase64EncodedValue(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

//SendRequest.. Common Func to send http request
func SendRequest(method, url string, reqdata interface{}, headers map[string]string) ([]byte, error) {

	var payload io.Reader
	payload = nil
	//create payload and client

	if reqdata != nil {
		payload = strings.NewReader(reqdata.(string))
		fmt.Println("Request Body :   ", reqdata.(string))
	}

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
