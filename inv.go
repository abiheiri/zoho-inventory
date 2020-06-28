package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
    "strings"
    "flag"
)

func main() {

	f, err := ioutil.ReadFile("auth.conf")
	if err != nil {
		fmt.Print(err)
	}

	s := strings.Split(string(f), ",")

	username, apikey := s[0], s[1]
	// fmt.Println(username, apikey)

    // Is there an existing token?
	t, err := ioutil.ReadFile("token.txt")
	if err != nil {
		getToken(username, apikey)
	}

	token := (string(t))
	fmt.Println(token)

    // Create flags
    // inv -g -i=all
    getPtr := flag.String("g", "get", "a string")
    // numbPtr := flag.Int("numb", 42, "an int")
    // boolPtr := flag.Bool("fork", false, "a bool")

    // var svar string
    // flag.StringVar(&svar, "svar", "bar", "a string var")
    flag.Parse()

    fmt.Println("word:", *getPtr)
    // fmt.Println("numb:", *numbPtr)
    // fmt.Println("fork:", *boolPtr)
    // fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())


    if getPtr != nil {
        fmt.Println("word:", *getPtr)
    }

    

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


	var re = regexp.MustCompile(`AUTHTOKEN=(?P<token>\S+)`)
	var str = string(body)
	match := re.FindStringSubmatch(str)
	fmt.Printf("TOKEN: %s", match[1])

	f, err := os.Create("token.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(match[1])
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "saved a new token successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
