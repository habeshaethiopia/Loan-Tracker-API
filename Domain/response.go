package domain
//write error and success response struct
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
//write error and success response struct
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}
// login resoponse
type LoginResponse struct {
	AccessToken string
	RefreshToken string

}

//write refreshtofken response and request
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
