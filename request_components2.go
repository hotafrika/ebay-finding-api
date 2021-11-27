package finding

import (
	"strconv"
)

type RequestItemFilter struct {
	ItemFilterMap map[ItemFilterParameter]ServiceItemFilter `json:"-"`
	ItemFilter    []ServiceItemFilter                       `json:"itemFilter,omitempty"`
}

// ServiceItemFilter represents ItemFilter
type ServiceItemFilter struct {
	Name       string   `json:"name"`
	Value      []string `json:"value"`
	ParamName  string   `json:"paramName,omitempty"`
	ParamValue string   `json:"paramValue,omitempty"`
}

func (sr *RequestItemFilter) Initialize() {
	sr.ItemFilterMap = make(map[ItemFilterParameter]ServiceItemFilter)
}

// Reload refreshes itemFilter
func (sr *RequestItemFilter) Reload() {
	sr.ItemFilter = nil
}

func (sr *RequestItemFilter) prepare() {
	if len(sr.ItemFilterMap) == 0 {
		return
	}
	if len(sr.ItemFilter) != 0 {
		return
	}
	var filter []ServiceItemFilter
	for _, val := range sr.ItemFilterMap {
		filter = append(filter, val)
	}
	sr.ItemFilter = filter
}

func (sr *RequestItemFilter) prepareIFMap(ifp ItemFilterParameter) {
	if _, ok := sr.ItemFilterMap[ifp]; !ok {
		sr.ItemFilterMap[ifp] = ServiceItemFilter{Name: string(ifp)}
	}
}

func (sr *RequestItemFilter) addIFValues(ifp ItemFilterParameter, limit int, values ...string) {
	if len(values) == 0 {
		return
	}
	sr.prepareIFMap(ifp)
	oldValues := sr.ItemFilterMap[ifp].Value
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
	newValues := make([]string, 0, len(valuesMap))
	for k := range valuesMap {
		newValues = append(newValues, k)
	}
	sr.ItemFilterMap[ifp] = ServiceItemFilter{
		Name:  string(ifp),
		Value: newValues,
	}
}

//func (sr *RequestItemFilter) addIFValuesWithParameter(ifp ItemFilterParameter, limit int, paramName string, paramValue string, values ...string) {
//	if len(values) == 0 {
//		return
//	}
//	sr.prepareIFMap(ItemFilterSeller)
//	oldValues := sr.ItemFilterMap[ItemFilterSeller].Value
//	for _, value := range values {
//		if len(oldValues) >= limit {
//			break
//		}
//		oldValues = append(oldValues, value)
//	}
//	sr.ItemFilterMap[ItemFilterSeller] = ServiceItemFilter{
//		Name:       string(ifp),
//		Value:      oldValues,
//		ParamName:  paramName,
//		ParamValue: paramValue,
//	}
//}

func (sr *RequestItemFilter) updateIFValue(ifp ItemFilterParameter, value string) {
	sr.prepareIFMap(ifp)
	sr.ItemFilterMap[ifp] = ServiceItemFilter{
		Name:  string(ifp),
		Value: []string{value},
	}
}

func (sr *RequestItemFilter) updateIFValueWithParameter(ifp ItemFilterParameter, paramName string, paramValue string, value string) {
	sr.prepareIFMap(ifp)
	sr.ItemFilterMap[ifp] = ServiceItemFilter{
		Name:       string(ifp),
		Value:      []string{value},
		ParamName:  paramName,
		ParamValue: paramValue,
	}
}

