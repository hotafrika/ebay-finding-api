package finding

import (
	"encoding/xml"
	"fmt"
)

// AdvancedRequest searches for items on eBay by category or keyword or both.
type AdvancedRequest struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services findItemsAdvancedRequest" json:"-"`
	RequestCategories
	RequestKeywords
	RequestDescriptionSearch
	RequestAspectFilter
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes AdvancedRequest for page #
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
	if err != nil {
		return AdvancedResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// Execute executes AdvancedRequest for the first page
func (sr *AdvancedRequest) Execute() (AdvancedResponse, error) {
	return sr.GetPage(1)
}

// GetBody return AdvancedRequest body as XML
func (sr *AdvancedRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *AdvancedRequest) getBody() ([]byte, error) {
	sr.prepare()
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// ByCategoryRequest searches for items on eBay using specific eBay category ID numbers
// (input category ID numbers using categoryId).
type ByCategoryRequest struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services findItemsByCategoryRequest" json:"-"`
	RequestCategories
	RequestAspectFilter
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes ByCategoryRequest for page #
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

// Execute executes ByCategoryRequest for the first page
func (sr *ByCategoryRequest) Execute() (ByCategoryResponse, error) {
	return sr.GetPage(1)
}

// GetBody return ByCategoryRequest body as XML
func (sr *ByCategoryRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *ByCategoryRequest) getBody() ([]byte, error) {
	sr.prepare()
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// ByKeywordsRequest searches for items on eBay by a keyword query (keywords).
type ByKeywordsRequest struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services findItemsByKeywordsRequest" json:"-"`
	RequestKeywords
	RequestDescriptionSearch
	RequestAspectFilter
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes ByKeywordsRequest for page #
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

// Execute executes ByKeywordsRequest for the first page
func (sr *ByKeywordsRequest) Execute() (ByKeywordsResponse, error) {
	return sr.GetPage(1)
}

// GetBody return ByKeywordsRequest body as XML
func (sr *ByKeywordsRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *ByKeywordsRequest) getBody() ([]byte, error) {
	sr.prepare()
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// ByProductRequest searches for items on eBay using specific eBay product values.
type ByProductRequest struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services findItemsByProductRequest" json:"-"`
	RequestProduct
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// GetPage executes ByProductRequest for page #
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

// Execute executes ByProductRequest for the first page
func (sr *ByProductRequest) Execute() (ByProductResponse, error) {
	return sr.GetPage(1)
}

// GetBody return ByProductRequest body as XML
func (sr *ByProductRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *ByProductRequest) getBody() ([]byte, error) {
	sr.prepare()
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// InEbayStoresRequest searches for items in the eBay store inventories.
// You can combine storeName with keywords to find specific items,
// or use keywords without storeName to search for items in all eBay stores.
type InEbayStoresRequest struct {
	XMLName   xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services findItemsIneBayStoresRequest" json:"-"`
	StoreName string   `json:"storeName,omitempty" xml:"storeName,omitempty"`
	RequestAspectFilter
	RequestCategories
	RequestKeywords
	RequestItemFilter
	RequestOutputSelector
	RequestStandard
}

// WithStoreName adds store's name to InEbayStoresRequest
func (sr *InEbayStoresRequest) WithStoreName(name string) *InEbayStoresRequest {
	sr.StoreName = name
	return sr
}

// GetPage executes InEbayStoresRequest for page #
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

// Execute executes InEbayStoresRequest for the first page
func (sr *InEbayStoresRequest) Execute() (InEbayStoresResponse, error) {
	return sr.GetPage(1)
}

// GetBody return InEbayStoresRequest body as XML
func (sr *InEbayStoresRequest) GetBody() ([]byte, error) {
	sr.prepare()
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *InEbayStoresRequest) getBody() ([]byte, error) {
	sr.prepare()
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// GetHistogramsRequest retrieves category and/or aspect
// histogram information for the eBay category you specify using the categoryId field.
type GetHistogramsRequest struct {
	XMLName    xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services getHistogramsRequest" json:"-"`
	CategoryId string   `json:"categoryId" xml:"categoryId"`
	RequestBasic
}

// WithCategory adds categoryId to GetHistogramsRequest
func (sr *GetHistogramsRequest) WithCategory(category string) *GetHistogramsRequest {
	sr.CategoryId = category
	return sr
}

// Execute executes GetHistogramsRequest
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

// GetBody return GetHistogramsRequest body as XML
func (sr *GetHistogramsRequest) GetBody() ([]byte, error) {
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *GetHistogramsRequest) getBody() ([]byte, error) {
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// GetKeywordsRecommendationRequest retrieves commonly used words found in eBay titles,
// based on the words you supply to the call.
type GetKeywordsRecommendationRequest struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services getSearchKeywordsRecommendationRequest" json:"-"`
	RequestKeywords
	RequestBasic
}

// Execute executes GetKeywordsRecommendationRequest
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

// GetBody return GetKeywordsRecommendationRequest body as XML
func (sr *GetKeywordsRecommendationRequest) GetBody() ([]byte, error) {
	return xml.MarshalIndent(sr, "", "  ")
}

func (sr *GetKeywordsRecommendationRequest) getBody() ([]byte, error) {
	return xml.Marshal(sr)
}

/*
==============================================================================
*/

// GetVersionRequest returns the current version of the service.
// This simple call can be used to monitor the service for availability.
// This call has no input parameters and the response contains only the standard output fields.
type GetVersionRequest struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/search/v1/services getVersionRequest" json:"-"`
	RequestBasic
}

// Execute executes GetVersionRequest
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

// GetBody return GetVersionRequest body as XML
func (sr *GetVersionRequest) GetBody() ([]byte, error) {
	return xml.MarshalIndent(sr, "", "  ")
}
