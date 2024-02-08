package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: how to declare images??
type Auction struct {
    Title string `json:"title"`    
    Description string `json:"description"`
    StartPrice int64 `json:"start_price"`
}

type AuctionRepository interface {
    CreateAuction(title, description string, startPrice int64) (int64, error)
    InsertAuctionImage() error
}

type mysqlAuctionRepository struct {
    DB *sql.DB
}

func NewMySQLAuctionRepository (db *sql.DB) (*mysqlAuctionRepository, error) {
    return &mysqlAuctionRepository{DB:db}, nil
}

func (r *mysqlAuctionRepository) CreateAuction(title, description string, startPrice int64) (int64, error) {
    stmt := `INSERT INTO auction (title, description, start_price)
    VALUES (?, ?, ?)`
    result, err := r.DB.Exec(stmt, title, description, startPrice)
    id, err := result.LastInsertId()
    fmt.Printf("error at createAuction is %v\n", err)
    fmt.Printf("Last inserted id: %d\n", id)
    return id, err
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
    id, err := s.Repo.CreateAuction(auction.Title, auction.Description, auction.StartPrice)
    log.Printf("created auction id: %d\n", id)
    return err
}


