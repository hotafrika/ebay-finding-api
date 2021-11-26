package ebay_finding_api

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type ServiceRequest struct {
	CategoryID        []string `json:"categoryId,omitempty"`
	Keywords          string   `json:"keywords,omitempty"`
	DescriptionSearch bool     `json:"descriptionSearch,omitempty"`
	//isSellerFilterApplied bool     // internal use; to check if request valid
	itemFilterMap   map[ItemFilterParameter]ServiceItemFilter
	OutputSelector  []string               `json:"outputSelector,omitempty"`
	ItemFilter      []ServiceItemFilter    `json:"itemFilter,omitempty"`
	PaginationInput ServicePaginationInput `json:"paginationInput,omitempty"`
	SortOrder       string                 `json:"sortOrder,omitempty"`
	//AspectFilter          []ServiceAspectFilter   `json:"aspectFilter,omitempty"` 	// removed
	//Affiliate				ServiceAffiliate		`json:"affiliate,omitempty"`		// to implement
	//BuyerPostalCode 		string 					`json:"buyerPostalCode,omitempty"`	// to implement

	client *resty.Client
}

// ServiceAspectFilter represents AspectFilter
//  It is NOT implemented yet
type ServiceAspectFilter struct {
	AspectName      string   `json:"aspectName"`
	AspectValueName []string `json:"aspectValueName"`
}

// ServiceItemFilter represents ItemFilter
type ServiceItemFilter struct {
	Name       string   `json:"name"`
	Value      []string `json:"value"`
	ParamName  string   `json:"paramName,omitempty"`
	ParamValue string   `json:"paramValue,omitempty"`
}

// ServicePaginationInput represents PaginationInput
type ServicePaginationInput struct {
	EntriesPerPage int `json:"entriesPerPage,omitempty"`
	PageNumber     int `json:"pageNumber,omitempty"`
}

func (sr *ServiceRequest) prepare() {
	if len(sr.itemFilterMap) == 0 {
		return
	}
	var filter []ServiceItemFilter
	for _, val := range sr.itemFilterMap {
		filter = append(filter, val)
	}
	sr.ItemFilter = filter
}

func (sr *ServiceRequest) Reload() *ServiceRequest {
	sr.ItemFilter = nil
	sr.prepare()
	return sr
}

func (sr *ServiceRequest) Get(page int) (ServiceResponse, error){
	if len(sr.ItemFilter) == 0 {
		sr.prepare()
	}
	sr.PaginationInput.PageNumber = page
	// TODO client work here
	return ServiceResponse{}, nil
}

func (sr *ServiceRequest) First() (ServiceResponse, error){
	return sr.Get(1)
}

func (sr *ServiceRequest) RequestBody() ([]byte, error) {
	return json.MarshalIndent(sr, "", "  ")
}


