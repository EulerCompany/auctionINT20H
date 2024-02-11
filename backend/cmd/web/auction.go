package main

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Base64File struct {
	Base64 []byte `json:"base64"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Type   string `json:"type"`
}
type CreateAuctionRequest struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartPrice  int64        `json:"start_price"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Files       []Base64File `json:"files"`
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

type Image struct {
	Image []byte
}

type AuctionRepository interface {
	CreateAuction(authorId int64, title, description string, startPrice int64, startDate, endDate time.Time) (int64, error)
	GetAllActiveAuctions() ([]Auction, error)
	GetAuctionById(auctionId int) (Auction, error)
	UpdateCurrentPriceAuction(auction Auction) error
	GetAllActiveAuctionsByUserId(userId int) ([]Auction, error)
	UpdateAuction(auctionId int64, title, description string) error
	InsertAuctionPhotos(auctionId int64, photos [][]byte) error
	GetImageByAuctionIdAndImageId(auctionId int64, imageId int64) (Image, error)
	GetImagesByAuctionId(auctionId int64) ([]Image, error)
}

type mysqlAuctionRepository struct {
	DB *sql.DB
}

func NewMySQLAuctionRepository(db *sql.DB) (*mysqlAuctionRepository, error) {
	return &mysqlAuctionRepository{DB: db}, nil
}

func (r *mysqlAuctionRepository) GetImagesByAuctionId(auctionId int64) ([]Image, error) {
	stmt := `SELECT img FROM auction_image WHERE auction_id = ?`

	rows, err := r.DB.Query(stmt, auctionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var imgs []Image
	for rows.Next() {
		var img Image
		if err := rows.Scan(
			&img.Image); err != nil {
			return nil, err
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func (r *mysqlAuctionRepository) GetImageByAuctionIdAndImageId(auctionId, imageId int64) (Image, error) {
	stmt := `SELECT img FROM auction_image WHERE auction_id = ? AND id = ?`
	row := r.DB.QueryRow(stmt, auctionId, imageId)
	var img Image
	err := row.Scan(&img.Image)
	return img, err
}

func (r *mysqlAuctionRepository) CreateAuction(authorId int64, title, description string, startPrice int64, startDate, endDate time.Time) (int64, error) {
	stmt := `INSERT INTO auction (author_id, title, description, start_price, status, start_date, end_date) VALUES (?, ?, ?, ?,'active', ?, ?)`
	result, err := r.DB.Exec(stmt, authorId, title, description, startPrice, startDate, endDate)
	if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
	return id, err
}

func (t *mysqlAuctionRepository) InsertAuctionPhotos(auctionId int64, photos [][]byte) error {
	stmt := `INSERT INTO auction_image (auction_id, img) VALUES `
	const rowSQL = "(?, ?)"
	
    var inserts []string
	vals := []interface{}{}

	for i := 0; i < len(photos); i++ {
		vals = append(vals, auctionId, photos[i])
		inserts = append(inserts, rowSQL)
	}
	sqlInsert := stmt + strings.Join(inserts, ",")
	_, err := t.DB.Exec(sqlInsert, vals...)
	return err
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
	var auctions Auction
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
		auctions = auction
	}
	return auctions, nil
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
	_, err := r.DB.Exec(stmt, auction.CurrentPrice, auction.Id)
	return err
}

func (r *mysqlAuctionRepository) UpdateAuction(auctionId int64, title, description string) error {
	stmt := `UPDATE auction SET title = ?, description = ? 
    WHERE id = ?`
	_, err := r.DB.Exec(stmt, title, description, auctionId)
	return err
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

	id, err := s.Repo.CreateAuction(userId, auction.Title, auction.Description, auction.StartPrice, auction.StartDate, auction.EndDate)
	if err != nil {
		return CreateAuctionResponse{}, err
	}
	photos := make([][]byte, len(auction.Files))

	for i, photo := range auction.Files {
		photos[i] = photo.Base64
	}
	if err = s.Repo.InsertAuctionPhotos(id, photos); err != nil {
        return CreateAuctionResponse{}, err
    }
    
    return CreateAuctionResponse{Id: id}, nil
}

func (s *AuctionService) UpdateAuction(auction Auction) error {
	if len(auction.Title) == 0 {
		return errors.New("title cannot be empty")
	}
	if len(auction.Description) == 0 {
		return errors.New("description cannot be empty")
	}

	err := s.Repo.UpdateAuction(auction.Id, auction.Title, auction.Description)
	return err
}

func (s *AuctionService) GetAllActiveAuctions() ([]Auction, error) {
	auctions, err := s.Repo.GetAllActiveAuctions()
	return auctions, err
}

func (s *AuctionService) AcceptBet(auctionId int, bet int64) (Auction, error) {
	auction, err := s.Repo.GetAuctionById(auctionId)
	if auction.CurrentPrice < bet {
		auction.CurrentPrice = bet
		err := s.Repo.UpdateCurrentPriceAuction(auction)
		return auction, err
	}

	return auction, err
}

func (s *AuctionService) GetAllActiveAuctionsByUserId(userId int) ([]Auction, error) {
	auctions, err := s.Repo.GetAllActiveAuctionsByUserId(userId)
	return auctions, err
}

type ImageResponse struct {
	Base64 []byte `json:"img_base64"`
}

func (s *AuctionService) GetImageByAuctionAndImageId(auctionId, imageId int64) (*ImageResponse, error) {
	img, err := s.Repo.GetImageByAuctionIdAndImageId(auctionId, imageId)
	if err != nil {
		return nil, err
	}

	return &ImageResponse{Base64: img.Image}, nil
}

func (s *AuctionService) GetImagesByAuctionId(auctionId int64) ([]ImageResponse, error) {
	imgs, err := s.Repo.GetImagesByAuctionId(auctionId)
	if err != nil {
		return nil, err
	}
	resp := make([]ImageResponse, len(imgs))
	for i, img := range imgs {
		resp[i].Base64 = img.Image
	}
	return resp, nil
}