// WithItemFilterTopRatedSellerOnly adds TopRatedSellerOnly ItemFilter
// The TopRatedSellerOnly item filter cannot be used together with either the Seller or ExcludeSeller item filters.
// The TopRatedSellerOnly item filter is supported for the following sites only:
// US (EBAY-US), Motors (EBAY-MOTOR), UK (EBAY-GB), IE (EBAY-IE), DE (EBAY-DE), AT (EBAY-AT), and CH (EBAY-CH).
func (sr *RequestItemFilter) WithItemFilterTopRatedSellerOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterTopRatedSellerOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterAuthorizedSellerOnly adds AuthorizedSellerOnly ItemFilter
// If set to true, returns only items listed by authorized sellers
func (sr *RequestItemFilter) WithItemFilterAuthorizedSellerOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterAuthorizedSellerOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterBestOfferOnly adds BestOfferOnly ItemFilter
// If true, the search results are limited to only items that have Best Offer enabled.
// Default is false.
func (sr *RequestItemFilter) WithItemFilterBestOfferOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterBestOfferOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterCharityOnly adds CharityOnly ItemFilter
// If true, the search results are limited to items for which all or part of the proceeds are given to a charity.
// Each item in the search results will include the ID of the given charity.
// Default is false.
func (sr *RequestItemFilter) WithItemFilterCharityOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterCharityOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterExcludeAutoPay adds ExcludeAutoPay ItemFilter
// If true, excludes all items requiring immediate payment. Default is false.
func (sr *RequestItemFilter) WithItemFilterExcludeAutoPay(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterExcludeAutoPay, strconv.FormatBool(b))
	return sr
}

// WithItemFilterFeaturedOnly adds FeaturedOnly ItemFilter
// If true, the search results are limited to featured item listings only. Default is false.
func (sr *RequestItemFilter) WithItemFilterFeaturedOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterFeaturedOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterFreeShippingOnly adds FreeShippingOnly ItemFilter
// If true, the search results are limited to only items with free shipping to the site specified in the req
// (see Global ID Values).
// Default is false.
func (sr *RequestItemFilter) WithItemFilterFreeShippingOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterFreeShippingOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterGetItFastOnly adds GetItFastOnly ItemFilter
// If true, the search results are limited to only Get It Fast listings. Default is false.
func (sr *RequestItemFilter) WithItemFilterGetItFastOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterGetItFastOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterHideDuplicateItems adds HideDuplicateItems ItemFilter
// If true, and there are duplicate items for an item in the search results,
// the subsequent duplicates will not appear in the results. Default is false.
func (sr *RequestItemFilter) WithItemFilterHideDuplicateItems(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterHideDuplicateItems, strconv.FormatBool(b))
	return sr
}

// WithItemFilterLocalPickupOnly adds LocalPickupOnly ItemFilter
// 	If true, the search results are limited to only items which have local pickup available. Default is false.
func (sr *RequestItemFilter) WithItemFilterLocalPickupOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterLocalPickupOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterLocalSearchOnly adds LocalSearchOnly ItemFilter
// If true, the search results are limited to only matching items with the Local Inventory Listing Options (LILO).
// Must be used together with the MaxDistance item filter, and the req must also specify buyerPostalCode.
// Currently, this is only available for the Motors site (global ID EBAY- MOTOR).
func (sr *RequestItemFilter) WithItemFilterLocalSearchOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterLocalSearchOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterLotsOnly adds LotsOnly ItemFilter
// If true, the search results are limited to only matching listings for which the lot size is 2 or more.
// Default is false.
func (sr *RequestItemFilter) WithItemFilterLotsOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterLotsOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterOutletSellerOnly adds OutletSellerOnly ItemFilter
// If set to true, returns only items listed by outlet sellers.
func (sr *RequestItemFilter) WithItemFilterOutletSellerOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterOutletSellerOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterReturnsAcceptedOnly adds ReturnsAcceptedOnly ItemFilter
// If set to true, returns only items where the seller accepts returns.
func (sr *RequestItemFilter) WithItemFilterReturnsAcceptedOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterReturnsAcceptedOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterSoldItemsOnly adds SoldItemsOnly ItemFilter
// Reserved for future use. If true, excludes all completed items which are not ended by being sold.
func (sr *RequestItemFilter) WithItemFilterSoldItemsOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterSoldItemsOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterWorldOfGoodOnly adds WorldOfGoodOnly ItemFilter
// If true, the search results are limited to only items listed in the World of Good marketplace.
// Defaults to false.
func (sr *RequestItemFilter) WithItemFilterWorldOfGoodOnly(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterWorldOfGoodOnly, strconv.FormatBool(b))
	return sr
}

// WithItemFilterAvailableTo adds AvailableTo ItemFilter
// Limits items to those available to the specified country only.
// Item filter LocatedIn cannot be used together with item filter AvailableTo.
// Expects the two-letter ISO 3166 country code to indicate the country where the item is located.
// For English names that correspond to each code (e.g., KY="Cayman Islands")
func (sr *RequestItemFilter) WithItemFilterAvailableTo(code string) *RequestItemFilter {
	sr.updateIFValue(ItemFilterAvailableTo, code)
	return sr
}

