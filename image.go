//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//
package gogyazo

type Images []Image

type Image struct {
	Id           string `json:"image_id"`
	ParmalinkUrl string `json:"parmalink_url"`
	Url          string `json:"url"`
	Type         string `json:"type"`
	ThumbUrl     string `json:"thumb_url"`
	CreatedAt    string `json:"created_at"`
}
