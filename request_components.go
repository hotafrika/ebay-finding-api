package finding

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

// RequestBasic is used for requests without pages
type RequestBasic struct {
	URL    string        `json:"-"`
	Client *resty.Client `json:"-"`
}

/*
================================================================
*/

type RequestStandard struct {
	PaginationInput ServicePaginationInput `json:"paginationInput,omitempty"`
	SortOrder       string                 `json:"sortOrder,omitempty"`
	//Affiliate				ServiceAffiliate		`json:"affiliate,omitempty"`		// to implement
	//BuyerPostalCode 		string 					`json:"buyerPostalCode,omitempty"`	// to implement

	RequestBasic
}

// ServicePaginationInput represents PaginationInput
type ServicePaginationInput struct {
	EntriesPerPage int `json:"entriesPerPage,omitempty"`
	PageNumber     int `json:"pageNumber,omitempty"`
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

/*
================================================================
*/

type RequestAspectFilter struct {
	////AspectFilter          []ServiceAspectFilter   `json:"aspectFilter,omitempty"` 	// removed
}

// ServiceAspectFilter represents AspectFilter
//  It is NOT implemented yet
type ServiceAspectFilter struct {
	AspectName      string   `json:"aspectName"`
	AspectValueName []string `json:"aspectValueName"`
}

/*
================================================================
*/

type RequestCategories struct {
	CategoryID []string `json:"categoryId,omitempty"`
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

type RequestDescriptionSearch struct {
	DescriptionSearch bool `json:"descriptionSearch,omitempty"`
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

type RequestKeywords struct {
	Keywords string `json:"keywords,omitempty"`
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

type RequestOutputSelector struct {
	OutputSelector []string `json:"outputSelector,omitempty"`
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

type RequestProduct struct {
	ProductID Product `json:"productId"`
}

type Product struct {
	Type string `json:"type"`
	Text string `json:"#text"`
}

// WithProductType adds productId to req
func (sr *RequestProduct) WithProductType(productType ProductTypeOption, id string) *RequestProduct {
	sr.ProductID.Type = string(productType)
	sr.ProductID.Text = id
	return sr
}
