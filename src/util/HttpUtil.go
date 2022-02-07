package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func MappingUrlParameter(originalUrl string, paramMapper map[string]string) string {
	returnUrl := originalUrl
	log.Println("originalUrl= ", originalUrl)
	if paramMapper != nil {
		for key, replaceValue := range paramMapper {
			returnUrl = strings.Replace(returnUrl, key, replaceValue, -1)
			// fmt.Println("Key:", key, "=>", "Element:", replaceValue+":"+returnUrl)
		}
	}
	log.Println("returnUrl= ", returnUrl)
	return returnUrl
}

// http 호출
func CommonHttp(url string, json []byte, httpMethod string) (*http.Response, error) {

	authInfo := AuthenticationHandler()

	log.Println("CommonHttp "+httpMethod+", ", url)
	// log.Println("authInfo ", authInfo)
	client := &http.Client{}
	req, err1 := http.NewRequest(httpMethod, url, bytes.NewBuffer(json))
	if err1 != nil {
		panic(err1)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Content-Type", "application/json")

	req.Header.Add("Authorization", authInfo)

	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	resp, err := client.Do(req) // err 자체는 nil 이고 resp 내에 statusCode가 500임...

	return resp, err
}

// ajax 호출할 때 header key 생성
func AuthenticationHandler() string {

	// conf 파일에 정의
	//api_username := os.Getenv("API_USERNAME")
	//api_password := os.Getenv("API_PASSWORD")
	api_username := "default"
	api_password := "default"

	//The header "KEY: VAL" is "Authorization: Basic {base64 encoded $USERNAME:$PASSWORD}".
	apiUserInfo := api_username + ":" + api_password
	encA := base64.StdEncoding.EncodeToString([]byte(apiUserInfo))
	//req.Header.Add("Authorization", "Basic"+encA)
	return "Basic " + encA

}
