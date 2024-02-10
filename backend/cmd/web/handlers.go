package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
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
	resp["message"] = "Signed up successfully!"

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

    id, err := app.user.Authenticate(loginData)
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
	resp := make(map[string]interface{})
	resp["token"] = tokenStr
	resp["message"] = "Logged in succesfully!"
    resp["userId"] = id
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (app *application) getMe(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	// Extract the token from the Authorization header
	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
		return
	}
	token := authToken[1]

	feedName := r.Context().Value("token info").(map[string]string)

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["user"] = feedName["username"]
	resp["token"] = token

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}

// TODO: where should validation go, service/controller????
// TODO: link auction with the author, needs Max's changes
func (app *application) createAuction(w http.ResponseWriter, r *http.Request) {
	log.Println("create auction executing")

	var auction Auction
	err := json.NewDecoder(r.Body).Decode(&auction)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	log.Printf("Parsed %v\n", auction)

	err = app.auction.CreateAuction(auction)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	log.Println("finished executing create auction")

}

func (app *application) getAllActive(w http.ResponseWriter, r *http.Request) {
	log.Println("show all active executing")

	resp, err := app.auction.GetAllActiveAuctions()
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
    
	for i, auction := range resp {
		log.Printf("auction %d = %v", i, auction)
	}
    data, err :=json.Marshal(resp)
    if err != nil {
        logErrorDumbExit(w, err)
        return
    }
    w.Write(data)
    fmt.Println(err)
}


func (app *application) getAllActiveByUserId(w http.ResponseWriter, r *http.Request) {
	log.Println("show all active by userId")

    id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	resp, err := app.auction.GetAllActiveAuctionsByUserId(id)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
    
	for i, auction := range resp {
		log.Printf("auction %d = %v", i, auction)
	}
    data, err :=json.Marshal(resp)
    if err != nil {
        logErrorDumbExit(w, err)
        return
    }
    w.Write(data)
    fmt.Println(err)
}
