package utils

const (
	//InternalServerError : Error 500
	InternalServerError = 500

	//Success : Success 200
	Success = 200

	//BadRequest : Error 400
	BadRequest = 400

	//ProductError : Error when distributor is empty
	ProductError = "'product' parameter empty"

	//PlaceError : Error when distributor is empty
	PlaceError = "'place' parameter emtpy"

	//DistributorError : Error when distributor is empty
	DistributorError = "'distributor' parameter emtpy"

	//NoRecordError : Error when no records found
	NoRecordError = "No records found"

	//DecodeError : error while decoding
	DecodeError = "Decode error, please check the input"

	//PasswordError : Error password
	PasswordError = "Password is empty"

	//UsernameError : Username error
	UsernameError = "Username is empty"

	//InvalidUsernameOrPasswordError : username or password error
	InvalidUsernameOrPasswordError = "Invalid username or password"

	//DatabaseError : Error in DB
	DatabaseError = "Database error"

	//JWTError : Error with JWT
	JWTError = "JWT Error"

	//AccessTokenInterval : Access token interval
	AccessTokenInterval = 100

	//InvalidTokenError : Invalid token error
	InvalidTokenError = "Invalid Token"

	//ForbiddenError : Forbidden http error code
	ForbiddenError = 403

	//TopNError : Error in number parameter
	TopNError = "'number' parameter empty"
)
