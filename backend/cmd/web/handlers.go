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
	UserName  string
	AuctionId int
	Bet       int64 `json:"bet"`
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

	tokenMap := r.Context().Value("token info").(map[string]string)

	userId, err := app.user.GetIdByUsername(tokenMap["username"])

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]interface{})
	resp["user"] = tokenMap["username"]
	resp["token"] = token
	resp["userId"] = userId

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

	tokenMap := r.Context().Value("token info").(map[string]string)
	userId, err := app.user.GetIdByUsername(tokenMap["username"])
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	resp, err := app.auction.CreateAuction(int64(userId), auctionReq)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	app.JSONResponse(w, http.StatusCreated, resp)
}

func (app *application) getActiveAuctions(w http.ResponseWriter, r *http.Request) {
	resp, err := app.auction.GetAllActiveAuctions()
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	w.Write(data)
	fmt.Println(err)
}

func (app *application) makebet(w http.ResponseWriter, r *http.Request) {
	tokenMap := r.Context().Value("token info").(map[string]string)
	betData, err := parseBetData(r)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	userId, err := app.user.GetIdByUsername(tokenMap["username"])
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	betData.UserId = userId
	betData.AuctionId, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	err = app.bet.MakeBet(betData)
    auction, err := app.auction.AcceptBet(betData.AuctionId, betData.Bet)

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["user"] = tokenMap["username"]
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	resp, err := app.auction.GetAllActiveAuctionsByUserId(id)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		logErrorDumbExit(w, err)
		return
	}
	w.Write(data)
	fmt.Println(err)
}

func (app *application) getAuctionBets(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	bets, err := app.bet.GetAllBetsByAuction(id)

	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, bets)
}

func (app *application) updateAuction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	var auctionUpdate UpdateAuctionRequest
	err = decodeJSONBody(w, r, &auctionUpdate)

	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	auction, err := app.auction.Repo.GetAuctionById(id)

	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	auction.Title = auctionUpdate.Title
	auction.Description = auctionUpdate.Description
	err = app.auction.UpdateAuction(auction)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, map[string]string{"status": "updated"})
}

func (app *application) getAuctionById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	auction, err := app.auction.Repo.GetAuctionById(id)

	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, auction)
}

func (app *application) getImagesByAuctionId(w http.ResponseWriter, r *http.Request) {
	auctionId, err := strconv.ParseInt(chi.URLParam(r, "auctionId"), 10, 64)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	imgs, err := app.auction.GetImagesByAuctionId(auctionId)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	app.JSONResponse(w, http.StatusOK, imgs)
}

func (app *application) getImageByAuctionImageId(w http.ResponseWriter, r *http.Request) {
	imageId, err := strconv.ParseInt(chi.URLParam(r, "imageId"), 10, 64)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}
	auctionId, err := strconv.ParseInt(chi.URLParam(r, "auctionId"), 10, 64)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	img, err := app.auction.GetImageByAuctionAndImageId(auctionId, imageId)
	if err != nil {
		app.JSONErrorResponse(w, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, img)

}
