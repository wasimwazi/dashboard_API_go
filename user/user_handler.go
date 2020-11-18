package user

import (
	"dashboard-api/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

//HandlerInterface : User handler interface
type HandlerInterface interface {
	LoginUser(w http.ResponseWriter, r *http.Request)
}

//Handler : Uer Handler
type Handler struct {
	us ServiceInterface
}

// NewHTTPHandler : Returns Admin HTTP Handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		us: NewService(db),
	}
}

// LoginUser : to Login User
func (userHandler *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /login API hit!")
	var login Login
	body := json.NewDecoder(r.Body)
	err := body.Decode(&login)
	if err != nil {
		log.Println("Error :", utils.DecodeError)
		utils.Fail(w, utils.BadRequest, utils.DecodeError)
		return
	}
	//Validate the login request
	err = ValidateLogin(login)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.BadRequest, err.Error())
		return
	}
	userLogged, err := userHandler.us.LoginUser(&login)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Error : Invalid username or password!")
			utils.Fail(w, utils.ForbiddenError, err.Error())
			return
		}
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.InternalServerError, err.Error())
		return
	}
	jwt, err := GenerateToken(userLogged.ID, userLogged.Username)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.ForbiddenError, err.Error())
		return
	}
	log.Println("App : User logged in successfully!")
	utils.Send(w, utils.Success, jwt)
	return
}
