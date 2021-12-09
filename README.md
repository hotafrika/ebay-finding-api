![example workflow](https://github.com/hotafrika/ebay-finding-api/actions/workflows/autotests.yml/badge.svg)

Golang library implementation of eBay [Finding API](https://developer.ebay.com/Devzone/finding/Concepts/FindingAPIGuide.html)

You can check all the Finding API search options and limitations [here](https://developer.ebay.com/Devzone/finding/CallRef/index.html)



### Example
```go
package main

import (
	"fmt"
	"github.com/hotafrika/ebay-finding-api"
)

func main() {
	// Create Finding API service
	s := finding.NewService("your-sec-app-name").WithPageLimit(50)
	// Create findItemsAdvanced call request
	r := s.NewAdvancedRequest()
	// Set few search parameters
	r.WithDescriptionSearch(true)
	r.WithKeywords("harry potter")
	r.WithItemFilterCondition(finding.ConditionGood, finding.ConditionNew)
	r.WithItemFilterMaxPriceWithCurrency(100, finding.CurrencyIDUSD)
	r.WithSortOrder(finding.SortOrderCurrentPriceHighest)
	r.WithPageLimit(1)
	r.WithOutputSelectors(finding.OutputSelectorAspectHistogram,
		finding.OutputSelectorSellerInfo,
		finding.OutputSelectorStoreInfo,
		finding.OutputSelectorUnitPriceInfo,
		finding.OutputSelectorGalleryInfo,
		finding.OutputSelectorPictureURLSuperSize,
		finding.OutputSelectorConditionHistogram,
		finding.OutputSelectorCategoryHistogram,
		finding.OutputSelectorPictureURLLarge)
	// Get first page
	res, err := r.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
	// Get second page
	res2, err := r.GetPage(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res2)
}

```

