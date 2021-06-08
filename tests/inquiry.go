package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Inquiry struct {
	Title      string
	Email      string
	CategoryId int
	Message    string
}

func main() {

	inquiry := Inquiry{
		Title:      "問い合わせテスト",
		Email:      "uyupun@gmail.com",
		CategoryId: 1,
		Message:    "インターネットが壊れました。",
	}

	apiEndpoint := "http://localhost:1323/api/v0/inquiry"
	reqBody := makeRequestBody(inquiry)
	req := createRequest(apiEndpoint, reqBody)

	res := postInquiry(req)
	fmt.Println(res)
}

func makeRequestBody(inquiry Inquiry) *bytes.Reader {
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
	req.Header.Set("Authorization", "Bearer hogehoge")
	return req
}

func postInquiry(req *http.Request) string {
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
