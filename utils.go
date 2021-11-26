package ebay_finding_api

import "time"

func fromEbayDateTime(ebayDT string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", ebayDT)
}

func toEbayDateTime(datetime time.Time) string {
	return datetime.Format("2006-01-02T15:04:05.000Z")
}
