package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	fmt.Println(string(body))
	// fmt.Println(s[1])

	//r := regexp.MustCompile(`(?P<first>.*)`)
	// fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
	//fmt.Printf(string(body), r.SubexpNames())

}
