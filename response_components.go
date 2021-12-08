package finding

type responseStandard struct {
	// Ack indicates whether the error is a fatal error (causing the request to fail) or a less severe
	// error (a warning) that should be communicated to the user.
	Ack string `xml:"ack"`
	// ErrorMessage Information regarding an error or warning that occurred when eBay processed the request.
	// Not returned when the ack value is Success. Run-time errors are not reported here,
	// but are instead reported as part of a SOAP fault.
	ErrorMessage []Error `xml:"errorMessage>error"`
	// Timestamp  represents the date and time when eBay processed the request.
	Timestamp string `xml:"timestamp"`
	// Version is the release version that eBay used to process the request. Developer Technical Support
	// may ask you for the version value if you work with them to troubleshoot issues.
	Version string `xml:"version"`
}

type Error struct {
	Category    string      `xml:"category"`
	Domain      string      `xml:"domain"`
	ErrorID     string      `xml:"errorId"`
	ExceptionID string      `xml:"exceptionId"`
	Message     string      `xml:"message"`
	Parameters  []Parameter `xml:"parameter"`
	Severity    string      `xml:"severity"`
	Subdomain   string      `xml:"subdomain"`
}

type Parameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

/*
=====================================================
*/

type ResponseAspectHistogramContainer struct {
	AspectHistogramContainer AspectHistogramContainer `xml:"aspectHistogramContainer"`
}

// AspectHistogramContainer is response container for aspect histograms.
type AspectHistogramContainer struct {
	Aspects           []Aspect `xml:"aspect"`
	DomainDisplayName string   `xml:"domainDisplayName"`
	DomainName        string   `xml:"domainName"`
}

// Aspect is a characteristic of an item in a domain.
type Aspect struct {
	Name            string           `xml:"name,attr"`
	ValueHistograms []ValueHistogram `xml:"valueHistogram"`
}

// ValueHistogram is a container that returns the name of the respective aspect value and the histogram
//  (the number of available items) that share that item characteristic.
type ValueHistogram struct {
	ValueName string `xml:"valueName,attr"`
	// Count is the number of items that share the characteristic the respective aspect value.
	Count int64 `xml:"count"`
}

/*
=====================================================
*/

type ResponseCategoryHistogramContainer struct {
	CategoryHistogramContainer CategoryHistogramContainer `json:"categoryHistogramContainer" xml:"categoryHistogramContainer"`
}

// CategoryHistogramContainer is a response container for category histograms. Only returned when one or
// more category histograms are returned. A category histogram is not returned if there are no
// matching items or if the search is restricted to a single leaf category.
type CategoryHistogramContainer struct {
	CategoryHistograms []CategoryHistogram `xml:"categoryHistogram"`
}

// CategoryHistogram Statistical (item count) information on the categories that contain
// items that match the search criteria or specified category or categories.
// A category histogram contains information for up to 10 child categories.
// Search result total entries may not necessarily match the sum of category histogram item counts.
type CategoryHistogram struct {
	CategoryId              string              `xml:"categoryId"`
	CategoryName            string              `xml:"categoryName"`
	ChildCategoryHistograms []CategoryHistogram `xml:"childCategoryHistogram"`
	Count                   int64               `xml:"count"`
}

/*
=====================================================
*/

type ResponseConditionHistogramContainer struct {
	ConditionHistogramContainer ConditionHistogramContainer `json:"conditionHistogramContainer" xml:"conditionHistogramContainer"`
}

// ConditionHistogramContainer is a response container for condition histograms.
// Not returned when Condition is specified in itemFilter.
// That is, only returned when you have not yet narrowed your search based on specific conditions.
type ConditionHistogramContainer struct {
	ConditionHistograms []ConditionHistogram `xml:"conditionHistogram"`
}

// ConditionHistogram Statistical (item count) information on the condition of items that match
// the search criteria (or specified category).
// For example, the number of brand new items that match the query.
type ConditionHistogram struct {
	Condition Condition `xml:"condition"`
	Count     int       `xml:"count"`
}

/*
=====================================================
*/

type ResponsePaginationOutput struct {
	PaginationOutput PaginationOutput `xml:"paginationOutput"`
}

// PaginationOutput Indicates the pagination of the result set. Child elements indicate the page
// number that is returned, the maximum number of item listings to return per page,
// total number of pages that can be returned, and the total number of listings that
// match the search criteria.
type PaginationOutput struct {
	PageNumber     int `xml:"pageNumber"`
	EntriesPerPage int `xml:"entriesPerPage"`
	TotalPages     int `xml:"totalPages"`
	TotalEntries   int `xml:"totalEntries"`
}

/*
=====================================================
*/

type ResponseSearchResult struct {
	SearchResult SearchResult `xml:"searchResult"`
}

// SearchResult is a container for the item listings that matched the search criteria.
// The data for each item is returned in individual containers, if any matches were found.
type SearchResult struct {
	Count int    `xml:"count,attr"`
	Items []Item `xml:"item"`
}

