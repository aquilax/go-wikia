package gowikia

import "net/url"

func generateApiUrl(wikiaUrl, path, query string) (apiURL string, err error) {
	parsedUrl, err := url.Parse(wikiaUrl)
	if err != nil {
		return "", err
	}
	parsedUrl.Path = path
	parsedUrl.RawQuery = query
	return parsedUrl.String(), nil
}
