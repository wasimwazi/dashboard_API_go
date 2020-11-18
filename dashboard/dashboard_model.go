package dashboard

//DataFromDB : Get data from the db
type DataFromDB struct {
	ProductID       string
	DistributorID   string
	Place           string
	ProductName     string
	DistributorName string
	Quantity        int
}

//TopNDataFromDB : Get data from the db
type TopNDataFromDB struct {
	ProductID       string
	DistributorID   string
	Place           string
	ProductName     string
	DistributorName string
	Quantity        int
	Rank            int
}

//Product : Product struct
type Product struct {
	ProductName string `json:"name"`
	Quantity    int    `json:"quantity"`
}

//Products :: Slice of product
type Products []Product

//Distributor : Distributor struct
type Distributor struct {
	DistributorName string   `json:"name"`
	Products        Products `json:"product"`
}

//AreaDistributionResponse : Response struct for area distribution
type AreaDistributionResponse struct {
	Name         string        `json:"name"`
	Distributors []Distributor `json:"distributor"`
}

//PlaceQuantity : struct for place response
type PlaceQuantity struct {
	PlaceName string `json:"name"`
	Quantity  int    `json:"quantity"`
}

//PlaceQuantities : Slice of places
type PlaceQuantities []PlaceQuantity

//DistributorPlace : Distributor struct
type DistributorPlace struct {
	DistributorName string          `json:"name"`
	Place           PlaceQuantities `json:"place"`
}

//ProductDistributionResponse : Response struct for product distribution
type ProductDistributionResponse struct {
	Name         string             `json:"name"`
	Distributors []DistributorPlace `json:"distributor"`
}

//Place : Place struct
type Place struct {
	Name    string   `json:"name"`
	Product Products `json:"product"`
}

//DistributorPerformanceResponse : Response struct for distributor performance
type DistributorPerformanceResponse struct {
	DistributorName string  `json:"name"`
	Place           []Place `json:"place"`
}
