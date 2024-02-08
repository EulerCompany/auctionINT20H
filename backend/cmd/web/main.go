package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	user *UserModel
    auction *AuctionService
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
    dsn := flag.String("dsn", "web:pass@tcp(db)/auction", "MySQL DSN")

	flag.Parse()
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    auctionRepo, _ := NewMySQLAuctionRepository(db)
    auctionService := NewAuctionService(auctionRepo)

	app := &application{
		user: &UserModel{db},
        auction: auctionService,
    }

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
