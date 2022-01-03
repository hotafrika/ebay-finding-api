package finding

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"sort"
	"testing"
)

func TestAdvancedRequest_GetBody(t *testing.T) {
	service := NewService("")

	type tcase struct {
		filename string
		request  *AdvancedRequest
	}

	tests := make([]tcase, 0)

	request1 := service.NewAdvancedRequest()
	request1.WithKeywords("tolkien")
	request1.WithPageLimit(2)
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	request2 := service.NewAdvancedRequest()
	request2.WithKeywords("iPad")
	request2.WithPageLimit(2)
	request2.WithItemFilterMaxHandlingTime(1)
	request2.WithItemFilterExpeditedShippingType(ExpeditedShippingTypeOneDayShipping)
	tests = append(tests, tcase{
		filename: "BasicFastShipping.xml",
		request:  request2,
	})

	request3 := service.NewAdvancedRequest()
	request3.WithKeywords("tolkien")
	request3.WithPageLimit(2)
	request3.WithItemFilterFreeShippingOnly(true)
	tests = append(tests, tcase{
		filename: "BasicFreeShipping.xml",
		request:  request3,
	})

	request4 := service.NewAdvancedRequest()
	request4.WithKeywords("iPad")
	request4.WithPageLimit(3)
	request4.WithItemFilterPaymentMethod(PaymentMethodPayPal)
	request4.WithItemFilterHideDuplicateItems(true)
	tests = append(tests, tcase{
		filename: "BasicFilteredPaymentMethod.xml",
		request:  request4,
	})

	request5 := service.NewAdvancedRequest()
	request5.WithKeywords("tolkien")
	request5.WithPageLimit(2)
	request5.WithItemFilterFeedbackScoreMin(200)
	tests = append(tests, tcase{
		filename: "BasicFilteredSellerFeedbackScore.xml",
		request:  request5,
	})

	request6 := service.NewAdvancedRequest()
	request6.WithKeywords("tolkien")
	request6.WithPageLimit(2)
	request6.WithDescriptionSearch(true)
	request6.WithCategoryIDInt(1)
	tests = append(tests, tcase{
		filename: "CategoryKeywordWDescriptionSearch.xml",
		request:  request6,
	})

	request7 := service.NewAdvancedRequest()
	request7.WithPageLimit(2)
	request7.WithCategoryIDInt(31388)
	request7.WithOutputSelectors(OutputSelectorAspectHistogram)
	request7.WithItemFilterConditionName(ConditionNameUsed)
	request7.WithItemFilterListingType(ListingTypeAuctionWithBIN)
	tests = append(tests, tcase{
		filename: "RetrievingAspectHistogramInformation.xml",
		request:  request7,
	})

	request8 := service.NewAdvancedRequest()
	request8.WithPageLimit(3)
	request8.WithOutputSelectors(OutputSelectorSellerInfo)
	request8.WithItemFilterSeller("e***y")
	tests = append(tests, tcase{
		filename: "RetrieveAllActiveListingsForSeller.xml",
		request:  request8,
	})

	request9 := service.NewAdvancedRequest()
	request9.WithPageLimit(2)
	request9.WithCategoryIDInt(31388)
	request9.WithItemFilterMaxPriceWithCurrency(75, CurrencyIDUSD)
	request9.WithItemFilterMinPriceWithCurrency(50, CurrencyIDUSD)
	request9.WithAspectFilter("Megapixels", "5.0 to 5.9 MP")
	tests = append(tests, tcase{
		filename: "RefiningResultsAspectFilters.xml",
		request:  request9,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 AdvancedRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "advanced", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 AdvancedRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			sort.Slice(req1.AspectFilter, func(i, j int) bool { return req1.AspectFilter[i].AspectName > req1.AspectFilter[j].AspectName })
			sort.Slice(req2.AspectFilter, func(i, j int) bool { return req2.AspectFilter[i].AspectName > req2.AspectFilter[j].AspectName })

			sort.Slice(req1.ItemFilter, func(i, j int) bool { return req1.ItemFilter[i].Name > req1.ItemFilter[j].Name })
			sort.Slice(req2.ItemFilter, func(i, j int) bool { return req2.ItemFilter[i].Name > req2.ItemFilter[j].Name })

			sort.Slice(req1.OutputSelector, func(i, j int) bool { return req1.OutputSelector[i] > req1.OutputSelector[j] })
			sort.Slice(req2.OutputSelector, func(i, j int) bool { return req2.OutputSelector[i] > req2.OutputSelector[j] })

			sort.Slice(req1.CategoryID, func(i, j int) bool { return req1.CategoryID[i] > req1.CategoryID[j] })
			sort.Slice(req2.CategoryID, func(i, j int) bool { return req2.CategoryID[i] > req2.CategoryID[j] })

			assert.EqualValues(t, req1.AspectFilter, req2.AspectFilter)
			assert.EqualValues(t, req1.ItemFilter, req2.ItemFilter)
			assert.EqualValues(t, req1.OutputSelector, req2.OutputSelector)
			assert.EqualValues(t, req1.CategoryID, req2.CategoryID)
			assert.Equal(t, req1.Keywords, req2.Keywords)
			assert.Equal(t, req1.DescriptionSearch, req2.DescriptionSearch)
			assert.Equal(t, req1.RequestStandard, req2.RequestStandard)
		})
	}
}