// WithItemFilterCurrency adds Currency ItemFilter
// Limits results to items listed with the specified currency only.
func (sr *RequestItemFilter) WithItemFilterCurrency(currency ItemFilterCurrencyIDOption) *RequestItemFilter {
	sr.updateIFValue(ItemFilterCurrency, string(currency))
	return sr
}

// WithItemFilterExpeditedShippingType adds ExpeditedShippingType ItemFilter
// Specifies the type of expedited shipping. You can specify either Expedited or OneDayShipping.
// Only items that can be shipped by the specified type are returned.
// ExpeditedShippingType is used together with the MaxHandlingTime and ReturnsAcceptedOnly filters
// to filter items for certain kinds of gifting events such as birthdays or holidays where the items must
// be delivered by a certain date.
func (sr *RequestItemFilter) WithItemFilterExpeditedShippingType(shipping ItemFilterExpeditedShippingTypeOption) *RequestItemFilter {
	sr.updateIFValue(ItemFilterExpeditedShippingType, string(shipping))
	return sr
}

// WithItemFilterFeedbackScoreMax adds FeedbackScoreMax ItemFilter
// Specifies the maximum feedback score of a seller whose items can be included in the response.
// If FeedbackScoreMin is also specified, the FeedbackScoreMax value must be greater than or equal to the
// FeedbackScoreMin value.
//
//  Values below 0 are ignored.
func (sr *RequestItemFilter) WithItemFilterFeedbackScoreMax(score int) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterFeedbackScoreMin(score int) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterListedIn(globalID GlobalID) *RequestItemFilter {
	sr.updateIFValue(ItemFilterListedIn, string(globalID))
	return sr
}

// WithItemFilterMaxBids adds MaxBids ItemFilter
// Limits the results to items with bid counts less than or equal to the specified value.
// If MinBids is also specified, the MaxBids value must be greater than or equal to the MinBids value.
//
//  Values below 0 are ignored.
func (sr *RequestItemFilter) WithItemFilterMaxBids(score int) *RequestItemFilter {
	if score < 0 {
		return sr
	}
	sr.updateIFValue(ItemFilterMaxBids, strconv.Itoa(score))
	return sr
}

// WithItemFilterMaxDistance adds MaxDistance ItemFilter
// Specifies the maximum distance from the specified postal code (buyerPostalCode) to search for items.
//  The req must also specify buyerPostalCode.
//  Values below 5 are ignored.
func (sr *RequestItemFilter) WithItemFilterMaxDistance(score int) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMaxHandlingTime(score int) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMaxQuantity(score int) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMinBids(score int) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMinQuantity(score int) *RequestItemFilter {
	if score < 1 {
		return sr
	}
	sr.updateIFValue(ItemFilterMinQuantity, strconv.Itoa(score))
	return sr
}

// WithItemFilterPaymentMethod adds PaymentMethod ItemFilter
// Limits results to items that accept the specified payment method.
func (sr *RequestItemFilter) WithItemFilterPaymentMethod(payment ItemFilterPaymentMethodOption) *RequestItemFilter {
	sr.updateIFValue(ItemFilterPaymentMethod, string(payment))
	return sr
}

// WithItemFilterSellerBusinessType adds SellerBusinessType ItemFilter
// Restricts the items to those that are from sellers whose
// business type is the specified value. Only one value can be specified.
func (sr *RequestItemFilter) WithItemFilterSellerBusinessType(businessType ItemFilterSellerBusinessTypeOption) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterValueBoxInventory(b bool) *RequestItemFilter {
	sr.updateIFValue(ItemFilterValueBoxInventory, strconv.FormatBool(b))
	return sr
}

// WithItemFilterSeller adds Seller ItemFilter
// The Seller item filter cannot be used together with either the ExcludeSeller or TopRatedSellerOnly item filters.
// Multiple values are allowed. Up to 100 sellers can be specified.
//  101th and next sellers will be skipped
func (sr *RequestItemFilter) WithItemFilterSeller(sellers ...string) *RequestItemFilter {
	sr.addIFValues(ItemFilterSeller, 100, sellers...)
	return sr
}

