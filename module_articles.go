package gowikia

import (
	"encoding/json"
)

type ArticlesAsSimpleJsonElementsItem struct {
	Text     string                             `json:"text"`
	Elements []ArticlesAsSimpleJsonElementsItem `json:"elements"`
}

type ArticlesAsSimpleJsonContentItem struct {
	Type     string                             `json:"type"`
	Text     string                             `json:"text"`
	Elements []ArticlesAsSimpleJsonElementsItem `json:"elements"`
}

type ArticlesAsSimpleJsonImagesItem struct {
	Src     string `json:"src"`
	Caption string `json:"caption"`
}

type ArticlesAsSimpleJsonSection struct {
	Title   string                            `json:"title"`
	Level   int                               `json:"level"`
	Content []ArticlesAsSimpleJsonContentItem `json:"content"`
	Images  []ArticlesAsSimpleJsonImagesItem  `json:"images"`
}

type ArticlesAsSimpleJsonResult struct {
	Sections []ArticlesAsSimpleJsonSection `json:"sections"`
}

// Get simplified article contents
// http://muppet.wikia.com/api/v1#!/Articles/getAsSimpleJson_get_0
func (wa *WikiaApi) ArticlesAsSimpleJson(id int) (*ArticlesAsSimpleJsonResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_AS_SIMPLE_JSON_SEGMENT},
		RequestParams{
			"id": intToStr(id),
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesAsSimpleJsonResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesDetailsResultItemRevision struct {
	Id        int    `json:"id"`
	User      string `json:"user"`
	UserId    int    `json:"user_id"`
	Timestamp string `json:"timestamp"` //WTF: TS as string
}

type ArticlesDetailsResultItem struct {
	Id                 int                               `json:"id"`
	Title              string                            `json:"title"`
	Ns                 int                               `json:"ns"`
	Url                string                            `json:"url"`
	Revision           ArticlesDetailsResultItemRevision `json:"revision"`
	Comments           int                               `json:"comments"`
	Type               string                            `json:"type"`
	Abstract           string                            `json:"abstract"`
	Thumbnail          string                            `json:"thumbnail"`
	OriginalDimensions OriginalDimensions                `json:"original_dimensions"`
	//Metadata	// WTF: Not documented
}

type ArticlesDetailsResult struct {
	Items    map[string]ArticlesDetailsResultItem `json:"items"`
	Basepath string                               `json:"basepath"`
}

type ArticlesDetailsRequest struct {
	ids      []int
	titles   []string
	abstract int
	width    int
	height   int
}

func DefaultArticlesDetailsRequest() ArticlesDetailsRequest {
	return ArticlesDetailsRequest{[]int{50}, []string{}, 100, 200, 200}
}

// Get details about one or more articles
// http://muppet.wikia.com/api/v1#!/Articles/getDetails_get_1
func (wa *WikiaApi) ArticlesDetails(params ArticlesDetailsRequest) (*ArticlesDetailsResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_DETAILS_SEGMENT},
		RequestParams{
			"ids":      intArrToStr(params.ids),
			"titles":   strArrToStr(params.titles),
			"abstract": intToStr(params.abstract),
			"width":    intToStr(params.width),
			"height":   intToStr(params.height),
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesDetailsResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ListResultItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Ns    int    `json:"ns"`
}

type ArticlesListResult struct {
	Items    []ListResultItem `json:"items"`
	Offset   string           `json:"offset"`
	Basepath string           `json:"basepath"`
}

type ArticlesListRequest struct {
	category   string
	namespaces []int
	limit      int
	offset     string
}

func DefaultArticlesListRequest() ArticlesListRequest {
	return ArticlesListRequest{"", []int{0}, 25, ""}
}

// Get a list of pages on the current wiki
// http://muppet.wikia.com/api/v1#!/Articles/getList_get_3
func (wa *WikiaApi) ArticlesList(params ArticlesListRequest) (*ArticlesListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_LIST_SEGMENT},
		RequestParams{
			"category":   params.category,
			"namespaces": intArrToStr(params.namespaces),
			"limit":      intToStr(params.limit),
			"offset":     params.offset,
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesListResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type MostLinkedResultItem struct {
	Id          string `json:"id"` // WTF: String
	Title       string `json:"title"`
	Url         string `json:"url"`
	Ns          int    `json:"ns"`
	BackLinkCnt int    `json:"backlink_cnt"` // WTF: missing from the result
}

type ArticlesMostLinkedResult struct {
	Items    []MostLinkedResultItem `json:"items"`
	Basepath string                 `json:"basepath"`
}

// Get the most linked articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getTop_get_4
func (wa *WikiaApi) ArticlesMostLinked() (*ArticlesMostLinkedResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_MOST_LINKED_SEGMENT},
		RequestParams{})
	if err != nil {
		return nil, err
	}
	var result *ArticlesMostLinkedResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type Creator struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type ArticlesNewResultItem struct {
	Id                 int                `json:"id"`
	Ns                 int                `json:"ns"`
	Title              string             `json:"title"`
	Abstract           string             `json:"abstract"`
	Quality            int                `json:"quality"`
	Url                string             `json:"url"`
	Creator            Creator            `json:"creator"`
	CreationDate       string             `json:"creation_date"` // WTF: String
	Thumbnail          string             `json:"thumbnail"`
	OriginalDimensions OriginalDimensions `json:"original_dimensions"`
}

// WTF: Wrong model schema
type ArticlesNewResult struct {
	Items []ArticlesNewResultItem `json:"items"`
}

type ArticlesNewRequest struct {
	Namespaces        []int
	Limit             int
	MinArticleQuality int
}

func DefaultArticlesNewRequest() ArticlesNewRequest {
	return ArticlesNewRequest{[]int{0}, 20, 10}
}

// Get list of new articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getNew_get_6
func (wa *WikiaApi) ArticlesNew(params ArticlesNewRequest) (*ArticlesNewResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_NEW_SEGMENT},
		RequestParams{
			"namespaces":        intArrToStr(params.Namespaces),
			"limit":             intToStr(params.Limit),
			"minArticleQuality": intToStr(params.MinArticleQuality),
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesNewResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesPopularItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ArticlesPopularResult struct {
	Items    []ArticlesPopularItem `json:"items"`
	Basepath string                `json:"basepath"`
}

// Get popular articles for the current wiki (from the beginning of time)
// http://muppet.wikia.com/api/v1#!/Articles/getPopular_get_7
func (wa *WikiaApi) ArticlesPopular(limit int) (*ArticlesPopularResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_POPULAR_SEGMENT},
		RequestParams{
			"limit": intToStr(limit),
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesPopularResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesTopItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Ns    int    `json:"ns"`
}

type ArticlesTopResult struct {
	Items    []ArticlesTopItem `json:"items"`
	Basepath string            `json:"basepath"`
}

type ArticlesTopRequest struct {
	Namespaces []int
	Category   string
	Limit      int // WTF: String in original doc
}

func DefaultArticlesTopRequest() ArticlesTopRequest {
	return ArticlesTopRequest{[]int{0}, "", 10}
}

// Get the most viewed articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getTop_get_9
func (wa *WikiaApi) ArticlesTop(params ArticlesTopRequest) (*ArticlesTopResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_TOP_SEGMENT},
		RequestParams{
			"namespaces": intArrToStr(params.Namespaces),
			"category":   params.Category,
			"limit":      intToStr(params.Limit),
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesTopResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesTopByHubWiki struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Domain   string `json:"domain"`
}

type ArticlesTopByHubArticle struct {
	Id int `json:"id"`
	Ns int `json:"ns"`
}

type ArticlesTopByHubItem struct {
	Wiki     ArticlesTopByHubWiki      `json:"wiki"`
	Articles []ArticlesTopByHubArticle `json:"articles"`
}

// WTF: Does not work BadRequest
type ArticlesTopByHubResult struct {
	Items []ArticlesTopByHubItem `json:"items"`
}

type ArticlesTopByHubRequest struct {
	Hub        string
	Lang       string
	Namespaces []int
}

func DefaultArticlesTopByHubRequest() ArticlesTopByHubRequest {
	return ArticlesTopByHubRequest{"gaming", "en", []int{0}}
}

// Get the top articles by pageviews for a hub
// http://muppet.wikia.com/api/v1#!/Articles/getTopByHub_get_11
func (wa *WikiaApi) ArticlesTopByHub(params ArticlesTopByHubRequest) (*ArticlesTopByHubResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, ARTICLES_TOP_BY_HUB_SEGMENT},
		RequestParams{
			"hub":        params.Hub,
			"lang":       params.Lang,
			"namespaces": intArrToStr(params.Namespaces),
		})
	if err != nil {
		return nil, err
	}
	var result *ArticlesTopByHubResult
	return result, json.Unmarshal(jsonBlob, &result)
}
