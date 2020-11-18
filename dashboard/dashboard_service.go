package dashboard

import (
	"dashboard-api/utils"
	"database/sql"
	"errors"
)

//ServiceInterface : Dashboard service interface
type ServiceInterface interface {
	GetAreaDistribution(string) (*AreaDistributionResponse, int, error)
	GetProductDistribution(string) (*ProductDistributionResponse, int, error)
	GetDistributorPerformance(string) (*DistributorPerformanceResponse, int, error)
	GetTopN(int, string) (*DistributorPerformanceResponse, int, error)
}

//Service : Dashboard service struct
type Service struct {
	dashboardRepo Repo
}

//NewService : Returns dashboard service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		dashboardRepo: NewRepo(db),
	}
}

//GetAreaDistribution : Get area distribution service
func (service *Service) GetAreaDistribution(place string) (*AreaDistributionResponse, int, error) {
	areaDistribution, err := service.dashboardRepo.GetDataFromDB(place, "area")
	if err != nil {
		return nil, utils.InternalServerError, err
	}
	if len(areaDistribution) <= 0 {
		return nil, utils.Success, errors.New(utils.NoRecordError)
	}
	var response AreaDistributionResponse
	response.Name = areaDistribution[0].Place
	distributorMap := make(map[string]Products)
	//Formating the data for response
	//Creating distributor map
	for _, v := range areaDistribution {
		var product Product
		product.ProductName = v.ProductName
		product.Quantity = v.Quantity
		distributorMap[v.DistributorName] = append(distributorMap[v.DistributorName], product)
	}
	var distributor Distributor
	var distributors []Distributor
	//Converting the map to array of objects
	for i, v := range distributorMap {
		distributor.DistributorName = i
		distributor.Products = v
		distributors = append(distributors, distributor)
	}
	response.Distributors = distributors
	return &response, utils.Success, nil
}

//GetProductDistribution : Get product distribution service
func (service *Service) GetProductDistribution(productName string) (*ProductDistributionResponse, int, error) {
	productDistribution, err := service.dashboardRepo.GetDataFromDB(productName, "product")
	if err != nil {
		return nil, utils.InternalServerError, err
	}
	if len(productDistribution) <= 0 {
		return nil, utils.Success, errors.New(utils.NoRecordError)
	}
	var response ProductDistributionResponse
	response.Name = productDistribution[0].ProductName
	distributorMap := make(map[string]PlaceQuantities)
	//Formating the data for response
	//Creating distributor map
	for _, v := range productDistribution {
		var place PlaceQuantity
		place.PlaceName = v.Place
		place.Quantity = v.Quantity
		distributorMap[v.DistributorName] = append(distributorMap[v.DistributorName], place)
	}
	var distributor DistributorPlace
	var distributors []DistributorPlace
	//Converting the map to array of objects 
	for i, v := range distributorMap {
		distributor.DistributorName = i
		distributor.Place = v
		distributors = append(distributors, distributor)
	}
	response.Distributors = distributors
	return &response, utils.Success, nil
}

//GetDistributorPerformance : Get distributor performance service
func (service *Service) GetDistributorPerformance(distributorName string) (*DistributorPerformanceResponse, int, error) {
	distributorPerformance, err := service.dashboardRepo.GetDataFromDB(distributorName, "distributor")
	if err != nil {
		return nil, utils.InternalServerError, err
	}
	if len(distributorPerformance) <= 0 {
		return nil, utils.Success, errors.New(utils.NoRecordError)
	}
	var response DistributorPerformanceResponse
	response.DistributorName = distributorPerformance[0].DistributorName
	placeMap := make(map[string]Products)
	//Formating the data for response
	//Creating distributor map
	for _, v := range distributorPerformance {
		var product Product
		product.ProductName = v.ProductName
		product.Quantity = v.Quantity
		placeMap[v.Place] = append(placeMap[v.Place], product)
	}
	var places []Place
	//Converting the map to array of objects
	for i, v := range placeMap {
		var place Place
		place.Name = i
		place.Product = v
		places = append(places, place)
	}
	response.Place = places
	return &response, utils.Success, nil
}

//GetTopN : get top N service
func (service *Service) GetTopN(n int, distributorName string) (*DistributorPerformanceResponse, int, error) {
	topNFromDB, err := service.dashboardRepo.GetTopN(n, distributorName)
	if err != nil {
		return nil, utils.InternalServerError, err
	}
	if len(topNFromDB) <= 0 {
		return nil, utils.Success, errors.New(utils.NoRecordError)
	}
	var response DistributorPerformanceResponse
	response.DistributorName = topNFromDB[0].DistributorName
	placeMap := make(map[string]Products)
	//Formating the data for response
	//Creating distributor map
	for _, v := range topNFromDB {
		var product Product
		product.ProductName = v.ProductName
		product.Quantity = v.Quantity
		placeMap[v.Place] = append(placeMap[v.Place], product)
	}
	var places []Place
	//Converting the map to array of objects
	for i, v := range placeMap {
		var place Place
		place.Name = i
		place.Product = v
		places = append(places, place)
	}
	response.Place = places
	return &response, utils.Success, nil
}