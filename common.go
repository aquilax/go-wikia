package gowikia

import (
	"io/ioutil"
	"net/http"
)

const (
	METHOD_GET = "GET"
)

func (wa *WikiaApi) fetchUrl(u string) ([]byte, error) {
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
