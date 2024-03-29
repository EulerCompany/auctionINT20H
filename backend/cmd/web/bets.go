package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Bet struct {
	User string `json:"user"`
	Bet  int64  `json:"bet"`
}

type BetRepository interface {
	MakeBet(userId int, auctionId int, bet int64) (int64, error)
	GetAllBetsByAuction(auctionId int) ([]Bet, error)
}

type mysqlBetRepository struct {
	DB *sql.DB
}

func NewMySQLBetRepository(db *sql.DB) (*mysqlBetRepository, error) {
	return &mysqlBetRepository{DB: db}, nil
}

func (r *mysqlBetRepository) MakeBet(userId int, auctionId int, bet int64) (int64, error) {
	stmt := `INSERT INTO auction_bet (auction_id, user_id, bet) VALUES (?, ?, ?)`
	result, err := r.DB.Exec(stmt, auctionId, userId, bet)
    if err != nil {
        return 0, err
    }
    id, err := result.RowsAffected()
	return id, err
}

func (r *mysqlBetRepository) GetAllBetsByAuction(auctionId int) ([]Bet, error) {
	stmt := `SELECT u.name, ab.bet
    FROM auction_bet ab
    LEFT JOIN user u
        ON u.id = ab.user_id
    WHERE auction_id = ?`
	rows, err := r.DB.Query(stmt, auctionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bets []Bet
	for rows.Next() {
		var bet Bet
		if err := rows.Scan(
			&bet.User,
			&bet.Bet,
		); err != nil {
			return bets, err
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
	_, err := s.Repo.MakeBet(bet.UserId, bet.AuctionId, bet.Bet)
	return err
}

func (s *BetService) GetAllBetsByAuction(auctionId int) ([]Bet, error) {
	bets, err := s.Repo.GetAllBetsByAuction(auctionId)
	return bets, err
}
