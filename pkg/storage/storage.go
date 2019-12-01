package storage

import (
	"github.com/undeadops/giles"
)

// DataAccessLayer - Storage Abstraction for objects
type DataAccessLayer interface {
	GetAllTickers() ([]*giles.Ticker, error)
	SaveTicker(*giles.Ticker) error
	GetTicker(string) (*giles.Ticker, error)
	DeleteTicker(string) error
	GetWatch() ([]*giles.Watch, error)
	UpsertWatch(*giles.Watch) error
	DeleteWatch(string) error
}
