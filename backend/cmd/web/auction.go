package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CreateAuctionRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartPrice  int64     `json:"start_price"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type CreateAuctionResponse struct {
	Id int64 `json:"auction_id"`
}

type UpdateAuctionRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Auction struct {
	Id           int64
	AuthorId     int64
	Title        string
	Description  string
	StartPrice   int64
	CurrentPrice int64
	Status       string
	StartDate    sql.NullTime
	EndDate      sql.NullTime
}

type AuctionRepository interface {
	CreateAuction(authorId int64, title, description string, startPrice int64, startDate, endDate time.Time) (int64, error)
	GetAllActiveAuctions() ([]Auction, error)
	GetAuctionById(auctionId int) (Auction, error)
	UpdateCurrentPriceAuction(auction Auction) error
	GetAllActiveAuctionsByUserId(userId int) ([]Auction, error)
}

type mysqlAuctionRepository struct {
	DB *sql.DB
}

func NewMySQLAuctionRepository(db *sql.DB) (*mysqlAuctionRepository, error) {
	return &mysqlAuctionRepository{DB: db}, nil
}

func (r *mysqlAuctionRepository) CreateAuction(authorId int64, title, description string, startPrice int64, startDate, endDate time.Time) (int64, error) {
	log.Println("Creating new auction")
	stmt := `INSERT INTO auction (author_id, title, description, start_price, status, start_date, end_date) VALUES (?, ?, ?, ?,'active', ?, ?)`
	result, err := r.DB.Exec(stmt, authorId, title, description, startPrice, startDate, endDate)
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

func (r *mysqlAuctionRepository) GetAuctionById(auctionId int) (Auction, error) {
	stmt := `SELECT * FROM auction WHERE id = ?`
	rows, err := r.DB.Query(stmt, auctionId)
	if err != nil {
		return Auction{}, err
	}
	defer rows.Close()
	var auction Auction
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
			return auction, err
		}
	}

	return auction, nil
}

func (r *mysqlAuctionRepository) GetAllActiveAuctionsByUserId(userId int) ([]Auction, error) {
	stmt := `SELECT * FROM auction WHERE status = 'active' and author_id = ?`
	rows, err := r.DB.Query(stmt, userId)
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

func (r *mysqlAuctionRepository) UpdateCurrentPriceAuction(auction Auction) error {
	stmt := `UPDATE auction SET current_price = ? WHERE id = ?`
	result, err := r.DB.Exec(stmt, auction.CurrentPrice, auction.Id)
	fmt.Printf("error at update auction is %v\n", err)
	fmt.Printf("result at update auction is %v\n", result)
	return err
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

func (s *AuctionService) validateAuctionRequset(auction CreateAuctionRequest) error {

	if len(auction.Title) == 0 {
		return errors.New("title cannot be empty")
	}
	if len(auction.Description) == 0 {
		return errors.New("description cannot be empty")
	}
	if auction.StartPrice <= 0 {
		return errors.New("start price must be greater than zero")
	}
	if auction.StartDate.Before(time.Now()) {
		return errors.New("start date must be in the future")
	}
	if auction.EndDate.Before(auction.StartDate) {
		return errors.New("end date must be after start date")
	}

	return nil
}

func (s *AuctionService) CreateAuction(userId int64, auction CreateAuctionRequest) (CreateAuctionResponse, error) {
	if err := s.validateAuctionRequset(auction); err != nil {
		return CreateAuctionResponse{}, err
	}
	log.Println("Validate")

	id, err := s.Repo.CreateAuction(userId, auction.Title, auction.Description, auction.StartPrice, auction.StartDate, auction.EndDate)
	if err != nil {
		return CreateAuctionResponse{}, err
	}
	return CreateAuctionResponse{Id: id}, nil
}

func (s *AuctionService) UpdateAuction(newAuction CreateAuctionRequest) error {

	return nil
}

func (s *AuctionService) GetAllActiveAuctions() ([]Auction, error) {
	log.Printf("Calling get all active auctions\n")
	auctions, err := s.Repo.GetAllActiveAuctions()
	return auctions, err
}

func (s *AuctionService) AcceptBet(auctionId int, bet int64) (Auction, error) {
	log.Printf("Accepting bet\n")
	auction, err := s.Repo.GetAuctionById(auctionId)
	fmt.Printf("Accept bet for %v\n", auction)
	fmt.Printf("Accept bet for id  %d\n", auction.Id)
	if auction.CurrentPrice < bet {
		auction.CurrentPrice = bet
		err := s.Repo.UpdateCurrentPriceAuction(auction)
		return auction, err
	}

	return auction, err
}

func (s *AuctionService) GetAllActiveAuctionsByUserId(userId int) ([]Auction, error) {
	log.Printf("Calling get all active auctions\n")
	auctions, err := s.Repo.GetAllActiveAuctionsByUserId(userId)
	return auctions, err
}
