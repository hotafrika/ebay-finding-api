package finding

type responseStandard struct {
	Ack          string  `xml:"ack"`
	ErrorMessage []Error `xml:"errorMessage>error"`
	Timestamp    string  `xml:"timestamp"`
	Version      string  `xml:"version"`
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
	AspectHistogramContainer AspectHistogramContainer `json:"aspectHistogramContainer" xml:"aspectHistogramContainer"`
}

type AspectHistogramContainer struct {
	// TODO
}

/*
=====================================================
*/

type ResponseCategoryHistogramContainer struct {
	CategoryHistogramContainer CategoryHistogramContainer `json:"categoryHistogramContainer" xml:"categoryHistogramContainer"`
}

type CategoryHistogramContainer struct {
	// TODO
}

/*
=====================================================
*/

type ResponseConditionHistogramContainer struct {
	ConditionHistogramContainer ConditionHistogramContainer `json:"conditionHistogramContainer" xml:"conditionHistogramContainer"`
}

type ConditionHistogramContainer struct {
	// TODO
}

/*
=====================================================
*/

type ResponsePaginationOutput struct {
	PaginationOutput PaginationOutput `xml:"paginationOutput"`
}

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

type SearchResult struct {
	Count int    `xml:"count,attr"`
	Items []Item `xml:"item"`
}

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
