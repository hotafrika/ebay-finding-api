package finding

const EbayFindingAPIVersion = "1.13.0"
const EbayRequestDataFormat = "JSON"
const EbayResponseDataFormat = "XML"
const EbayServiceName = "FindingService"
const DefaultItemsPerPage = 100

type EbayEndpoint string

const (
	// EbayEndpointProduction is a production endpoint for Finding API
	EbayEndpointProduction = "https://svcs.ebay.com/services/search/FindingService/v1"
	// EbayEndpointSandbox is a sandbox endpoint for Finding API
	EbayEndpointSandbox = "https://svcs.sandbox.ebay.com/services/search/FindingService/v1"
)

type EbayOperation string

const (
	OperationFindItemsAdvanced               EbayOperation = "findItemsAdvanced"
	OperationFindItemsByCategory             EbayOperation = "findItemsByCategory"
	OperationFindItemsByKeywords             EbayOperation = "findItemsByKeywords"
	OperationFindItemsByProduct              EbayOperation = "findItemsByProduct"
	OperationFindItemsIneBayStores           EbayOperation = "findItemsIneBayStores"
	OperationGetSearchKeywordsRecommendation EbayOperation = "getSearchKeywordsRecommendation"
	OperationGetHistograms                   EbayOperation = "getHistograms"
	OperationGetVersion                      EbayOperation = "getVersion"
)

type GlobalID string

const (
	GlobalIDEbayUS    GlobalID = "EBAY-US"
	GlobalIDEbayENCA  GlobalID = "EBAY-ENCA"
	GlobalIDEbayGB    GlobalID = "EBAY-GB"
	GlobalIDEbayAU    GlobalID = "EBAY-AU"
	GlobalIDEbayAT    GlobalID = "EBAY-AT"
	GlobalIDEbayFRBE  GlobalID = "EBAY-FRBE"
	GlobalIDEbayFR    GlobalID = "EBAY-FR"
	GlobalIDEbayDE    GlobalID = "EBAY-DE"
	GlobalIDEbayMOTOR GlobalID = "EBAY-MOTOR"
	GlobalIDEbayIT    GlobalID = "EBAY-IT"
	GlobalIDEbayNLBE  GlobalID = "EBAY-NLBE"
	GlobalIDEbayNL    GlobalID = "EBAY-NL"
	GlobalIDEbayES    GlobalID = "EBAY-ES"
	GlobalIDEbayCH    GlobalID = "EBAY-CH"
	GlobalIDEbayHK    GlobalID = "EBAY-HK"
	GlobalIDEbayIN    GlobalID = "EBAY-IN"
	GlobalIDEbayIE    GlobalID = "EBAY-IE"
	GlobalIDEbayMY    GlobalID = "EBAY-MY"
	GlobalIDEbayFRCA  GlobalID = "EBAY-FRCA"
	GlobalIDEbayPH    GlobalID = "EBAY-PH"
	GlobalIDEbaySG    GlobalID = "EBAY-SG"
)

type OutputSelectorParameter string

const (
	OutputSelectorAspectHistogram     OutputSelectorParameter = "AspectHistogram"
	OutputSelectorCategoryHistogram   OutputSelectorParameter = "CategoryHistogram"
	OutputSelectorConditionHistogram  OutputSelectorParameter = "ConditionHistogram"
	OutputSelectorGalleryInfo         OutputSelectorParameter = "GalleryInfo"
	OutputSelectorPictureURLLarge     OutputSelectorParameter = "PictureURLLarge"
	OutputSelectorPictureURLSuperSize OutputSelectorParameter = "PictureURLSuperSize"
	OutputSelectorSellerInfo          OutputSelectorParameter = "SellerInfo"
	OutputSelectorStoreInfo           OutputSelectorParameter = "StoreInfo"
	OutputSelectorUnitPriceInfo       OutputSelectorParameter = "UnitPriceInfo"
)

type AspectNameParameter string

type ItemFilterParameter string

