/*
Package gowikia provides a way to accessing Wikia.com's API from go

Example usage:

package main

import (
	"fmt"
	"github.com/aquilax/go-wikia"
)

func main() {
	wa, err := gowikia.NewWikiaApi("http://muppet.wikia.com/")
	if err != nil {
		fmt.Printf("Error creating the API: %s", err)
	}
	latestActivity, err := wa.ActivityLatestActivity(gowikia.DefaultActivityRequest())
	if err != nil {
		fmt.Printf("Error getting result: %s", err)
	}
	....
}
*/

package gowikia

type WikiaApi struct {
	url string
}

// NewWikiaApi function creates new wikiaApi or throws an error on invalid url
func NewWikiaApi(wikiaUrl string) (*WikiaApi, error) {
	valid, err := isValidUrl(wikiaUrl)
	if !valid {
		return nil, err
	}
	return &WikiaApi{wikiaUrl}, nil
}
