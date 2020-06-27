package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
    b, err := ioutil.ReadFile("auth.conf") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    s := strings.Split(string(b),",")
    username, key := s[0], s[1]
    fmt.Println(username, key)
}