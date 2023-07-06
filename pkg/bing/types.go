package bing

type TimeResponse struct {
	ResourceSets []ResourceSet `json:"resourceSets"`
	StatusCode   int           `json:"statusCode"`
}

type ResourceSet struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	TimeZoneAtLocation []TimeZoneAtLocation `json:"timeZoneAtLocation"`
}

type TimeZoneAtLocation struct {
	PlaceName string     `json:"placeName"`
	TimeZone  []TimeZone `json:"timeZone"`
}

type TimeZone struct {
	GenericName       string        `json:"genericName"`
	Abbreviation      string        `json:"abbreviation"`
	IanaTimeZoneId    string        `json:"ianaTimeZoneId"`
	WindowsTimeZoneId string        `json:"windowsTimeZoneId"`
	UtcOffset         string        `json:"utcOffset"`
	ConvertedTime     ConvertedTime `json:"convertedTime"`
}

type ConvertedTime struct {
	LocalTime           string `json:"localTime"`
	UtcOffsetWithDst    string `json:"utcOffsetWithDst"`
	TimeZoneDisplayName string `json:"timeZoneDisplayName"`
	TimeZoneDisplayAbbr string `json:"timeZoneDisplayAbbr"`
}

type LocationResource struct {
	BoundingBox   []float64 `json:"bbox"`
	Name          string    `json:"name"`
	Point         Point     `json:"point"`
	Address       Address   `json:"address"`
	Confidence    string    `json:"confidence"`
	EntityType    string    `json:"entityType"`
	GeocodePoints []Geocode `json:"geocodePoints"`
	MatchCodes    []string  `json:"matchCodes"`
}

type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Address struct {
	AddressLine      string       `json:"addressLine"`
	AdminDistrict    string       `json:"adminDistrict"`
	AdminDistrict2   string       `json:"adminDistrict2"`
	CountryRegion    string       `json:"countryRegion"`
	FormattedAddress string       `json:"formattedAddress"`
	Intersection     Intersection `json:"intersection"`
	Locality         string       `json:"locality"`
	PostalCode       string       `json:"postalCode"`
}

type Intersection struct {
	BaseStreet       string `json:"baseStreet"`
	SecondaryStreet1 string `json:"secondaryStreet1"`
	SecondaryStreet2 string `json:"secondaryStreet2"`
	IntersectionType string `json:"intersectionType"`
	DisplayName      string `json:"displayName"`
}

type Geocode struct {
	Type              string    `json:"type"`
	Coordinates       []float64 `json:"coordinates"`
	CalculationMethod string    `json:"calculationMethod"`
	UsageTypes        []string  `json:"usageTypes"`
}

type LocationResponse struct {
	ResourceSets []LocationResourceSet `json:"resourceSets"`
	StatusCode   int                   `json:"statusCode"`
}

type LocationResourceSet struct {
	Resources []LocationResource `json:"resources"`
}