// Item is a container for the data of a single item that matches the search criteria.
type Item struct {
	Attributes             []ItemAttribute      `xml:"attribute"`
	CharityID              string               `xml:"charityId"`
	Condition              Condition            `xml:"condition"`
	Country                string               `xml:"country"`
	DiscountPriceInfo      DiscountPriceInfo    `xml:"discountPriceInfo"`
	Distance               float64              `xml:"distance"`
	DistanceUnit           string               `xml:"distance,attr"`
	EekStatuses            []string             `xml:"eekStatus"`
	GalleryInfoContainer   GalleryInfoContainer `xml:"galleryInfoContainer"`
	GalleryPlusPictureURLs []string             `xml:"galleryPlusPictureURL"`
	GalleryURL             string               `xml:"galleryURL"`
	GlobalID               string               `xml:"globalId"`
	ItemID                 string               `xml:"itemId"`
	ListingInfo            ListingInfo          `xml:"listingInfo"`
	Location               string               `xml:"location"`
	PaymentMethods         []string             `xml:"paymentMethod"`
	PictureURLLarge        string               `xml:"pictureURLLarge"`
	PictureURLSuperSize    string               `xml:"pictureURLSuperSize"`
	PostalCode             string               `xml:"postalCode"`
	PrimaryCategory        Category             `xml:"primaryCategory"`
	ProductID              string               `xml:"productId"`
	SecondaryCategory      Category             `xml:"secondaryCategory"`
	SellerInfo             SellerInfo           `xml:"sellerInfo"`
	SellingStatus          SellingStatus        `xml:"sellingStatus"`
	ShippingInfo           ShippingInfo         `xml:"shippingInfo"`
	StoreInfo              StoreInfo            `xml:"storeInfo"`
	Subtitle               string               `xml:"subtitle"`
	Title                  string               `xml:"title"`
	UnitPrice              UnitPriceInfo        `xml:"unitPrice"`
	ViewItemURL            string               `xml:"viewItemURL"`
	AutoPay                bool                 `xml:"autoPay"`
	EbayPlusEnabled        bool                 `xml:"eBayPlusEnabled"`
	ReturnsAccepted        bool                 `xml:"returnsAccepted"`
	TopRatedListing        bool                 `xml:"topRatedListing"`
}

type ItemAttribute struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Condition struct {
	ConditionDisplayName string `xml:"conditionDisplayName"`
	ConditionId          string `xml:"conditionId"`
}

type DiscountPriceInfo struct {
	MinimumAdvertisedPriceExposure string  `xml:"minimumAdvertisedPriceExposure"`
	OriginalRetailPrice            float64 `xml:"originalRetailPrice"`
	PricingTreatment               string  `xml:"pricingTreatment"`
	SoldOffEbay                    bool    `xml:"soldOffEbay"`
	SoldOnEbay                     bool    `xml:"soldOnEbay"`
}

type GalleryInfoContainer struct {
	GalleryURL  string `xml:"galleryURL"`
	GallerySize string `xml:"gallerySize,attr"`
}

type ListingInfo struct {
	BestOfferEnabled       bool   `xml:"bestOfferEnabled"`
	BuyItNowAvailable      bool   `xml:"buyItNowAvailable"`
	Gift                   bool   `xml:"gift"`
	BuyItNowPrice          Price  `xml:"buyItNowPrice"`
	ConvertedBuyItNowPrice Price  `xml:"convertedBuyItNowPrice"`
	EndTime                string `xml:"endTime"`
	ListingType            string `xml:"listingType"`
	StartTime              string `xml:"startTime"`
	WatchCount             int    `xml:"watch_count"`
}

type Price struct {
	Value      float64 `xml:",cdata"`
	CurrencyID string  `xml:"currencyId,attr"`
}

type Category struct {
	CategoryId   string `xml:"categoryId"`
	CategoryName string `xml:"categoryName"`
}

type SellerInfo struct {
	FeedbackRatingStar      string  `xml:"feedbackRatingStar"`
	FeedbackScore           int64   `xml:"feedbackScore"`
	PositiveFeedbackPercent float64 `xml:"positiveFeedbackPercent"`
	SellerUserName          string  `xml:"sellerUserName"`
	TopRatedSeller          bool    `xml:"topRatedSeller"`
}

type SellingStatus struct {
	BidCount              int    `xml:"bidCount"`
	ConvertedCurrentPrice Price  `xml:"convertedCurrentPrice"`
	CurrentPrice          Price  `xml:"currentPrice"`
	SellingState          string `xml:"sellingState"`
	TimeLeft              string `xml:"timeLeft"`
}

type ShippingInfo struct {
	ExpeditedShipping       bool     `xml:"expeditedShipping"`
	OneDayShippingAvailable bool     `xml:"oneDayShippingAvailable"`
	HandlingTime            int      `xml:"handlingTime"`
	ShippingServiceCost     Price    `xml:"shippingServiceCost"`
	ShippingType            string   `xml:"shippingType"`
	ShipToLocations         []string `xml:"shipToLocations"`
}

type StoreInfo struct {
	StoreName string `xml:"storeName"`
	StoreURL  string `xml:"storeURL"`
}

type UnitPriceInfo struct {
	Quantity int64  `xml:"quantity"`
	Type     string `xml:"type"`
}
