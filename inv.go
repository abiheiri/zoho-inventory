package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("auth.conf") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	s := strings.Split(string(b), ",")

	username, apikey := s[0], s[1]
	// fmt.Println(username, apikey)

	getToken(username, apikey)
}

func getToken(username, apikey string) {
	url := "https://accounts.zoho.com/apiauthtoken/nb/create?SCOPE=ZohoInventory/inventoryapi&EMAIL_ID=" + username + "&PASSWORD=" + apikey
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// s := strings.Split(string(body),"=")
	// fmt.Println(string(body))
	// fmt.Println(s[1])

	var re = regexp.MustCompile(`AUTHTOKEN=(?P<token>\S+)`)
	var str = string(body)
	match := re.FindStringSubmatch(str)
	fmt.Printf("TOKEN: %s", match[1])

	// r, err := regexp.MatchString(`AUTHTOKEN=(.*)`, string(body))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(r))
}
