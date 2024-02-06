package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginData struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func parseLoginData(r *http.Request) (LoginData, error) {
	var ld LoginData
	// TODO: add proper validation
	err := json.NewDecoder(r.Body).Decode(&ld)

	return ld, err
}

// TODO: temp function to just log errors and send 500 to client
func logErrorDumbExit(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	log.Println("signup user called")

	loginData, err := parseLoginData(r)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	err = app.user.CreateUser(loginData)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	log.Println(err)

	tokenStr, err := createToken(loginData.Name)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	fmt.Println(tokenStr)

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["token"] = tokenStr
	resp["message"] = "Logged in succesfully!"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("login user called")
	loginData, err := parseLoginData(r)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}

	_, err = app.user.Authenticate(loginData)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	tokenStr, err := createToken(loginData.Name)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	fmt.Println(tokenStr)

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["token"] = tokenStr
	resp["message"] = "Logged in succesfully!"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
