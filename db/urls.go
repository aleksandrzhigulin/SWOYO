package db

import (
	"database/sql"
	"errors"
)

func (db Database) ShortUrlExists(url string) bool {
	var rowUrl string
	query := `SELECT url FROM urls WHERE short_url = $1;`
	row := db.Connection.QueryRow(query, url).Scan(&rowUrl)

	if errors.Is(row, sql.ErrNoRows) {
		return false
	}
	return true
}

func (db Database) SaveShortUrl(shortUrl string, url string) error {
	query := `INSERT INTO urls (short_url, url) VALUES ($1, $2)`
	var shortUrlResult string
	err := db.Connection.QueryRow(query, shortUrl, url).Scan(&shortUrlResult)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) GetFullUrl(shortUrl string) (error, string) {
	var value string
	query := `SELECT url FROM urls WHERE short_url = $1`
	err := db.Connection.QueryRow(query, shortUrl).Scan(&value)
	if err != nil {
		return err, ""
	}
	return nil, value
}
