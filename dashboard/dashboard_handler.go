package dashboard

import (
	"dashboard-api/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

//HandlerInterface : Interface to handle the requests
type HandlerInterface interface {
	GetAreaDistribution(w http.ResponseWriter, r *http.Request)
	GetProductDistribution(w http.ResponseWriter, r *http.Request)
	GetDistributorPerformance(w http.ResponseWriter, r *http.Request)
	GetTopN(w http.ResponseWriter, r *http.Request)
}

//Handler : Story Handler Struct
type Handler struct {
	dashboardService ServiceInterface
}

//NewHTTPHandler : Returns Story HTTP Handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		dashboardService: NewService(db),
	}
}

//GetAreaDistribution : function to get the area distribution map
func (dh *Handler) GetAreaDistribution(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /getAreaDistribution API hit!")
	//Get place parameter from request
	place := r.URL.Query().Get("place")
	//Check if the request parameter is empty or not
	if len(place) <= 0 {
		log.Println("Error in GetAreaDistribution() :", utils.PlaceError)
		utils.Fail(w, utils.BadRequest, utils.PlaceError)
		return
	}
	areaDistribution, status, err := dh.dashboardService.GetAreaDistribution(place)
	if err != nil {
		log.Println("Error in GetAreaDistribution() :", err.Error())
		utils.Fail(w, status, err.Error())
		return
	}
	utils.Send(w, status, areaDistribution)
	return
}

//GetProductDistribution : function to get the product distribution map
func (dh *Handler) GetProductDistribution(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /getProductDistribution API hit!")
	//Get product parameter from request
	productName := r.URL.Query().Get("product")
	//Check if the request parameter is empty or not
	if len(productName) <= 0 {
		log.Println("Error in GetProductDistribution() :", utils.ProductError)
		utils.Fail(w, utils.BadRequest, utils.ProductError)
		return
	}
	productDistribution, status, err := dh.dashboardService.GetProductDistribution(productName)
	if err != nil {
		log.Println("Error in GetProductDistribution() :", err.Error())
		utils.Fail(w, status, err.Error())
		return
	}
	utils.Send(w, status, productDistribution)
	return
}

//GetDistributorPerformance : function to get the distributor performance graph
func (dh *Handler) GetDistributorPerformance(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /getDistributorPerformance API hit!")
	//Get distributor parameter from request
	distributorName := r.URL.Query().Get("distributor")
	//Check if the request parameter is empty or not
	if len(distributorName) <= 0 {
		log.Println("Error in GetDistributorPerformance() :", utils.DistributorError)
		utils.Fail(w, utils.BadRequest, utils.DistributorError)
		return
	}
	distributorPerformance, status, err := dh.dashboardService.GetDistributorPerformance(distributorName)
	if err != nil {
		log.Println("Error in GetDistributorPerformance() :", err.Error())
		utils.Fail(w, status, err.Error())
		return
	}
	utils.Send(w, status, distributorPerformance)
	return
}

//GetTopN : Function to get top N for given distributor
func (dh *Handler) GetTopN(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /getTopN API hit")
	//Get number parameter from request
	n := r.URL.Query().Get("number")
	//Check if the request parameter is empty or not
	var number int
	var err error
	if n != "" {
		number, err = strconv.Atoi(n)
		if err != nil {
			log.Println("Error in GetTopN() :", err.Error())
			utils.Fail(w, utils.BadRequest, err.Error())
			return
		}
	} else {
		log.Println("Error in GetTopN() :", utils.TopNError)
		utils.Fail(w, utils.BadRequest, utils.TopNError)
		return
	}
	distributorName := r.URL.Query().Get("distributor")
	if len(distributorName) <= 0 {
		log.Println("Error in GetDistributorPerformance() :", utils.DistributorError)
		utils.Fail(w, utils.BadRequest, utils.DistributorError)
		return
	}
	topN, status, err := dh.dashboardService.GetTopN(number, distributorName)
	if err != nil {
		log.Println("Error in GetTopN() :", err.Error())
		utils.Fail(w, status, err.Error())
		return
	}
	utils.Send(w, status, topN)
	return
}
