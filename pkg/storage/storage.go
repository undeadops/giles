package storage

import (
	ticker "github.com/undeadops/giles"
)

// DataAccessLayer - Storage Abstraction for objects
type DataAccessLayer interface {
	GetAllTickers() ([]*ticker.Ticker, error)
	SaveTicker(*ticker.Ticker) error
	GetTicker(string) (*ticker.Ticker, error)
	DeleteTicker(string) error
}
