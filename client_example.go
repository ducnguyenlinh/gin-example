package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func basicAuth(url string, info map[string]string) {
	var name = info["name"]
	var password = info["password"]

	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}

	request, err := http.NewRequest("GET", url, nil)
	check(err)
	request.SetBasicAuth(name, password)

	result, err := client.Do(request)
	check(err)

	bodyText, err := ioutil.ReadAll(result.Body)
	s := string(bodyText)

	fmt.Println(s)
}

func main() {
	info := make(map[string]string)

	info["name"] = os.Args[2]
	info["password"] = os.Args[3]

	url := os.Args[1]

	basicAuth(url, info)
}
