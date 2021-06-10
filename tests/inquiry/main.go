package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/uyupun/regret/models"
)

func main() {

	inquiry := models.Inquiry{
		Subject:    "問い合わせテスト",
		Email:      "uyupun@gmail.com",
		CategoryId: 1,
		Text:       "インターネットが壊れました。",
	}

	apiEndpoint := "http://localhost:1323/api/v0/inquiry"
	res := postInquiry(apiEndpoint, inquiry)
	fmt.Println(res)
}

func postInquiry(apiEndpoint string, inquiry models.Inquiry) string {
	reqBody := makeRequestBody(inquiry)
	req := createRequest(apiEndpoint, reqBody)
	res := doRequest(req)
	return res
}

func makeRequestBody(inquiry models.Inquiry) *bytes.Reader {
	reqBodyBytes, err := json.Marshal(inquiry)
	if err != nil {
		panic(err)
	}

	reqBodyReader := bytes.NewReader(reqBodyBytes)
	return reqBodyReader
}

func createRequest(apiEndpoint string, reqBody *bytes.Reader) *http.Request {
	req, err := http.NewRequest("POST", apiEndpoint, reqBody)
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
