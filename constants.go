package ebay_finding_api

const EbayFindingAPIVersion = "1.13.0"
const EbayRequestDataFormat = "JSON"
const EbayResponseDataFormat = "JSON"
const EbayServiceName = "FindingService"
const DefaultItemsPerPage = 100

type EbayEndpoint string

const (
	EbayEndpointProduction = "https://svcs.ebay.com/services/search/FindingService/v1"
	EbayEndpointSandbox    = "https://svcs.sandbox.ebay.com/services/search/FindingService/v1"
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
	ItemFilterAuthorizedSellerOnly  ItemFilterParameter = "AuthorizedSellerOnly" // done
	ItemFilterAvailableTo           ItemFilterParameter = "AvailableTo"          // done
	ItemFilterBestOfferOnly         ItemFilterParameter = "BestOfferOnly"        // done
	ItemFilterCharityOnly           ItemFilterParameter = "CharityOnly"          // done
	ItemFilterCondition             ItemFilterParameter = "Condition"            // done
	ItemFilterCurrency              ItemFilterParameter = "Currency"             // done
	ItemFilterEndTimeFrom           ItemFilterParameter = "EndTimeFrom"
	ItemFilterEndTimeTo             ItemFilterParameter = "EndTimeTo"
	ItemFilterExcludeAutoPay        ItemFilterParameter = "ExcludeAutoPay"        // done
	ItemFilterExcludeCategory       ItemFilterParameter = "ExcludeCategory"       // done
	ItemFilterExcludeSeller         ItemFilterParameter = "ExcludeSeller"         // done
	ItemFilterExpeditedShippingType ItemFilterParameter = "ExpeditedShippingType" // done
	ItemFilterFeaturedOnly          ItemFilterParameter = "FeaturedOnly"          // done
	ItemFilterFeedbackScoreMax      ItemFilterParameter = "FeedbackScoreMax"      // done
	ItemFilterFeedbackScoreMin      ItemFilterParameter = "FeedbackScoreMin"      // done
	ItemFilterFreeShippingOnly      ItemFilterParameter = "FreeShippingOnly"      // done
	ItemFilterGetItFastOnly         ItemFilterParameter = "GetItFastOnly"         // done
	ItemFilterHideDuplicateItems    ItemFilterParameter = "HideDuplicateItems"    // done
	ItemFilterListedIn              ItemFilterParameter = "ListedIn"              // done
	ItemFilterListingType           ItemFilterParameter = "ListingType"           // done
	ItemFilterLocalPickupOnly       ItemFilterParameter = "LocalPickupOnly"       // done
	ItemFilterLocalSearchOnly       ItemFilterParameter = "LocalSearchOnly"       // done
	ItemFilterLocatedIn             ItemFilterParameter = "LocatedIn"             // done
	ItemFilterLotsOnly              ItemFilterParameter = "LotsOnly"              // done
	ItemFilterMaxBids               ItemFilterParameter = "MaxBids"               // done
	ItemFilterMaxDistance           ItemFilterParameter = "MaxDistance"           // done
	ItemFilterMaxHandlingTime       ItemFilterParameter = "MaxHandlingTime"       // done
	ItemFilterMaxPrice              ItemFilterParameter = "MaxPrice"
	ItemFilterMaxQuantity           ItemFilterParameter = "MaxQuantity" // done
	ItemFilterMinBids               ItemFilterParameter = "MinBids"     // done
	ItemFilterMinPrice              ItemFilterParameter = "MinPrice"
	ItemFilterMinQuantity           ItemFilterParameter = "MinQuantity" // done
	ItemFilterModTimeFrom           ItemFilterParameter = "ModTimeFrom"
	ItemFilterOutletSellerOnly      ItemFilterParameter = "OutletSellerOnly"    // done
	ItemFilterPaymentMethod         ItemFilterParameter = "PaymentMethod"       // done
	ItemFilterReturnsAcceptedOnly   ItemFilterParameter = "ReturnsAcceptedOnly" // done
	ItemFilterSeller                ItemFilterParameter = "Seller"              // done
	ItemFilterSellerBusinessType    ItemFilterParameter = "SellerBusinessType"  // done
	ItemFilterSoldItemsOnly         ItemFilterParameter = "SoldItemsOnly"       // done
	ItemFilterStartTimeFrom         ItemFilterParameter = "StartTimeFrom"
	ItemFilterStartTimeTo           ItemFilterParameter = "StartTimeTo"
	ItemFilterTopRatedSellerOnly    ItemFilterParameter = "TopRatedSellerOnly" // done
	ItemFilterValueBoxInventory     ItemFilterParameter = "ValueBoxInventory"  // done
	ItemFilterWorldOfGoodOnly       ItemFilterParameter = "WorldOfGoodOnly"    // done
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
