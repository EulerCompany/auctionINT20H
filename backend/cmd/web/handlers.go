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

type BetData struct {
	UserId    int
	AuctionId int
	Bet       float64 `json:"bet"`
}

func parseLoginData(r *http.Request) (LoginData, error) {
	var ld LoginData
	// TODO: add proper validation
	err := json.NewDecoder(r.Body).Decode(&ld)

	return ld, err
}

func parseBetData(r *http.Request) (BetData, error) {
	var bd BetData
	// TODO: add proper validation
	err := json.NewDecoder(r.Body).Decode(&bd)

	return bd, err
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

func (app *application) createAuction(w http.ResponseWriter, r *http.Request) {
	var auctionReq CreateAuctionRequest
	err := decodeJSONBody(w, r, &auctionReq)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
    log.Printf("Parsed create auction payload: %v\n", auctionReq)
    
    resp, err := app.auction.CreateAuction(1, auctionReq)
	if err != nil {
        app.JSONErrorResponse(w, err)
		return
	}

    app.JSONResponse(w, http.StatusCreated, resp)
}

func (app *application) getActiveAuctions(w http.ResponseWriter, r *http.Request) {
	log.Println("show all active executing")

	resp, err := app.auction.GetAllActiveAuctions()
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}

	for i, auction := range resp {
		log.Printf("auction %d = %v", i, auction)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Write(data)
	fmt.Println(err)
}

func (app *application) makebet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Making bet0")
	feedName := r.Context().Value("token info").(map[string]string)
	betData, err := parseBetData(r)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	userId, err := app.user.GetIdByUsername(feedName["username"])

	if err != nil {
		logErrorDumbExit(w, err)
		return
	}

	betData.UserId = userId
	betData.AuctionId, err = strconv.Atoi(chi.URLParam(r, "id"))

	fmt.Println("Making bet1")
	fmt.Printf("result at update auction is %v\n", betData)
	err = app.bet.MakeBet(betData)
	fmt.Println("Making bet2")
	auction, err := app.auction.AcceptBet(betData.AuctionId, betData.Bet)

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["user"] = feedName["username"]
	resp["auction"] = auction.Title

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}


func (app *application) getActiveAuctionsByUser(w http.ResponseWriter, r *http.Request) {
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