func TestByCategoryRequest_GetBody(t *testing.T) {
	service := NewService("")

	type tcase struct {
		filename string
		request  *ByCategoryRequest
	}

	tests := make([]tcase, 0)

	request1 := service.NewByCategoryRequest()
	request1.WithCategoryIDInt(10181)
	request1.WithPageLimit(1)
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	request2 := service.NewByCategoryRequest()
	request2.WithCategoryIDInt(10181)
	request2.WithPageLimit(2)
	request2.WithOutputSelectors(OutputSelectorCategoryHistogram)
	tests = append(tests, tcase{
		filename: "RetrievingCategoryHistogram.xml",
		request:  request2,
	})

	request3 := service.NewByCategoryRequest()
	request3.WithCategoryIDInt(307)
	request3.WithPageLimit(2)
	request3.WithOutputSelectors(OutputSelectorAspectHistogram)
	tests = append(tests, tcase{
		filename: "ReturningAspectHistogramInformation.xml",
		request:  request3,
	})

	request4 := service.NewByCategoryRequest()
	request4.WithCategoryIDInt(307)
	request4.WithPageLimit(3)
	request4.WithPageNumber(11)
	request4.WithAspectFilter("Genre", "Rock")
	request4.WithAspectFilter("Condition", "Brand New", "Like New")
	tests = append(tests, tcase{
		filename: "RefiningResultsUsingAspectFilters.xml",
		request:  request4,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			fmt.Println(string(got))
			if !assert.NoError(t, err) {
				return
			}

			var req1 ByCategoryRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "bycategory", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 ByCategoryRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			sort.Slice(req1.AspectFilter, func(i, j int) bool { return req1.AspectFilter[i].AspectName > req1.AspectFilter[j].AspectName })
			sort.Slice(req2.AspectFilter, func(i, j int) bool { return req2.AspectFilter[i].AspectName > req2.AspectFilter[j].AspectName })

			sort.Slice(req1.ItemFilter, func(i, j int) bool { return req1.ItemFilter[i].Name > req1.ItemFilter[j].Name })
			sort.Slice(req2.ItemFilter, func(i, j int) bool { return req2.ItemFilter[i].Name > req2.ItemFilter[j].Name })

			sort.Slice(req1.OutputSelector, func(i, j int) bool { return req1.OutputSelector[i] > req1.OutputSelector[j] })
			sort.Slice(req2.OutputSelector, func(i, j int) bool { return req2.OutputSelector[i] > req2.OutputSelector[j] })

			sort.Slice(req1.CategoryID, func(i, j int) bool { return req1.CategoryID[i] > req1.CategoryID[j] })
			sort.Slice(req2.CategoryID, func(i, j int) bool { return req2.CategoryID[i] > req2.CategoryID[j] })

			assert.EqualValues(t, req1.AspectFilter, req2.AspectFilter)
			assert.EqualValues(t, req1.ItemFilter, req2.ItemFilter)
			assert.EqualValues(t, req1.OutputSelector, req2.OutputSelector)
			assert.EqualValues(t, req1.CategoryID, req2.CategoryID)
			assert.Equal(t, req1.RequestStandard, req2.RequestStandard)
		})
	}
}