const (
	ItemFilterAuthorizedSellerOnly  ItemFilterParameter = "AuthorizedSellerOnly"
	ItemFilterAvailableTo           ItemFilterParameter = "AvailableTo"
	ItemFilterBestOfferOnly         ItemFilterParameter = "BestOfferOnly"
	ItemFilterCharityOnly           ItemFilterParameter = "CharityOnly"
	ItemFilterCondition             ItemFilterParameter = "Condition"
	ItemFilterCurrency              ItemFilterParameter = "Currency"
	ItemFilterEndTimeFrom           ItemFilterParameter = "EndTimeFrom"
	ItemFilterEndTimeTo             ItemFilterParameter = "EndTimeTo"
	ItemFilterExcludeAutoPay        ItemFilterParameter = "ExcludeAutoPay"
	ItemFilterExcludeCategory       ItemFilterParameter = "ExcludeCategory"
	ItemFilterExcludeSeller         ItemFilterParameter = "ExcludeSeller"
	ItemFilterExpeditedShippingType ItemFilterParameter = "ExpeditedShippingType"
	ItemFilterFeaturedOnly          ItemFilterParameter = "FeaturedOnly"
	ItemFilterFeedbackScoreMax      ItemFilterParameter = "FeedbackScoreMax"
	ItemFilterFeedbackScoreMin      ItemFilterParameter = "FeedbackScoreMin"
	ItemFilterFreeShippingOnly      ItemFilterParameter = "FreeShippingOnly"
	ItemFilterGetItFastOnly         ItemFilterParameter = "GetItFastOnly"
	ItemFilterHideDuplicateItems    ItemFilterParameter = "HideDuplicateItems"
	ItemFilterListedIn              ItemFilterParameter = "ListedIn"
	ItemFilterListingType           ItemFilterParameter = "ListingType"
	ItemFilterLocalPickupOnly       ItemFilterParameter = "LocalPickupOnly"
	ItemFilterLocalSearchOnly       ItemFilterParameter = "LocalSearchOnly"
	ItemFilterLocatedIn             ItemFilterParameter = "LocatedIn"
	ItemFilterLotsOnly              ItemFilterParameter = "LotsOnly"
	ItemFilterMaxBids               ItemFilterParameter = "MaxBids"
	ItemFilterMaxDistance           ItemFilterParameter = "MaxDistance"
	ItemFilterMaxHandlingTime       ItemFilterParameter = "MaxHandlingTime"
	ItemFilterMaxPrice              ItemFilterParameter = "MaxPrice"
	ItemFilterMaxQuantity           ItemFilterParameter = "MaxQuantity"
	ItemFilterMinBids               ItemFilterParameter = "MinBids"
	ItemFilterMinPrice              ItemFilterParameter = "MinPrice"
	ItemFilterMinQuantity           ItemFilterParameter = "MinQuantity"
	ItemFilterModTimeFrom           ItemFilterParameter = "ModTimeFrom"
	ItemFilterOutletSellerOnly      ItemFilterParameter = "OutletSellerOnly"
	ItemFilterPaymentMethod         ItemFilterParameter = "PaymentMethod"
	ItemFilterReturnsAcceptedOnly   ItemFilterParameter = "ReturnsAcceptedOnly"
	ItemFilterSeller                ItemFilterParameter = "Seller"
	ItemFilterSellerBusinessType    ItemFilterParameter = "SellerBusinessType"
	ItemFilterSoldItemsOnly         ItemFilterParameter = "SoldItemsOnly"
	ItemFilterStartTimeFrom         ItemFilterParameter = "StartTimeFrom"
	ItemFilterStartTimeTo           ItemFilterParameter = "StartTimeTo"
	ItemFilterTopRatedSellerOnly    ItemFilterParameter = "TopRatedSellerOnly"
	ItemFilterValueBoxInventory     ItemFilterParameter = "ValueBoxInventory"
	ItemFilterWorldOfGoodOnly       ItemFilterParameter = "WorldOfGoodOnly"
)

type SortOrderParameter string

const (
	SortOrderBestMatch                SortOrderParameter = "BestMatch"
	SortOrderBidCountFewest           SortOrderParameter = "BidCountFewest"
	SortOrderBidCountMost             SortOrderParameter = "BidCountMost"
	SortOrderCountryAscending         SortOrderParameter = "CountryAscending"
	SortOrderCountryDescending        SortOrderParameter = "CountryDescending"
	SortOrderCurrentPriceHighest      SortOrderParameter = "CurrentPriceHighest"
	SortOrderDistanceNearest          SortOrderParameter = "DistanceNearest"
	SortOrderEndTimeSoonest           SortOrderParameter = "EndTimeSoonest"
	SortOrderPricePlusShippingHighest SortOrderParameter = "PricePlusShippingHighest"
	SortOrderPricePlusShippingLowest  SortOrderParameter = "PricePlusShippingLowest"
	SortOrderStartTimeNewest          SortOrderParameter = "StartTimeNewest"
	SortOrderWatchCountDecreaseSort   SortOrderParameter = "WatchCountDecreaseSort"
)

type ItemFilterCurrencyIDOption string

