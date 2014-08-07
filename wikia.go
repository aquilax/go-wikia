package gowikia

const (
	API_URL = "https://api.tictail.com"
	API_VER = "v1"

	API_GET = "GET"

	API_STORES = "stores"
)

type WikiaApi struct {
	url string
}

func NewWikiaApi(wikiaUrl string) (*WikiaApi, error) {
	valid, err := isValidUrl(wikiaUrl)
	if !valid {
		return nil, err
	}
	return &WikiaApi{wikiaUrl}, nil
}