// WithItemFilterExcludeCategory adds ExcludeCategory ItemFilter
// Specify one or more category IDs. Search results will not include
// items from the specified categories or their child categories.
// Valid category IDs.
//  26th and next categories will be skipped
func (sr *RequestItemFilter) WithItemFilterExcludeCategory(categories ...string) *RequestItemFilter {
	sr.addIFValues(ItemFilterExcludeCategory, 25, categories...)
	return sr
}

// WithItemFilterExcludeSeller adds ExcludeSeller ItemFilter
// Specify one or more seller names. Search results will not include items from the specified sellers.
// The ExcludeSeller item filter cannot be used together with either the Seller or TopRatedSellerOnly item filters.
//  101th and next sellers will be skipped
func (sr *RequestItemFilter) WithItemFilterExcludeSeller(sellers ...string) *RequestItemFilter {
	sr.addIFValues(ItemFilterExcludeSeller, 100, sellers...)
	return sr
}

// WithItemFilterCondition adds Condition ItemFilter
// This item condition filter allows a user to filter items based on item condition.
func (sr *RequestItemFilter) WithItemFilterCondition(conditions ...ItemFilterConditionOption) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterListingType(listingTypes ...ItemFilterListingTypeOption) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterLocatedIn(codes ...string) *RequestItemFilter {
	sr.addIFValues(ItemFilterLocatedIn, 25, codes...)
	return sr
}

// WithItemFilterMaxPrice adds MaxPrice ItemFilter
// Specifies the maximum current price an item can have to be included in the response.
// If using with MinPrice to specify a price range, the MaxPrice value must be greater than or equal to MinPrice.
//  Values below 0 are ignored.
func (sr *RequestItemFilter) WithItemFilterMaxPrice(price float64) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMaxPriceWithCurrency(price float64, currency ItemFilterCurrencyIDOption) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMinPrice(price float64) *RequestItemFilter {
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
func (sr *RequestItemFilter) WithItemFilterMinPriceWithCurrency(price float64, currency ItemFilterCurrencyIDOption) *RequestItemFilter {
	if price < 0 {
		return sr
	}
	sr.updateIFValueWithParameter(ItemFilterMinPrice, "Currency", string(currency), strconv.FormatFloat(price, 'f', 2, 64))
	return sr
}

// WithItemFilterEndTimeFrom adds EndTimeFrom ItemFilter
// Limits the results to items ending on or after the specified time. Specify a time in the future.
func (sr *RequestItemFilter) WithItemFilterEndTimeFrom(datetime string) *RequestItemFilter {
	sr.updateIFValue(ItemFilterEndTimeFrom, datetime)
	return sr
}

// WithItemFilterEndTimeTo adds EndTimeTo ItemFilter
// Limits the results to items ending on or before the specified time. Specify a time in the future.
func (sr *RequestItemFilter) WithItemFilterEndTimeTo(datetime string) *RequestItemFilter {
	sr.updateIFValue(ItemFilterEndTimeTo, datetime)
	return sr
}

// WithItemFilterModTimeFrom adds ModTimeFrom ItemFilter
// Limits the results to active items whose status has changed since the specified time.
// Specify a time in the past. Time must be in GMT.
func (sr *RequestItemFilter) WithItemFilterModTimeFrom(datetime string) *RequestItemFilter {
	sr.updateIFValue(ItemFilterModTimeFrom, datetime)
	return sr
}

// WithItemFilterStartTimeFrom adds StartTimeFrom ItemFilter
// Limits the results to items started on or after the specified time. Specify a time in the future.
func (sr *RequestItemFilter) WithItemFilterStartTimeFrom(datetime string) *RequestItemFilter {
	sr.updateIFValue(ItemFilterStartTimeFrom, datetime)
	return sr
}

// WithItemFilterStartTimeTo adds StartTimeTo ItemFilter
// Limits the results to items started on or before the specified time. Specify a time in the future.
func (sr *RequestItemFilter) WithItemFilterStartTimeTo(datetime string) *RequestItemFilter {
	sr.updateIFValue(ItemFilterStartTimeTo, datetime)
	return sr
}
