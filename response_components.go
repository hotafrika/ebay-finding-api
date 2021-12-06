package finding

type responseStandard struct {
	Ack string `xml:"ack"`
	//ErrorMessage []Error `xml:"errorMessage"`
	Timestamp string `xml:"timestamp"`
	Version   string `xml:"version"`
}

type Error struct {
	Category    string      `json:"category"`
	Domain      string      `json:"domain"`
	ErrorID     string      `json:"errorId"`
	ExceptionID string      `json:"exceptionId"`
	Message     string      `json:"message"`
	Parameter   []Parameter `json:"parameter"`
	Severity    string      `json:"severity"`
	Subdomain   string      `json:"subdomain"`
}

type Parameter struct {
	Name  string `json:"@name"`
	Value string `json:"__value__"`
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
	PaginationOutput PaginationOutput `json:"paginationOutput"`
}

type PaginationOutput struct {
	PageNumber     int `json:"pageNumber"`
	EntriesPerPage int `json:"entriesPerPage"`
	TotalPages     int `json:"totalPages"`
	TotalEntries   int `json:"totalEntries"`
}

/*
=====================================================
*/

type ResponseSearchResult struct {
	SearchResult SearchResult `json:"searchResult"`
}

type SearchResult struct {
	// TODO
}