func TestByKeywordsRequest_GetBody(t *testing.T) {
	service := NewService("")

	type tcase struct {
		filename string
		request  *ByKeywordsRequest
	}

	tests := make([]tcase, 0)

	request1 := service.NewByKeywordsRequest()
	request1.WithKeywords("harry potter phoenix")
	request1.WithPageLimit(2)
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	request2 := service.NewByKeywordsRequest()
	request2.WithKeywords("harry potter phoenix")
	request2.WithPageLimit(2)
	request2.WithItemFilterMaxPriceWithCurrency(10, CurrencyIDUSD)
	tests = append(tests, tcase{
		filename: "MaxPriceItemFilter.xml",
		request:  request2,
	})

	request3 := service.NewByKeywordsRequest()
	request3.WithKeywords("harry potter phoenix")
	request3.WithPageLimit(2)
	request3.WithItemFilterMaxPriceWithCurrency(10, CurrencyIDUSD)
	request3.WithItemFilterFreeShippingOnly(true)
	tests = append(tests, tcase{
		filename: "FilterPriceFreeShipping.xml",
		request:  request3,
	})

	request4 := service.NewByKeywordsRequest()
	request4.WithKeywords("bagpipes")
	request4.WithBuyerPostalCode("9***5")
	request4.WithItemFilterMaxDistance(25)
	tests = append(tests, tcase{
		filename: "ProximitySearch.xml",
		request:  request4,
	})

	request5 := service.NewByKeywordsRequest()
	request5.WithKeywords("pre-CBS")
	request5.WithBuyerPostalCode("9***0")
	request5.WithSortOrder(SortOrderDistanceNearest)
	request5.WithOutputSelectors(OutputSelectorSellerInfo)
	request5.WithItemFilterMaxPriceWithCurrency(2500, CurrencyIDUSD)
	request5.WithItemFilterMinPriceWithCurrency(1000, CurrencyIDUSD)
	tests = append(tests, tcase{
		filename: "ReviewingSellerInformationItemsSortedDistance.xml",
		request:  request5,
	})

	request6 := service.NewByKeywordsRequest()
	request6.WithKeywords("camalots")
	request6.WithPageLimit(2)
	request6.WithSortOrder(SortOrderEndTimeSoonest)
	request6.WithAffiliate("9", "1********0", "k***n")
	tests = append(tests, tcase{
		filename: "SpecifyingAffiliateInformation.xml",
		request:  request6,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			//fmt.Println(string(got))
			if !assert.NoError(t, err) {
				return
			}

			var req1 ByKeywordsRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "bykeywords", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 ByKeywordsRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			sort.Slice(req1.AspectFilter, func(i, j int) bool { return req1.AspectFilter[i].AspectName > req1.AspectFilter[j].AspectName })
			sort.Slice(req2.AspectFilter, func(i, j int) bool { return req2.AspectFilter[i].AspectName > req2.AspectFilter[j].AspectName })

			sort.Slice(req1.ItemFilter, func(i, j int) bool { return req1.ItemFilter[i].Name > req1.ItemFilter[j].Name })
			sort.Slice(req2.ItemFilter, func(i, j int) bool { return req2.ItemFilter[i].Name > req2.ItemFilter[j].Name })

			sort.Slice(req1.OutputSelector, func(i, j int) bool { return req1.OutputSelector[i] > req1.OutputSelector[j] })
			sort.Slice(req2.OutputSelector, func(i, j int) bool { return req2.OutputSelector[i] > req2.OutputSelector[j] })

			assert.EqualValues(t, req1.AspectFilter, req2.AspectFilter)
			assert.EqualValues(t, req1.ItemFilter, req2.ItemFilter)
			assert.EqualValues(t, req1.OutputSelector, req2.OutputSelector)
			assert.Equal(t, req1.RequestStandard, req2.RequestStandard)
			assert.Equal(t, req1.Keywords, req2.Keywords)
			assert.Equal(t, req1.DescriptionSearch, req2.DescriptionSearch)
		})
	}
}
