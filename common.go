package gowikia

import (
	"io/ioutil"
	"net/http"
)

const (
	METHOD_GET = "GET"
)

type RequestParams map[string]string

func getJsonBlob(wikiaUrl string, segments []string, params RequestParams) ([]byte, error) {
	path := generatePath(segments)
	query := generateQuery(params)
	url := generateApiUrl(wikiaUrl, path, query)
	return fetchUrl(url)
}

func fetchUrl(u string) ([]byte, error) {
	response, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}