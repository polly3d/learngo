package tool

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

//FetchURLs fetch url from markdown file
func FetchURLs(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r, _ := regexp.Compile(`https?:\/\/github\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

	urls := r.FindAllString(string(data), -1)

	//remove https://github.com from urls
	len := len(urls)
	for i := 0; i < len; i++ {
		url := urls[i]
		url = strings.Replace(url, "https://github.com/", "", -1)
		urls[i] = url
	}
	fmt.Println(urls)

	return urls
}
