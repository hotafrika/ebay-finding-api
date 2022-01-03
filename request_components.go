package finding

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

// RequestBasic is used for requests without pages
type RequestBasic struct {
	URL    string        `json:"-" xml:"-"`
	Client *resty.Client `json:"-" xml:"-"`
}

/*
================================================================
*/

// RequestStandard represents basic part of ebay Finding request
type RequestStandard struct {
	PaginationInput ServicePaginationInput `json:"paginationInput,omitempty" xml:"paginationInput,omitempty"`
	SortOrder       string                 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	Affiliate       *ServiceAffiliate      `json:"affiliate,omitempty" xml:"affiliate,omitempty"`
	BuyerPostalCode string                 `json:"buyerPostalCode,omitempty" xml:"buyerPostalCode,omitempty"`

	RequestBasic
}

// ServiceAffiliate represents Affiliate Program
type ServiceAffiliate struct {
	NetworkID  string `json:"networkId,omitempty" xml:"networkId,omitempty"`
	TrackingID string `json:"trackingId,omitempty" xml:"trackingId,omitempty"`
	CustomID   string `json:"customId,omitempty" xml:"customId,omitempty"`
}

// WithAffiliate adds ServiceAffiliate
func (sr *RequestStandard) WithAffiliate(networkID, trackingID, customID string) *RequestStandard {
	sr.Affiliate = &ServiceAffiliate{
		NetworkID:  networkID,
		TrackingID: trackingID,
		CustomID:   customID,
	}
	return sr
}

// ServicePaginationInput represents PaginationInput
type ServicePaginationInput struct {
	EntriesPerPage int `json:"entriesPerPage,omitempty" xml:"entriesPerPage,omitempty"`
	PageNumber     int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

// WithPageLimit sets page limit to list of items
//  Min: 1. Max: 100. Default: 100.
func (sr *RequestStandard) WithPageLimit(limit int) *RequestStandard {
	if limit < 1 {
		limit = 1
	} else {
		if limit > 100 {
			limit = 100
		}
	}
	sr.PaginationInput.EntriesPerPage = limit
	return sr
}

// WithPageNumber sets page to get
//  Min: 1. Max: 100. Default: 100.
func (sr *RequestStandard) WithPageNumber(page int) *RequestStandard {
	if page < 1 {
		page = 1
	} else {
		if page > 100 {
			page = 100
		}
	}
	sr.PaginationInput.PageNumber = page
	return sr
}

// WithSortOrder sorts the returned items according to a single specified sort order.
// Default: BestMatch.
func (sr *RequestStandard) WithSortOrder(order SortOrderParameter) *RequestStandard {
	sr.SortOrder = string(order)
	return sr
}

// WithBuyerPostalCode adds BuyerPostalCode to requests
func (sr *RequestStandard) WithBuyerPostalCode(code string) *RequestStandard {
	sr.BuyerPostalCode = code
	return sr
}

/*
================================================================
*/

// RequestAspectFilter represents AspectFilter part of ebay Finding requests
type RequestAspectFilter struct {
	AspectFilter []ServiceAspectFilter `json:"aspectFilter,omitempty" xml:"aspectFilter,omitempty"` // removed
}

// ServiceAspectFilter represents AspectFilter
type ServiceAspectFilter struct {
	AspectName      string   `json:"aspectName" xml:"aspectName"`
	AspectValueName []string `json:"aspectValueName" xml:"aspectValueName"`
}

// WithAspectFilter adds AspectFilter to requests
func (sr *RequestAspectFilter) WithAspectFilter(aspectName string, aspectValues ...string) *RequestAspectFilter {
	if len(aspectValues) < 1 {
		return sr
	}
	if aspectName == "" {
		return sr
	}
	sr.AspectFilter = append(sr.AspectFilter, ServiceAspectFilter{aspectName, aspectValues})
	return sr
}

/*
================================================================
*/

// RequestCategories represents filter by categories in ebay Finding requests
type RequestCategories struct {
	CategoryID []string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
}

// WithCategoryIDInt adds category for searching (up to 3 categories)
//  4th and next categoryIDs will be skipped silently
func (sr *RequestCategories) WithCategoryIDInt(categoryID int) *RequestCategories {
	if len(sr.CategoryID) >= 3 {
		return sr
	}
	sr.CategoryID = append(sr.CategoryID, strconv.Itoa(categoryID))
	return sr
}

// WithCategoryID adds category for searching (up to 3 categories)
//  4th and next categoryIDs will be skipped silently
func (sr *RequestCategories) WithCategoryID(categoryID string) *RequestCategories {
	if len(sr.CategoryID) >= 3 {
		return sr
	}
	sr.CategoryID = append(sr.CategoryID, categoryID)
	return sr
}

// WithCategoriesID adds categories for searching (up to 3 categories)
//  4th and next categoryIDs will be skipped silently
func (sr *RequestCategories) WithCategoriesID(categoriesID ...string) *RequestCategories {
	for _, categoryID := range categoriesID {
		if len(sr.CategoryID) >= 3 {
			return sr
		}
		sr.CategoryID = append(sr.CategoryID, categoryID)
	}
	return sr
}

/*
================================================================
*/

// RequestDescriptionSearch represents filter by description search in ebay Finding requests
type RequestDescriptionSearch struct {
	DescriptionSearch bool `json:"descriptionSearch,omitempty" xml:"descriptionSearch,omitempty"`
}

// WithDescriptionSearch specifies whether your keyword query should be applied
// to item descriptions in addition to titles
func (sr *RequestDescriptionSearch) WithDescriptionSearch(ds bool) *RequestDescriptionSearch {
	sr.DescriptionSearch = ds
	return sr
}

/*
================================================================
*/

// RequestKeywords represents filter by keywords in ebay Finding requests
type RequestKeywords struct {
	Keywords string `json:"keywords,omitempty" xml:"keywords,omitempty"`
}

// WithKeywords adds keywords for searching
// Max length: 350. The maximum length for a single word is 98. Min length: 2.
// Key longer than 350 symbols will be trimmed. Key shorter 2 (0 < n < 2) symbols won't change Keywords field.
//  Empty string remove Keywords field.
func (sr *RequestKeywords) WithKeywords(key string) *RequestKeywords {
	if len(key) > 350 {
		key = key[:350]
	} else {
		if len(key) == 1 {
			return sr
		}
	}
	sr.Keywords = key
	return sr
}

/*
================================================================
*/

// RequestOutputSelector represents outputSelector part of ebay Finding requests
type RequestOutputSelector struct {
	OutputSelector []string `json:"outputSelector,omitempty" xml:"outputSelector,omitempty"`
}

// WithOutputSelectors adds OutputSelectors
func (sr *RequestOutputSelector) WithOutputSelectors(osps ...OutputSelectorParameter) *RequestOutputSelector {
	for _, osp := range osps {
		sr.OutputSelector = append(sr.OutputSelector, string(osp))
	}
	return sr
}

/*
================================================================
*/

// RequestProduct represents filter by products in ebay Finding requests
type RequestProduct struct {
	ProductID Product `json:"productId,omitempty" xml:"productId,omitempty"`
}

// Product ...
type Product struct {
	Type string `json:"type" xml:"type,attr"`
	Text string `json:"#text" xml:",cdata"`
}

// WithProductType adds productId to req
func (sr *RequestProduct) WithProductType(productType ProductTypeOption, id string) *RequestProduct {
	sr.ProductID.Type = string(productType)
	sr.ProductID.Text = id
	return sr
}
