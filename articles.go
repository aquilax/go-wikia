package gowikia

import (
	"encoding/json"
)

type AsSimpleJsonElementsItem struct {
	Text     string                     `json:"text"`
	Elements []AsSimpleJsonElementsItem `json:"elements"`
}

type AsSimpleJsonContentItem struct {
	Type     string                     `json:"type"`
	Text     string                     `json:"text"`
	Elements []AsSimpleJsonElementsItem `json:"elements"`
}

type AsSimpleJsonImagesItem struct {
	Src     string `json:"src"`
	Caption string `json:"caption"`
}

type AsSimpleJsonSections struct {
	Title   string                    `json:"title"`
	Level   int                       `json:"level"`
	Content []AsSimpleJsonContentItem `json:"content"`
	Images  []AsSimpleJsonImagesItem  `json:"images"`
}

type AsSimpleJsonResult struct {
	Sections AsSimpleJsonSections `json:"sections"`
}

// Get simplified article contents
// http://muppet.wikia.com/api/v1#!/Articles/getAsSimpleJson_get_0
func (wa *WikiaApi) AsSimpleJson(id int) (AsSimpleJsonResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "AsSimpleJson"},
		RequestParams{
			"id": intToStr(id),
		})
	if err != nil {
		return AsSimpleJsonResult{}, err
	}
	var ares AsSimpleJsonResult
	return ares, json.Unmarshal(jsonBlob, &ares)
}

type DetailsResultItemRevision struct {
	Id        int    `json:"id"`
	User      string `json:"user"`
	UserId    int    `json:"user_id"`
	Timestamp int    `json:"timestamp"`
}

type DetailsResultItem struct {
	Id                 int                       `json:"id"`
	Title              string                    `json:"title"`
	Ns                 int                       `json:"ns"`
	Url                string                    `json:"url"`
	Revision           DetailsResultItemRevision `json:"revision"`
	Comments           int                       `json:"comments"`
	Type               string                    `json:"type"`
	Abstract           string                    `json:"abstract"`
	Thumbnail          string                    `json:"thumbnail"`
	OriginalDimensions OriginalDimensions        `json:"original_dimensions"`
}

type DetailsResult struct {
	Items    []DetailsResultItem `json:"items"`
	Basepath string              `json:"basepath"`
}

type DetailsRequest struct {
	ids      []int
	titles   []string
	abstract int
	width    int
	height   int
}

func DetailsDefaults() DetailsRequest {
	return DetailsRequest{[]int{50}, []string{}, 100, 200, 200}
}

// Get details about one or more articles
// http://muppet.wikia.com/api/v1#!/Articles/getDetails_get_1
func (wa *WikiaApi) Details(dr DetailsRequest) (DetailsResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "Details"},
		RequestParams{
			"ids":      intArrToStr(dr.ids),
			"titles":   strArrToStr(dr.titles),
			"abstract": intToStr(dr.abstract),
			"width":    intToStr(dr.width),
			"height":   intToStr(dr.height),
		})
	if err != nil {
		return DetailsResult{}, err
	}
	var dres DetailsResult
	return dres, json.Unmarshal(jsonBlob, &dres)
}

type ListResultItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Ns    int    `json:"ns"`
}

type ListResult struct {
	Items    []ListResultItem `json:"items"`
	Offset   string           `json:"offset"`
	Basepath string           `json:"basepath"`
}

type ListRequest struct {
	category   string
	namespaces []int
	limit      int
	offset     string
}

// Get a list of pages on the current wiki
// http://muppet.wikia.com/api/v1#!/Articles/getList_get_3
func (wa *WikiaApi) ArticlesList(lr ListRequest) (ListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "List"},
		RequestParams{
			"category":   lr.category,
			"namespaces": intArrToStr(lr.namespaces),
			"limit":      intToStr(lr.limit),
			"offset":     lr.offset,
		})
	if err != nil {
		return ListResult{}, err
	}
	var lres ListResult
	return lres, json.Unmarshal(jsonBlob, &lres)
}

type MostLinkedResultItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Ns          int    `json:"ns"`
	BackLinkCnt int    `json:"backlink_cnt"`
}

type MostLinkedResult struct {
	Items    []MostLinkedResultItem `json:"items"`
	Basepath string                 `json:"basepath"`
}

// Get the most linked articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getTop_get_4
func (wa *WikiaApi) ArticlesMostLinked() (MostLinkedResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "MostLinked"},
		RequestParams{})
	if err != nil {
		return MostLinkedResult{}, err
	}
	var mlres MostLinkedResult
	return mlres, json.Unmarshal(jsonBlob, &mlres)
}

type Creator struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type ArticlesNewResult struct {
	Id                 int                `json:"id"`
	Ns                 int                `json:"ns"`
	Title              string             `json:"title"`
	Abstract           string             `json:"abstract"`
	Quality            int                `json:"quality"`
	Url                string             `json:"url"`
	Creator            Creator            `json:"creator"`
	CreationDate       string             `json:"creation_date"`
	Thumbnail          string             `json:"thumbnail"`
	OriginalDimensions OriginalDimensions `json:"original_dimensions"`
}

type ArticlesNewRequest struct {
	namespaces        []int
	limit             int
	minArticleQuality int
}

func (wa *WikiaApi) ArticlesNew(nr ArticlesNewRequest) (ArticlesNewResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "MostLinked"},
		RequestParams{})
	if err != nil {
		return ArticlesNewResult{}, err
	}
	var nres ArticlesNewResult
	return nres, json.Unmarshal(jsonBlob, &nres)
}
