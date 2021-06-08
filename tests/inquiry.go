package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Inquiry struct {
	Message    string
	CategoryId int
}

func main() {
	apiEndpoint := "http://localhost:1323/api/v0/inquiry"

	inquiry := Inquiry{
		Message:    "hogehoge",
		CategoryId: 1,
	}

	reqBody, err := json.Marshal(inquiry)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(reqBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer hogehoge")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resBody))
}
