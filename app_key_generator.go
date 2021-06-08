package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/thanhpk/randstr"
)

func main() {
	appKey := generateAppKey()
	setAppKey(".env", appKey)
	setAppKey("cpanel/.env", appKey)
}

func generateAppKey() string {
	appKey := randstr.String(20)
	return appKey
}

func setAppKey(path string, appKey string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	line := fmt.Sprintf("APP_KEY=%s", appKey)
	rep := regexp.MustCompile(`APP_KEY=.*`)
	newContents := rep.ReplaceAllString(string(file), line)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