const (
	CurrencyIDAUD ItemFilterCurrencyIDOption = "AUD"
	CurrencyIDCAD ItemFilterCurrencyIDOption = "CAD"
	CurrencyIDCHF ItemFilterCurrencyIDOption = "CHF"
	CurrencyIDCNY ItemFilterCurrencyIDOption = "CNY"
	CurrencyIDEUR ItemFilterCurrencyIDOption = "EUR"
	CurrencyIDGBP ItemFilterCurrencyIDOption = "GBP"
	CurrencyIDHKD ItemFilterCurrencyIDOption = "HKD"
	CurrencyIDINR ItemFilterCurrencyIDOption = "INR"
	CurrencyIDMYR ItemFilterCurrencyIDOption = "MYR"
	CurrencyIDPHP ItemFilterCurrencyIDOption = "PHP"
	CurrencyIDPLN ItemFilterCurrencyIDOption = "PLN"
	CurrencyIDSEK ItemFilterCurrencyIDOption = "SEK"
	CurrencyIDSGD ItemFilterCurrencyIDOption = "SGD"
	CurrencyIDTWD ItemFilterCurrencyIDOption = "TWD"
	// CurrencyIDUSD is US Dollar.
	// For eBay, you can only specify this currency for listings you submit to the US (site ID 0), eBayMotors (site 100), and Canada (site 2) sites.
	CurrencyIDUSD ItemFilterCurrencyIDOption = "USD"
)

type ItemFilterExpeditedShippingTypeOption string

const (
	ExpeditedShippingTypeExpedited      ItemFilterExpeditedShippingTypeOption = "Expedited"
	ExpeditedShippingTypeOneDayShipping ItemFilterExpeditedShippingTypeOption = "OneDayShipping"
)

type ItemFilterPaymentMethodOption string

const (
	PaymentMethodPayPal      ItemFilterPaymentMethodOption = "PayPal"
	PaymentMethodPaisaPay    ItemFilterPaymentMethodOption = "PaisaPay"
	PaymentMethodPaisaPayEMI ItemFilterPaymentMethodOption = "PaisaPayEMI"
)

type ItemFilterSellerBusinessTypeOption string

const (
	SellerBusinessTypeBusiness ItemFilterSellerBusinessTypeOption = "Business"
	SellerBusinessTypePrivate  ItemFilterSellerBusinessTypeOption = "Private"
)

type ItemFilterConditionOption string

const (
	ConditionNew                  ItemFilterConditionOption = "1000"
	ConditionNewOther             ItemFilterConditionOption = "1500"
	ConditionNewWithDefects       ItemFilterConditionOption = "1750"
	ConditionCertifiedRefurbished ItemFilterConditionOption = "2000"
	// ConditionExcellentRefurbished is used in Cell Phones & Smartphones category (9355)
	// on the US, Canada, UK, Germany, and Australia
	ConditionExcellentRefurbished ItemFilterConditionOption = "2010"
	// ConditionVeryGoodRefurbished is used in Cell Phones & Smartphones category (9355)
	// on the US, Canada, UK, Germany, and Australia
	ConditionVeryGoodRefurbished ItemFilterConditionOption = "2020"
	// ConditionGoodRefurbished is used in Cell Phones & Smartphones category (9355)
	// on the US, Canada, UK, Germany, and Australia
	ConditionGoodRefurbished ItemFilterConditionOption = "2030"
	// ConditionSellerRefurbished can't be used in Cell Phones & Smartphones category (9355)
	// on the US, Canada, UK, Germany, and Australia
	ConditionSellerRefurbished    ItemFilterConditionOption = "2500"
	ConditionLikeNew              ItemFilterConditionOption = "2750"
	ConditionUsed                 ItemFilterConditionOption = "3000"
	ConditionVeryGood             ItemFilterConditionOption = "4000"
	ConditionGood                 ItemFilterConditionOption = "5000"
	ConditionAcceptable           ItemFilterConditionOption = "6000"
	ConditionForPartsOrNotWorking ItemFilterConditionOption = "7000"
)

type ItemFilterListingTypeOption string

const (
	ListingTypeAuction        ItemFilterListingTypeOption = "Auction"
	ListingTypeAuctionWithBIN ItemFilterListingTypeOption = "AuctionWithBIN"
	ListingTypeClassified     ItemFilterListingTypeOption = "Classified"
	ListingTypeFixedPrice     ItemFilterListingTypeOption = "FixedPrice"
	ListingTypeStoreInventory ItemFilterListingTypeOption = "StoreInventory"
	ListingTypeAll            ItemFilterListingTypeOption = "All"
)

type ProductTypeOption string

const (
	ProductTypeISBN        ProductTypeOption = "ISBN"
	ProductTypeUPC         ProductTypeOption = "UPC"
	ProductTypeEAN         ProductTypeOption = "EAN"
	ProductTypeReferenceID ProductTypeOption = "ReferenceID"
)
