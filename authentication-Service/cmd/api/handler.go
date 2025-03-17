package main

import (
	"errors"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		password string `json:"password"`
	}
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	user, err := app.Models.User.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invalid Credentials"), http.StatusBadRequest)
		return
	}
	valid, err := app.Models.User.PasswordMatches(requestPayload.password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invalid Credentials"), http.StatusBadRequest)
		return
	}
	// token, err := app.Models.User.GenerateToken(user.ID)
	// if err != nil {
	//     app.errorJSON(w, err, http.StatusInternalServerError)
	//     return
	// }
	// responsePayload := struct {
	//     Token string `json:"token"`
	// }{
	//     Token: token,
	// }
	requestPayloads := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    user,
	}
	app.writeJson(w, http.StatusOK, requestPayloads)

}
