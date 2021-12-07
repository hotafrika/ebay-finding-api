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
	SearchResultCount int    `xml:"count,attr"`
	Items             []Item `xml:"item"`
}

type Item struct {
	Attributes             []ItemAttribute      `xml:"attribute"`
	AutoPay                bool                 `xml:"autoPay"`
	CharityID              string               `xml:"charityId"`
	Condition              Condition            `xml:"condition"`
	Country                string               `xml:"country"`
	DiscountPriceInfo      DiscountPriceInfo    `xml:"discountPriceInfo"`
	Distance               float64              `xml:"distance"`
	DistanceUnit           string               `xml:"distance,attr"`
	EbayPlusEnabled        bool                 `xml:"eBayPlusEnabled"`
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
	ReturnsAccepted        bool                 `xml:"returnsAccepted"`
	SecondaryCategory      Category             `xml:"secondaryCategory"`
	SellerInfo             SellerInfo           `xml:"sellerInfo"`
	SellingStatus          SellingStatus        `xml:"sellingStatus"`
	ShippingInfo           ShippingInfo         `xml:"shippingInfo"`
	StoreInfo              StoreInfo            `xml:"storeInfo"`
	Subtitle               string               `xml:"subtitle"`
	Title                  string               `xml:"title"`
	TopRatedListing        bool                 `xml:"topRatedListing"`
	UnitPrice              UnitPriceInfo        `xml:"unitPrice"`
	ViewItemURL            string               `xml:"viewItemURL"`
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
	// TODO
}

type GalleryInfoContainer struct {
	GalleryURL  string `xml:"galleryURL"`
	gallerySize string `xml:"gallerySize,attr"`
}

type ListingInfo struct {
	// TODO
}

type Category struct {
	CategoryId   string `xml:"categoryId"`
	CategoryName string `xml:"categoryName"`
}

type SellerInfo struct {
	// TODO
}

type SellingStatus struct {
	// TODO
}

type ShippingInfo struct {
	// TODO
}

type StoreInfo struct {
	// TODO
}

type UnitPriceInfo struct {
	Quantity int64  `xml:"quantity"`
	Type     string `xml:"type"`
}
