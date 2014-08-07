/*
Package gowikia provides a way to accessing Wikia.com's API from go

Example usage:

  package main

  import "github.com/aquilax/go-wikia"

  func main() {
      wikiaApi = NewWikiaApi("http://muppet.wikia.com")
      latestActivity, err := wikiaApi.LatestActivity(ActivityDefaults())
	  ...
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
