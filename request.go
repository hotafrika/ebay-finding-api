package finding

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type AdvancedRequest struct {
	RequestCategories
	RequestKeywords
	RequestDescriptionSearch
	RequestAspectFilter
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes req for page #
// Valid pages # 1 - 100
func (sr *AdvancedRequest) GetPage(page int) (AdvancedResponse, error) {
	if page < 1 {
		page = 1
	}
	sr.WithPageNumber(page)
	body, err := sr.getBody()
	if err != nil {
		return AdvancedResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return AdvancedResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return AdvancedResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := AdvancedResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return AdvancedResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// Execute executes req for the first page
func (sr *AdvancedRequest) Execute() (AdvancedResponse, error) {
	return sr.GetPage(1)
}

// GetBody return req body as JSON
func (sr *AdvancedRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *AdvancedRequest) getBody() ([]byte, error) {
	sr.prepare()
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type ByCategoryRequest struct {
	RequestCategories
	RequestAspectFilter
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes req for page #
// Valid pages # 1 - 100
func (sr *ByCategoryRequest) GetPage(page int) (ByCategoryResponse, error) {
	if page < 1 {
		page = 1
	}
	sr.WithPageNumber(page)
	body, err := sr.getBody()
	if err != nil {
		return ByCategoryResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return ByCategoryResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return ByCategoryResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	bcr := ByCategoryResponse{}
	err = xml.Unmarshal(res.Body(), &bcr)
	if err != nil {
		return ByCategoryResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return bcr, nil
}

// Execute executes req for the first page
func (sr *ByCategoryRequest) Execute() (ByCategoryResponse, error) {
	return sr.GetPage(1)
}

// GetBody return req body as JSON
func (sr *ByCategoryRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *ByCategoryRequest) getBody() ([]byte, error) {
	sr.prepare()
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type ByKeywordsRequest struct {
	RequestKeywords
	RequestDescriptionSearch
	RequestAspectFilter
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes req for page #
// Valid pages # 1 - 100
func (sr *ByKeywordsRequest) GetPage(page int) (ByKeywordsResponse, error) {
	if page < 1 {
		page = 1
	}
	sr.WithPageNumber(page)
	body, err := sr.getBody()
	if err != nil {
		return ByKeywordsResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return ByKeywordsResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return ByKeywordsResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	bkr := ByKeywordsResponse{}
	err = xml.Unmarshal(res.Body(), &bkr)
	if err != nil {
		return ByKeywordsResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return bkr, nil
}

// Execute executes req for the first page
func (sr *ByKeywordsRequest) Execute() (ByKeywordsResponse, error) {
	return sr.GetPage(1)
}

// GetBody return req body as JSON
func (sr *ByKeywordsRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *ByKeywordsRequest) getBody() ([]byte, error) {
	sr.prepare()
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type ByProductRequest struct {
	RequestProduct
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes req for page #
// Valid pages # 1 - 100
func (sr *ByProductRequest) GetPage(page int) (ByProductResponse, error) {
	if page < 1 {
		page = 1
	}
	sr.WithPageNumber(page)
	body, err := sr.getBody()
	if err != nil {
		return ByProductResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return ByProductResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return ByProductResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	bpr := ByProductResponse{}
	err = xml.Unmarshal(res.Body(), &bpr)
	if err != nil {
		return ByProductResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return bpr, nil
}

// Execute executes req for the first page
func (sr *ByProductRequest) Execute() (ByProductResponse, error) {
	return sr.GetPage(1)
}

// GetBody return req body as JSON
func (sr *ByProductRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *ByProductRequest) getBody() ([]byte, error) {
	sr.prepare()
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type InEbayStoresRequest struct {
	StoreName string `json:"storeName,omitempty"`
	RequestAspectFilter
	RequestCategories
	RequestKeywords
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// WithStoreName adds store's name to req
func (sr *InEbayStoresRequest) WithStoreName(name string) *InEbayStoresRequest {
	sr.StoreName = name
	return sr
}

// GetPage executes req for page #
// Valid pages # 1 - 100
func (sr *InEbayStoresRequest) GetPage(page int) (InEbayStoresResponse, error) {
	if page < 1 {
		page = 1
	}
	sr.WithPageNumber(page)
	body, err := sr.getBody()
	if err != nil {
		return InEbayStoresResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return InEbayStoresResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return InEbayStoresResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	iesr := InEbayStoresResponse{}
	err = xml.Unmarshal(res.Body(), &iesr)
	if err != nil {
		return InEbayStoresResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return iesr, nil
}

// Execute executes req for the first page
func (sr *InEbayStoresRequest) Execute() (InEbayStoresResponse, error) {
	return sr.GetPage(1)
}

// GetBody return req body as JSON
func (sr *InEbayStoresRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *InEbayStoresRequest) getBody() ([]byte, error) {
	sr.prepare()
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type GetHistogramsRequest struct {
	CategoryId string `json:"categoryId"`
	RequestBasic
}

// WithCategory adds categoryId to req
func (sr *GetHistogramsRequest) WithCategory(category string) *GetHistogramsRequest {
	sr.CategoryId = category
	return sr
}

// Execute executes req
func (sr *GetHistogramsRequest) Execute() (GetHistogramsResponse, error) {
	body, err := sr.getBody()
	if err != nil {
		return GetHistogramsResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return GetHistogramsResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetHistogramsResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ghr := GetHistogramsResponse{}
	err = xml.Unmarshal(res.Body(), &ghr)
	if err != nil {
		return GetHistogramsResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ghr, nil
}

// GetBody return req body as JSON
func (sr *GetHistogramsRequest) GetBody() ([]byte, error) {
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *GetHistogramsRequest) getBody() ([]byte, error) {
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type GetKeywordsRecommendationRequest struct {
	RequestKeywords
	RequestBasic
}

// Execute executes req
func (sr *GetKeywordsRecommendationRequest) Execute() (GetKeywordsRecommendationResponse, error) {
	body, err := sr.getBody()
	if err != nil {
		return GetKeywordsRecommendationResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	res, err := sr.Client.R().SetBody(body).Post(sr.URL)
	if err != nil {
		return GetKeywordsRecommendationResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetKeywordsRecommendationResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	gkrr := GetKeywordsRecommendationResponse{}
	err = xml.Unmarshal(res.Body(), &gkrr)
	if err != nil {
		return GetKeywordsRecommendationResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return gkrr, nil
}

// GetBody return req body as JSON
func (sr *GetKeywordsRecommendationRequest) GetBody() ([]byte, error) {
	return json.MarshalIndent(sr, "", "  ")
}

func (sr *GetKeywordsRecommendationRequest) getBody() ([]byte, error) {
	return json.Marshal(sr)
}

/*
==============================================================================
*/

type GetVersionRequest struct {
	RequestBasic
}

// Execute executes req
func (sr *GetVersionRequest) Execute() (GetVersionResponse, error) {
	res, err := sr.Client.R().Post(sr.URL)
	if err != nil {
		return GetVersionResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetVersionResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	gvr := GetVersionResponse{}
	err = xml.Unmarshal(res.Body(), &gvr)
	if err != nil {
		return GetVersionResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return gvr, nil
}
