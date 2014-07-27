//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//
package main

import (
	"flag"
	"fmt"
	"github.com/ksato9700/gogyazo"
	"os"
)

var (
	code = flag.String("code", "", "authcode to get access token")
)

func main() {
	flag.Parse()

	client := gogyazo.NewClient(
		os.Getenv("OAUTH2_GYAZO_CLIENT_ID"),
		os.Getenv("OAUTH2_GYAZO_CLIENT_SECRET"),
		os.Getenv("OAUTH2_GYAZO_REDIRECT_URI"),
	)
	if client.HttpClient == nil {
		if *code == "" {
			fmt.Println(client.AuthCodeUrl)
			return
		} else {
			client.Exchange(*code)
		}
	}
	fmt.Printf("%+v\n", client.Transport.Token.AccessToken)
}
