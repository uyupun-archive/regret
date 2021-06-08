package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/thanhpk/randstr"
)

func main() {
	accessToken := generateAccessToken()
	setAccessToken(".env", accessToken)
	setAccessToken("cpanel/.env", accessToken)
}

func generateAccessToken() string {
	accessToken := randstr.String(20)
	return accessToken
}

func setAccessToken(path string, accessToken string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	line := fmt.Sprintf("ACCESS_TOKEN=%s", accessToken)
	rep := regexp.MustCompile(`ACCESS_TOKEN=.*`)
	newContents := rep.ReplaceAllString(string(file), line)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
