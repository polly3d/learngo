package tool

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Ls() {
	dir, _ := os.Getwd()
	result, _ := ioutil.ReadDir(dir)
	for _, file := range result {
		fmt.Printf("%s\t", file.Name())
	}
}