// WithPageLimit sets page limit to list of items
//  Min: 1. Max: 100. Default: 100.
func (sr *ServiceRequest) WithPageLimit(limit int) *ServiceRequest {
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

// WithOutputSelectors adds OutputSelectors
func (sr *ServiceRequest) WithOutputSelectors(osps ...OutputSelectorParameter) *ServiceRequest {
	for _, osp := range osps {
		sr.OutputSelector = append(sr.OutputSelector, string(osp))
	}
	return sr
}

// WithSortOrder sorts the returned items according to a single specified sort order.
// Default: BestMatch.
func (sr *ServiceRequest) WithSortOrder(order SortOrderParameter) *ServiceRequest {
	sr.SortOrder = string(order)
	return sr
}

// WithCategoryIDInt adds category for searching (up to 3 categories)
//  4th and next categoryIDs will be skipped silently
func (sr *ServiceRequest) WithCategoryIDInt(categoryID int) *ServiceRequest {
	if len(sr.CategoryID) >= 3 {
		return sr
	}
	sr.CategoryID = append(sr.CategoryID, strconv.Itoa(categoryID))
	return sr
}

// WithCategoryID adds category for searching (up to 3 categories)
//  4th and next categoryIDs will be skipped silently
func (sr *ServiceRequest) WithCategoryID(categoryID string) *ServiceRequest {
	if len(sr.CategoryID) >= 3 {
		return sr
	}
	sr.CategoryID = append(sr.CategoryID, categoryID)
	return sr
}

// WithCategoriesID adds categories for searching (up to 3 categories)
//  4th and next categoryIDs will be skipped silently
func (sr *ServiceRequest) WithCategoriesID(categoriesID ...string) *ServiceRequest {
	for _, categoryID := range categoriesID {
		if len(sr.CategoryID) >= 3 {
			return sr
		}
		sr.CategoryID = append(sr.CategoryID, categoryID)
	}
	return sr
}

// WithKeywords adds keywords for searching
// Max length: 350. The maximum length for a single word is 98. Min length: 2.
// Key longer than 350 symbols will be trimmed. Key shorter 2 (0 < n < 2) symbols won't change Keywords field.
//  Empty string remove Keywords field.
func (sr *ServiceRequest) WithKeywords(key string) *ServiceRequest {
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

// WithDescriptionSearch specifies whether your keyword query should be applied
// to item descriptions in addition to titles
func (sr *ServiceRequest) WithDescriptionSearch(ds bool) *ServiceRequest {
	sr.DescriptionSearch = ds
	return sr
}

// WithItemFilterTopRatedSellerOnly adds TopRatedSellerOnly ItemFilter
// The TopRatedSellerOnly item filter cannot be used together with either the Seller or ExcludeSeller item filters.
// The TopRatedSellerOnly item filter is supported for the following sites only:
// US (EBAY-US), Motors (EBAY-MOTOR), UK (EBAY-GB), IE (EBAY-IE), DE (EBAY-DE), AT (EBAY-AT), and CH (EBAY-CH).
func (sr *ServiceRequest) WithItemFilterTopRatedSellerOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterTopRatedSellerOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterAuthorizedSellerOnly adds AuthorizedSellerOnly ItemFilter
// If set to true, returns only items listed by authorized sellers
func (sr *ServiceRequest) WithItemFilterAuthorizedSellerOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterAuthorizedSellerOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterBestOfferOnly adds BestOfferOnly ItemFilter
// If true, the search results are limited to only items that have Best Offer enabled.
// Default is false.
func (sr *ServiceRequest) WithItemFilterBestOfferOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterBestOfferOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterCharityOnly adds CharityOnly ItemFilter
// If true, the search results are limited to items for which all or part of the proceeds are given to a charity.
// Each item in the search results will include the ID of the given charity.
// Default is false.
func (sr *ServiceRequest) WithItemFilterCharityOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterCharityOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterExcludeAutoPay adds ExcludeAutoPay ItemFilter
// If true, excludes all items requiring immediate payment. Default is false.
func (sr *ServiceRequest) WithItemFilterExcludeAutoPay(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterExcludeAutoPay, strconv.FormatBool(b))
	return sr
}

// WithItemFilterFeaturedOnly adds FeaturedOnly ItemFilter
// If true, the search results are limited to featured item listings only. Default is false.
func (sr *ServiceRequest) WithItemFilterFeaturedOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterFeaturedOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterFreeShippingOnly adds FreeShippingOnly ItemFilter
// If true, the search results are limited to only items with free shipping to the site specified in the request
// (see Global ID Values).
// Default is false.
func (sr *ServiceRequest) WithItemFilterFreeShippingOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterFreeShippingOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterGetItFastOnly adds GetItFastOnly ItemFilter
// If true, the search results are limited to only Get It Fast listings. Default is false.
func (sr *ServiceRequest) WithItemFilterGetItFastOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterGetItFastOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterHideDuplicateItems adds HideDuplicateItems ItemFilter
// If true, and there are duplicate items for an item in the search results,
// the subsequent duplicates will not appear in the results. Default is false.
func (sr *ServiceRequest) WithItemFilterHideDuplicateItems(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterHideDuplicateItems, strconv.FormatBool(b))
	return sr
}

