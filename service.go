package finding

import (
	"github.com/go-resty/resty/v2"
	"time"
)

// Service represents Ebay Finding API service
type Service struct {
	version         string
	endpoint        string
	globalID        string
	securityAppName string
	timeout         time.Duration
	pageLimit       int
}

// NewService creates new Ebay Finding API service
// API version according to EbayFindingAPIVersion (1.13.0)
// Default endpoint: EbayEndpointProduction (https://svcs.ebay.com/services/search/FindingService/v1)
// Default GlobalID: GlobalIDEbayUS (EBAY-US)
// Default Page Limit: DefaultItemsPerPage (100)
// Default timeout for requests: 10 seconds
func NewService(securityAppName string) *Service {
	s := &Service{
		version:         EbayFindingAPIVersion,
		securityAppName: securityAppName,
		timeout:         10 * time.Second,
	}
	s.WithEndpoint(EbayEndpointProduction)
	s.WithGlobalID(GlobalIDEbayUS)
	s.WithPageLimit(DefaultItemsPerPage)
	return s
}

// WithEndpoint changes endpoint for service
// You can add your own endpoint (for tests purposes)
func (s *Service) WithEndpoint(endpoint string) *Service {
	s.endpoint = endpoint
	return s
}

// WithGlobalID changes site for search
func (s *Service) WithGlobalID(globalID GlobalID) *Service {
	s.globalID = string(globalID)
	return s
}

// WithTimeout changes default timeout for search requests
func (s *Service) WithTimeout(timeout time.Duration) *Service {
	s.timeout = timeout
	return s
}

// WithPageLimit changes limit of items for any child req
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

// NewAdvancedRequest creates new search req
func (s *Service) NewAdvancedRequest() *AdvancedRequest {
	sr := AdvancedRequest{}
	sr.Initialize()
	sr.Client = resty.New().
		SetHeader("X-EBAY-SOA-SERVICE-VERSION", s.version).
		SetHeader("X-EBAY-SOA-OPERATION-NAME", string(OperationFindItemsAdvanced)).
		SetHeader("X-EBAY-SOA-SECURITY-APPNAME", s.securityAppName).
		SetHeader("X-EBAY-SOA-REQUEST-DATA-FORMAT", EbayRequestDataFormat).
		SetHeader("X-EBAY-SOA-RESPONSE-DATA-FORMAT", EbayResponseDataFormat).
		SetHeader("X-EBAY-SOA-GLOBAL-ID", s.globalID).
		SetTimeout(s.timeout)
	sr.URL = s.endpoint
	sr.WithPageLimit(s.pageLimit)
	return &sr
}
