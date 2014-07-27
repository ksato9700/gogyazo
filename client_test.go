//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//
package gogyazo

import (
	"os"
	"testing"
)

var cid = os.Getenv("OAUTH2_GYAZO_CLIENT_ID")
var csec = os.Getenv("OAUTH2_GYAZO_CLIENT_SECRET")
var redurl = os.Getenv("OAUTH2_GYAZO_REDIRECT_URI")

func TestClient(t *testing.T) {
	client := NewClient(cid, csec, redurl)
	if client.HttpClient == nil {
		t.Error("client.HttpClient is nil")
		return
	}
	images, err := client.Images()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", images)
}
