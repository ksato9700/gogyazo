//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//

// Package gogyazo proivdes structs and functions for accessing Gyazo API.
package gogyazo

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BaseUrl       = "https://api.gyazo.com"
	UploadUrlBase = "https://upload.gyazo.com"
	ApiBaseUrl    = BaseUrl + "/api"
	AuthUrl       = "https://gyazo.com/oauth/authorize"
	TokenUrl      = "https://gyazo.com/oauth/token"
	CacheFile     = "cache.json"
)

type Client struct {
	HttpClient  *http.Client
	Transport   *oauth.Transport
	AuthCodeUrl string
}

func (c *Client) setToken(token *oauth.Token) {
	c.Transport.Token = token
	c.HttpClient = c.Transport.Client()
}

// Create a new instance of Client
func NewClient(cid string, csec string, redurl string) *Client {
	config := &oauth.Config{
		ClientId:     cid,
		ClientSecret: csec,
		Scope:        "public",
		AuthURL:      AuthUrl,
		TokenURL:     TokenUrl,
		RedirectURL:  redurl,
		TokenCache:   oauth.CacheFile(CacheFile),
	}
	// fmt.Printf("%+v\n", config)

	transport := &oauth.Transport{Config: config}
	token, err := config.TokenCache.Token()
	if err != nil {
		return &Client{
			Transport:   transport,
			AuthCodeUrl: config.AuthCodeURL(""),
		}
	} else {
		c := Client{Transport: transport}
		c.setToken(token)
		return &c
	}
}

//
func (c *Client) Exchange(code string) error {
	token, err := c.Transport.Exchange(code)
	if err != nil {
		return err
	}
	c.setToken(token)
	return nil
}

//
func (c *Client) apiGet(path string, params url.Values, out interface{}) error {
	urlstring := ApiBaseUrl + path
	if params != nil {
		urlstring += "?" + params.Encode()
	}
	// fmt.Println("urlstring", urlstring)
	res, err := c.HttpClient.Get(urlstring)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if strout, ok := out.(*string); ok {
		*strout = string(data)
	} else {
		err = json.Unmarshal(data, out)
	}
	return err
}

//
func (c *Client) Get(resource string, out interface{}) error {
	return c.apiGet(resource, nil, out)
}

//
func (c *Client) Images() (Images, error) {
	var images Images
	err := c.Get("/images", &images)
	return images, err
}
