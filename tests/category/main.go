package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var ApiEndpoint = "http://localhost:1323/api/v0/category"

func main() {
	res := getCategories(ApiEndpoint)
	fmt.Println(res)
}

func getCategories(apiEndpoint string) string {
	req := createRequest(apiEndpoint)
	res := doRequest(req)
	return res
}

func createRequest(apiEndpoint string) *http.Request {
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer testtest")
	return req
}

func doRequest(req *http.Request) string {
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody := makeResponseBody(res)
	return resBody
}

func makeResponseBody(res *http.Response) string {
	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(resBodyBytes)
}
