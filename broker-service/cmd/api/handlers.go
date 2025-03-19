package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}
type AuthPayload struct {
	Email    string `json:"email"`
	password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    "broker response",
	}
	_ = app.writeJson(w, http.StatusOK, payload)

}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	//handle submission request from frontend..get json data
	//we will read json data
	var payload RequestPayload
	err := app.readJson(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch RequestPayload.Action {
	case "auth":
		//auth request

	default:
		app.errorJSON(w, errors.New("Invalid action"))

	}

}
func (app *Config) Authenticate(w http.ResponseWriter, payload AuthPayload) {
	//create a json we send to the authentication micro service

	jsonData, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	//call the authentication service
	resp, err := http.Post("http://localhost:8081/auth", "application/json", bytes.NewBuffer(jsonData))
}