// WithItemFilterLocalPickupOnly adds LocalPickupOnly ItemFilter
// 	If true, the search results are limited to only items which have local pickup available. Default is false.
func (sr *ServiceRequest) WithItemFilterLocalPickupOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterLocalPickupOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterLocalSearchOnly adds LocalSearchOnly ItemFilter
// If true, the search results are limited to only matching items with the Local Inventory Listing Options (LILO).
// Must be used together with the MaxDistance item filter, and the request must also specify buyerPostalCode.
// Currently, this is only available for the Motors site (global ID EBAY- MOTOR).
func (sr *ServiceRequest) WithItemFilterLocalSearchOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterLocalSearchOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterLotsOnly adds LotsOnly ItemFilter
// If true, the search results are limited to only matching listings for which the lot size is 2 or more.
// Default is false.
func (sr *ServiceRequest) WithItemFilterLotsOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterLotsOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterOutletSellerOnly adds OutletSellerOnly ItemFilter
// If set to true, returns only items listed by outlet sellers.
func (sr *ServiceRequest) WithItemFilterOutletSellerOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterOutletSellerOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterReturnsAcceptedOnly adds ReturnsAcceptedOnly ItemFilter
// If set to true, returns only items where the seller accepts returns.
func (sr *ServiceRequest) WithItemFilterReturnsAcceptedOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterReturnsAcceptedOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterSoldItemsOnly adds SoldItemsOnly ItemFilter
// Reserved for future use. If true, excludes all completed items which are not ended by being sold.
func (sr *ServiceRequest) WithItemFilterSoldItemsOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterSoldItemsOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterWorldOfGoodOnly adds WorldOfGoodOnly ItemFilter
// If true, the search results are limited to only items listed in the World of Good marketplace.
// Defaults to false.
func (sr *ServiceRequest) WithItemFilterWorldOfGoodOnly(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterWorldOfGoodOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterAvailableTo adds AvailableTo ItemFilter
// Limits items to those available to the specified country only.
// Item filter LocatedIn cannot be used together with item filter AvailableTo.
// Expects the two-letter ISO 3166 country code to indicate the country where the item is located.
// For English names that correspond to each code (e.g., KY="Cayman Islands")
func (sr *ServiceRequest) WithItemFilterAvailableTo(code string) *ServiceRequest {
	sr.updateIFValue(ItemFilterAvailableTo, code)
	return sr
}

// WithItemFilterCurrency adds Currency ItemFilter
// Limits results to items listed with the specified currency only.
func (sr *ServiceRequest) WithItemFilterCurrency(currency ItemFilterCurrencyIDOption) *ServiceRequest {
	sr.updateIFValue(ItemFilterCurrency, string(currency))
	return sr
}

// WithItemFilterExpeditedShippingType adds ExpeditedShippingType ItemFilter
// Specifies the type of expedited shipping. You can specify either Expedited or OneDayShipping.
// Only items that can be shipped by the specified type are returned.
// ExpeditedShippingType is used together with the MaxHandlingTime and ReturnsAcceptedOnly filters
// to filter items for certain kinds of gifting events such as birthdays or holidays where the items must
// be delivered by a certain date.
func (sr *ServiceRequest) WithItemFilterExpeditedShippingType(shipping ItemFilterExpeditedShippingTypeOption) *ServiceRequest {
	sr.updateIFValue(ItemFilterExpeditedShippingType, string(shipping))
	return sr
}

// WithItemFilterFeedbackScoreMax adds FeedbackScoreMax ItemFilter
// Specifies the maximum feedback score of a seller whose items can be included in the response.
// If FeedbackScoreMin is also specified, the FeedbackScoreMax value must be greater than or equal to the
// FeedbackScoreMin value.
//
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterFeedbackScoreMax(score int) *ServiceRequest {
	if score < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterFeedbackScoreMax, strconv.Itoa(score))
	return sr
}

// WithItemFilterFeedbackScoreMin adds FeedbackScoreMin ItemFilter
// Specifies the mininum feedback score of a seller whose items can be included in the response.
// If FeedbackScoreMax is also specified, the FeedbackScoreMax value must be greater than or equal to the
// FeedbackScoreMin value.
//
//  Values below are 0 ignored.
func (sr *ServiceRequest) WithItemFilterFeedbackScoreMin(score int) *ServiceRequest {
	if score < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterFeedbackScoreMin, strconv.Itoa(score))
	return sr
}

// WithItemFilterListedIn adds ListedIn ItemFilter
// The site on which the items were originally listed.
// This can be useful for buyers who wish to see only items on their domestic site
// either for delivery cost reasons or time reasons, such as for gifting occasions like birthdays or holidays.
func (sr *ServiceRequest) WithItemFilterListedIn(globalID GlobalID) *ServiceRequest {
	sr.updateIFValue(ItemFilterListedIn, string(globalID))
	return sr
}

// WithItemFilterMaxBids adds MaxBids ItemFilter
// Limits the results to items with bid counts less than or equal to the specified value.
// If MinBids is also specified, the MaxBids value must be greater than or equal to the MinBids value.
//
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterMaxBids(score int) *ServiceRequest {
	if score < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterMaxBids, strconv.Itoa(score))
	return sr
}

