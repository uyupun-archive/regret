package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://localhost:1323/api/v0/inquiry"

	reqBody := strings.NewReader("{}")
	res, err := http.Post(url, "application/json", reqBody)
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
