package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: how to declare images??
type Auction struct {
	Id           int
	AuthorId     int
    Title        string `json:"title"`
	Description  string `json:"description"`
	StartPrice   int64  `json:"start_price"`
	CurrentPrice int64  `json:"current_price"`
	Status       string `json:"status"`
	StartDate    sql.NullTime `json:"start_date"`
	EndDate      sql.NullTime `json:"end_date"`
}

type AuctionRepository interface {
	CreateAuction(authorId int, title, description string, startPrice int64) (int64, error)
	InsertAuctionImage() error
	GetAllActiveAuctions() ([]Auction, error)
}

type mysqlAuctionRepository struct {
	DB *sql.DB
}

func NewMySQLAuctionRepository(db *sql.DB) (*mysqlAuctionRepository, error) {
	return &mysqlAuctionRepository{DB: db}, nil
}

func (r *mysqlAuctionRepository) CreateAuction(authorId int, title, description string, startPrice int64) (int64, error) {
	stmt := `INSERT INTO auction (author_id, title, description, start_price)
    VALUES (?, ?, ?, ?)`
	result, err := r.DB.Exec(stmt, authorId, title, description, startPrice)
	id, err := result.LastInsertId()
	fmt.Printf("error at createAuction is %v\n", err)
	fmt.Printf("Last inserted id: %d\n", id)
	return id, err
}

func (r *mysqlAuctionRepository) GetAllActiveAuctions() ([]Auction, error) {
	stmt := `SELECT * FROM auction WHERE status = 'active'`
	rows, err := r.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auctions []Auction
	for rows.Next() {
		var auction Auction
		if err := rows.Scan(
			&auction.Id,
            &auction.AuthorId,
			&auction.Title,
			&auction.Description,
			&auction.StartPrice,
			&auction.CurrentPrice,
			&auction.Status,
			&auction.StartDate,
			&auction.EndDate); err != nil {
			return nil, err
		}
		auctions = append(auctions, auction)
	}

	return auctions, nil
}

func (r *mysqlAuctionRepository) InsertAuctionImage() error {
	return nil
}

type AuctionService struct {
	Repo AuctionRepository
}

func NewAuctionService(repo AuctionRepository) *AuctionService {
	return &AuctionService{repo}
}

func (s *AuctionService) CreateAuction(auction Auction) error {
    // TODO: add author id
    id, err := s.Repo.CreateAuction(1, auction.Title, auction.Description, auction.StartPrice)
	log.Printf("created auction id: %d\n", id)
	return err
}

func (s *AuctionService) GetAllActiveAuctions() ([]Auction, error) {
    log.Printf("Calling get all active auctions\n")
	auctions, err := s.Repo.GetAllActiveAuctions()
	return auctions, err
}