// WithItemFilterMaxDistance adds MaxDistance ItemFilter
// Specifies the maximum distance from the specified postal code (buyerPostalCode) to search for items.
//  The request must also specify buyerPostalCode.
//  Values below 5 are ignored.
func (sr *ServiceRequest) WithItemFilterMaxDistance(score int) *ServiceRequest {
	if score < 5 {
		return sr
	}
	sr.updateIFValue(ItemFilterMaxDistance, strconv.Itoa(score))
	return sr
}

// WithItemFilterMaxHandlingTime adds MaxHandlingTime ItemFilter
// Specifies the maximum number of handling days the seller requires to ship the item.
// Only items with a handling time less than or equal to this number will be returned.
// (The handling time is the amount of time, in days, required by the seller to get the item
// ready to ship and handed off to the actual carrier who does the delivery.
// It does not include the time required by the carrier to deliver the item.
//  Values below 1 are ignored.
func (sr *ServiceRequest) WithItemFilterMaxHandlingTime(score int) *ServiceRequest {
	if score < 1 {
		return sr
	}
	sr.updateIFValue(ItemFilterMaxHandlingTime, strconv.Itoa(score))
	return sr
}

// WithItemFilterMaxQuantity adds MaxQuantity ItemFilter
// Limits the results to listings with a quantity less than or equal to the specified value.
// If MinQuantity is also specified, the MaxQuantity value must be greater than or equal to the MinQuantity value.
//
//  Values below 1 are ignored.
func (sr *ServiceRequest) WithItemFilterMaxQuantity(score int) *ServiceRequest {
	if score < 1 {
		return sr
	}
	sr.updateIFValue(ItemFilterMaxQuantity, strconv.Itoa(score))
	return sr
}

// WithItemFilterMinBids adds MinBids ItemFilter
// Limits the results to items with bid counts greater than or equal to the specified value.
// If MaxBids is also specified, the MaxBids value must be greater than or equal to the MinBids value.
//
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterMinBids(score int) *ServiceRequest {
	if score < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterMinBids, strconv.Itoa(score))
	return sr
}

// WithItemFilterMinQuantity adds MinQuantity ItemFilter
// Limits the results to listings with a quantity greater than or equal to the specified value.
// If MaxQuantity is also specified, the MaxQuantity value must be greater than or equal to the MinQuantity value.
//
//  Values below 1 are ignored.
func (sr *ServiceRequest) WithItemFilterMinQuantity(score int) *ServiceRequest {
	if score < 1 {
		return sr
	}
	sr.updateIFValue(ItemFilterMinQuantity, strconv.Itoa(score))
	return sr
}

// WithItemFilterPaymentMethod adds PaymentMethod ItemFilter
// Limits results to items that accept the specified payment method.
func (sr *ServiceRequest) WithItemFilterPaymentMethod(payment ItemFilterPaymentMethodOption) *ServiceRequest {
	sr.updateIFValue(ItemFilterPaymentMethod, string(payment))
	return sr
}

// WithItemFilterSellerBusinessType adds SellerBusinessType ItemFilter
// Restricts the items to those that are from sellers whose
// business type is the specified value. Only one value can be specified.
func (sr *ServiceRequest) WithItemFilterSellerBusinessType(businessType ItemFilterSellerBusinessTypeOption) *ServiceRequest {
	sr.updateIFValue(ItemFilterSellerBusinessType, string(businessType))
	return sr
}

// WithItemFilterValueBoxInventory adds ValueBoxInventory ItemFilter
//  Coming Soon!
// Coming Soon: This filter can be used in conjunction with the sortOrder PricePlusShippingLowest
// to return competitively priced items from eBay top-rated sellers that have a BuyItNow price,
// with the lowest priced item at the top of the list. This filter returns items from categories
// that are catalog-enabled; items from non catalog-enabled categories are not returned.
// Sellers can use this item filter to determine competitive pricing; buying applications can use
// it to obtain competitive items from top rated sellers that are likely to sell quickly.
//
// If set to 1, the item filter constraints are applied and the items are returned accordingly.
// If set to 0 (zero) the item filter is not applied. Defaults to 0.
func (sr *ServiceRequest) WithItemFilterValueBoxInventory(b bool) *ServiceRequest {
	sr.updateIFValue(ItemFilterValueBoxInventory, strconv.FormatBool(b))
	return sr
}

// WithItemFilterSeller adds Seller ItemFilter
// The Seller item filter cannot be used together with either the ExcludeSeller or TopRatedSellerOnly item filters.
// Multiple values are allowed. Up to 100 sellers can be specified.
//  101th and next sellers will be skipped
func (sr *ServiceRequest) WithItemFilterSeller(sellers ...string) *ServiceRequest {
	sr.addIFValues(ItemFilterSeller, 100, sellers...)
	return sr
}

