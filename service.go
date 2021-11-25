package ebay_finding_api

type Service struct {
	version   string
	endpoint  string
	operation string
	globalID  string
	pageLimit int
}

func NewService(operation EbayOperation) Service {
	s := Service{
		version:   EbayFindingAPIVersion,
		operation: string(operation),
	}
	s.WithEndpoint(EbayEndpointProduction)
	s.WithGlobalID(GlobalIDEbayUS)
	s.WithPageLimit(DefaultItemsPerPage)
	return s
}

func (s *Service) WithEndpoint(endpoint string) *Service {
	s.endpoint = endpoint
	return s
}

func (s *Service) WithOperation(operation EbayOperation) *Service {
	s.operation = string(operation)
	return s
}

func (s *Service) WithGlobalID(globalID GlobalID) *Service {
	s.globalID = string(globalID)
	return s
}

func (s *Service) WithPageLimit(limit int) *Service {
	if limit < 1 {
		limit = 1
	} else {
		if limit > 100 {
			limit = 100
		}
	}

	s.pageLimit = limit
	return s
}

func (s *Service) NewRequest() ServiceRequest {
	return ServiceRequest{
		itemFilterMap: make(map[ItemFilterParameter]ServiceItemFilter),
		PaginationInput: ServicePaginationInput{
			EntriesPerPage: s.pageLimit,
		},
	}
}
