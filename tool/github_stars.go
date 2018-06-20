package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Repos infomation
type Repos struct {
	Name  string `json:"name"`
	Stars int    `json:"stargazers_count"`
}

//FetchStars is to fetch stars from github
func FetchStars(repos string) Repos {
	apiURL := "https://api.github.com/repos"
	apiURL += "/" + repos
	fmt.Println(apiURL)

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if response.Status != "200" {
		//fmt.Println(string(body))
		//os.Exit(1)
	}

	data := &Repos{}
	err = json.Unmarshal(body, data)
	return *data
}