// WithItemFilterExcludeCategory adds ExcludeCategory ItemFilter
// Specify one or more category IDs. Search results will not include
// items from the specified categories or their child categories.
// Valid category IDs.
//  26th and next categories will be skipped
func (sr *ServiceRequest) WithItemFilterExcludeCategory(categories ...string) *ServiceRequest {
	sr.addIFValues(ItemFilterExcludeCategory, 25, categories...)
	return sr
}

// WithItemFilterExcludeSeller adds ExcludeSeller ItemFilter
// Specify one or more seller names. Search results will not include items from the specified sellers.
// The ExcludeSeller item filter cannot be used together with either the Seller or TopRatedSellerOnly item filters.
//  101th and next sellers will be skipped
func (sr *ServiceRequest) WithItemFilterExcludeSeller(sellers ...string) *ServiceRequest {
	sr.addIFValues(ItemFilterExcludeSeller, 100, sellers...)
	return sr
}

// WithItemFilterCondition adds Condition ItemFilter
// This item condition filter allows a user to filter items based on item condition.
func (sr *ServiceRequest) WithItemFilterCondition(conditions ...ItemFilterConditionOption) *ServiceRequest {
	var values []string
	for _, condition := range conditions {
		values = append(values, string(condition))
	}
	// 14 - quantity of available conditions, but let's use 20
	sr.addIFValues(ItemFilterCondition, 14, values...)
	return sr
}

// WithItemFilterListingType adds ListingType ItemFilter
// Filters items based listing type information. Default behavior is to return all matching items,
// except Store Inventory format listings.
func (sr *ServiceRequest) WithItemFilterListingType(listingTypes ...ItemFilterListingTypeOption) *ServiceRequest {
	var values []string
	for _, listingType := range listingTypes {
		values = append(values, string(listingType))
	}
	// 6 - quantity of available Listing Types
	sr.addIFValues(ItemFilterListingType, 6, values...)
	return sr
}

// WithItemFilterLocatedIn adds LocatedIn ItemFilter
// Limits the result set to just those items located in the specified country.
// Item filter AvailableTo cannot be used together with item filter LocatedIn.
//
// Expects the two-letter ISO 3166 country code to indicate the country where the item is located.
// For English names that correspond to each code (e.g., KY="Cayman Islands")
func (sr *ServiceRequest) WithItemFilterLocatedIn(codes ...string) *ServiceRequest {
	sr.addIFValues(ItemFilterLocatedIn, 25, codes...)
	return sr
}

// WithItemFilterMaxPrice adds MaxPrice ItemFilter
// Specifies the maximum current price an item can have to be included in the response.
// If using with MinPrice to specify a price range, the MaxPrice value must be greater than or equal to MinPrice.
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterMaxPrice(price float64) *ServiceRequest {
	if price < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterMaxPrice, strconv.FormatFloat(price, 'f', 2, 64))
	return sr
}

// WithItemFilterMaxPriceWithCurrency adds MaxPrice ItemFilter
// Specifies the maximum current price an item can have to be included in the response.
// If using with MinPrice to specify a price range, the MaxPrice value must be greater than or equal to MinPrice.
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterMaxPriceWithCurrency(price float64, currency ItemFilterCurrencyIDOption) *ServiceRequest {
	if price < 0 {
		return sr
	}
	sr.updateIFValueWithParameter(ItemFilterMaxPrice, "Currency", string(currency), strconv.FormatFloat(price, 'f', 2, 64))
	return sr
}

// WithItemFilterMinPrice adds MinPrice ItemFilter
// Specifies the minimum current price an item can have to be included in the response.
// If using with MaxPrice to specify a price range, the MaxPrice value must be greater than or equal to MinPrice.
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterMinPrice(price float64) *ServiceRequest {
	if price < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterMinPrice, strconv.FormatFloat(price, 'f', 2, 64))
	return sr
}

// WithItemFilterMinPriceWithCurrency adds MinPrice ItemFilter
// Specifies the minimum current price an item can have to be included in the response.
// If using with MaxPrice to specify a price range, the MaxPrice value must be greater than or equal to MinPrice.
//  Values below 0 are ignored.
func (sr *ServiceRequest) WithItemFilterMinPriceWithCurrency(price float64, currency ItemFilterCurrencyIDOption) *ServiceRequest {
	if price < 0 {
		return sr
	}
	sr.updateIFValueWithParameter(ItemFilterMinPrice, "Currency", string(currency), strconv.FormatFloat(price, 'f', 2, 64))
	return sr
}

