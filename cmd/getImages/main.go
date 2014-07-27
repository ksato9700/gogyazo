//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//
package main

import (
	"flag"
	"fmt"
	"github.com/ksato9700/gogyazo"
	"log"
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
		log.Fatalf("Get access token first")
	}

	images, err := client.Images()
	if err != nil {
		log.Fatal(err)
	}
	for _, image := range images {
		fmt.Printf("%+v\n", image)
	}
}
