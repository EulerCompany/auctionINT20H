package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Bet struct {
	Id        int
	UserId    int
	AuctionId int
	bet       float64 `json:"bet"`
}

type BetRepository interface {
	MakeBet(userId int, auctionId int, bet float64) (int64, error)
	GetAllBetsByAuction(auctionId int) ([]Bet, error)
}

type mysqlBetRepository struct {
	DB *sql.DB
}

func NewMySQLBetRepository(db *sql.DB) (*mysqlBetRepository, error) {
	return &mysqlBetRepository{DB: db}, nil
}

func (r *mysqlBetRepository) MakeBet(userId int, auctionId int, bet float64) (int64, error) {
	stmt := `INSERT INTO auction_bet (auction_id, user_id, bet) VALUES (?, ?, ?)`
	fmt.Printf("Last inserted id: %d\n", auctionId)
	fmt.Printf("Last inserted id: %d\n", userId)
	fmt.Printf("Last inserted id: %f\n", bet)
	result, err := r.DB.Exec(stmt, auctionId, userId, bet)
	id, err := result.RowsAffected()
	fmt.Printf("error at makeBet is %v\n", err)
	fmt.Printf("Last inserted id: %d\n", id)
	return id, err
}

func (r *mysqlBetRepository) GetAllBetsByAuction(auctionId int) ([]Bet, error) {
	stmt := `SELECT * FROM auction_bet WHERE auctionId = ?`
	rows, err := r.DB.Query(stmt, auctionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bets []Bet
	for rows.Next() {
		var bet Bet
		if err := rows.Scan(
			&bet.Id,
			&bet.UserId,
			&bet.AuctionId,
			&bet.bet,
		); err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}

	return bets, nil
}

type BetService struct {
	Repo BetRepository
}

func NewBetService(repo BetRepository) *BetService {
	return &BetService{repo}
}

func (s *BetService) MakeBet(bet BetData) error {
	fmt.Println("Making bet1.1")
	id, err := s.Repo.MakeBet(bet.UserId, bet.AuctionId, bet.Bet)
	log.Printf("created bet id: %d\n", id)
	return err
}

func (s *BetService) GetAllBetsByAuction(auctionId int) ([]Bet, error) {
	log.Printf("Calling get all bets by auction\n")
	auctions, err := s.Repo.GetAllBetsByAuction(auctionId)
	return auctions, err
}