// WithItemFilterEndTimeFrom adds EndTimeFrom ItemFilter
// Limits the results to items ending on or after the specified time. Specify a time in the future.
func (sr *ServiceRequest) WithItemFilterEndTimeFrom(datetime string) *ServiceRequest {
	sr.updateIFValue(ItemFilterEndTimeFrom, datetime)
	return sr
}

// WithItemFilterEndTimeTo adds EndTimeTo ItemFilter
// Limits the results to items ending on or before the specified time. Specify a time in the future.
func (sr *ServiceRequest) WithItemFilterEndTimeTo(datetime string) *ServiceRequest {
	sr.updateIFValue(ItemFilterEndTimeTo, datetime)
	return sr
}

// WithItemFilterModTimeFrom adds ModTimeFrom ItemFilter
// Limits the results to active items whose status has changed since the specified time.
// Specify a time in the past. Time must be in GMT.
func (sr *ServiceRequest) WithItemFilterModTimeFrom(datetime string) *ServiceRequest {
	sr.updateIFValue(ItemFilterModTimeFrom, datetime)
	return sr
}

// WithItemFilterStartTimeFrom adds StartTimeFrom ItemFilter
// Limits the results to items started on or after the specified time. Specify a time in the future.
func (sr *ServiceRequest) WithItemFilterStartTimeFrom(datetime string) *ServiceRequest {
	sr.updateIFValue(ItemFilterStartTimeFrom, datetime)
	return sr
}

// WithItemFilterStartTimeTo adds StartTimeTo ItemFilter
// Limits the results to items started on or before the specified time. Specify a time in the future.
func (sr *ServiceRequest) WithItemFilterStartTimeTo(datetime string) *ServiceRequest {
	sr.updateIFValue(ItemFilterStartTimeTo, datetime)
	return sr
}

func (sr *ServiceRequest) prepareIFMap(ifp ItemFilterParameter) {
	if _, ok := sr.itemFilterMap[ifp]; !ok {
		sr.itemFilterMap[ifp] = ServiceItemFilter{Name: string(ifp)}
	}
}

func (sr *ServiceRequest) addIFValues(ifp ItemFilterParameter, limit int, values ...string) {
	if len(values) == 0 {
		return
	}
	sr.prepareIFMap(ifp)
	oldValues := sr.itemFilterMap[ifp].Value
	valuesMap := make(map[string]struct{})
	for _, v := range oldValues {
		valuesMap[v] = struct{}{}
	}
	for _, value := range values {
		if len(valuesMap) >= limit {
			break
		}
		valuesMap[value] = struct{}{}
	}
	newValues := make([]string, len(valuesMap))
	for k := range valuesMap {
		newValues = append(newValues, k)
	}
	sr.itemFilterMap[ifp] = ServiceItemFilter{
		Name:  string(ifp),
		Value: newValues,
	}
}

//func (sr *ServiceRequest) addIFValuesWithParameter(ifp ItemFilterParameter, limit int, paramName string, paramValue string, values ...string) {
//	if len(values) == 0 {
//		return
//	}
//	sr.prepareIFMap(ItemFilterSeller)
//	oldValues := sr.itemFilterMap[ItemFilterSeller].Value
//	for _, value := range values {
//		if len(oldValues) >= limit {
//			break
//		}
//		oldValues = append(oldValues, value)
//	}
//	sr.itemFilterMap[ItemFilterSeller] = ServiceItemFilter{
//		Name:       string(ifp),
//		Value:      oldValues,
//		ParamName:  paramName,
//		ParamValue: paramValue,
//	}
//}

func (sr *ServiceRequest) updateIFValue(ifp ItemFilterParameter, value string) {
	sr.prepareIFMap(ifp)
	sr.itemFilterMap[ifp] = ServiceItemFilter{
		Name:  string(ifp),
		Value: []string{value},
	}
}

func (sr *ServiceRequest) updateIFValueWithParameter(ifp ItemFilterParameter, paramName string, paramValue string, value string) {
	sr.prepareIFMap(ifp)
	sr.itemFilterMap[ifp] = ServiceItemFilter{
		Name:       string(ifp),
		Value:      []string{value},
		ParamName:  paramName,
		ParamValue: paramValue,
	}
}
