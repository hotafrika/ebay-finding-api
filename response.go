package finding

import "encoding/xml"

// AdvancedResponse represents findItemsAdvancedResponse
type AdvancedResponse struct {
	XMLName xml.Name `xml:"findItemsAdvancedResponse"`
	ResponseAspectHistogramContainer
	ResponseCategoryHistogramContainer
	ResponseConditionHistogramContainer
	responseStandard
	// ItemSearchURL is a URL to view the search results on the eBay web site.
	// The search results on the web site will use the same pagination as the API search results.
	ItemSearchURL string `xml:"itemSearchURL"`
	ResponsePaginationOutput
	ResponseSearchResult
}

// ByCategoryResponse represents findItemsByCategoryResponse
type ByCategoryResponse struct {
	XMLName xml.Name `xml:"findItemsByCategoryResponse"`
	ResponseAspectHistogramContainer
	ResponseCategoryHistogramContainer
	ResponseConditionHistogramContainer
	responseStandard
	// ItemSearchURL is a URL to view the search results on the eBay web site.
	// The search results on the web site will use the same pagination as the API search results.
	ItemSearchURL string `xml:"itemSearchURL"`
	ResponsePaginationOutput
	ResponseSearchResult
}

// ByKeywordsResponse represents findItemsByKeywordsResponse
type ByKeywordsResponse struct {
	XMLName xml.Name `xml:"findItemsByKeywordsResponse"`
	ResponseAspectHistogramContainer
	ResponseCategoryHistogramContainer
	ResponseConditionHistogramContainer
	responseStandard
	// ItemSearchURL is a URL to view the search results on the eBay web site.
	// The search results on the web site will use the same pagination as the API search results.
	ItemSearchURL string `xml:"itemSearchURL"`
	PaginationOutput
	ResponseSearchResult
}

// ByProductResponse represents findItemsByProductResponse
type ByProductResponse struct {
	XMLName xml.Name `xml:"findItemsByProductResponse"`
	ResponseAspectHistogramContainer
	ResponseConditionHistogramContainer
	responseStandard
	// ItemSearchURL is a URL to view the search results on the eBay web site.
	// The search results on the web site will use the same pagination as the API search results.
	ItemSearchURL string `xml:"itemSearchURL"`
	ResponsePaginationOutput
	ResponseSearchResult
}

// InEbayStoresResponse represents findItemsIneBayStoresResponse
type InEbayStoresResponse struct {
	XMLName xml.Name `xml:"findItemsIneBayStoresResponse"`
	ResponseAspectHistogramContainer
	ResponseCategoryHistogramContainer
	ResponseConditionHistogramContainer
	responseStandard
	// ItemSearchURL is a URL to view the search results on the eBay web site.
	// The search results on the web site will use the same pagination as the API search results.
	ItemSearchURL string `xml:"itemSearchURL"`
	ResponsePaginationOutput
	ResponseSearchResult
}

// GetHistogramsResponse represents getHistogramsResponse
type GetHistogramsResponse struct {
	XMLName xml.Name `xml:"getHistogramsResponse"`
	responseStandard
}

// GetKeywordsRecommendationResponse represents getSearchKeywordsRecommendationResponse
type GetKeywordsRecommendationResponse struct {
	XMLName  xml.Name `xml:"getSearchKeywordsRecommendationResponse"`
	Keywords string   `xml:"keywords"`
	responseStandard
}

// GetVersionResponse represents getVersionResponse
type GetVersionResponse struct {
	XMLName xml.Name `xml:"getVersionResponse"`
	responseStandard
}
