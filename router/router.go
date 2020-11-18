package router

import (
	"dashboard-api/dashboard"
	"dashboard-api/user"
	"database/sql"

	"github.com/go-chi/chi"
)

// Router : Basic Router
type Router interface {
	Setup() *chi.Mux
}

// ChiRouter : Router that holds DB connection
type ChiRouter struct {
	DB *sql.DB
}

// NewRouter : Returns Basic Router
func NewRouter(db *sql.DB) Router {
	return &ChiRouter{
		DB: db,
	}
}

// Setup : chi Router
func (r *ChiRouter) Setup() *chi.Mux {
	cr := chi.NewRouter()
	dashboardHandler := dashboard.NewHTTPHandler(r.DB)
	userHandler := user.NewHTTPHandler(r.DB)
	cr.Post("/login", userHandler.LoginUser)
	cr.Group(func(cr chi.Router) {
		cr.Use(user.VerifyToken)
		cr.Get("/getAreaDistribution", dashboardHandler.GetAreaDistribution)
		cr.Get("/getProductDistribution", dashboardHandler.GetProductDistribution)
		cr.Get("/getDistributorPerformance", dashboardHandler.GetDistributorPerformance)
		cr.Get("/getTopN", dashboardHandler.GetTopN)
	})
	return cr
}
