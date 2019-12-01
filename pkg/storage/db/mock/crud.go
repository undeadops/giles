package mockdb

import (
	"fmt"
	"time"

	"github.com/undeadops/giles"
)

func (mdb *MockDB) GetWatch() ([]*giles.Watch, error) {
	w := make([]*giles.Watch, 0)
	w = append(w, &giles.Watch{Symbol: "MDB"})

	return w, nil
}

func (mdb *MockDB) UpsertWatch(w *giles.Watch) error {
	if w.Symbol == "FAIL" {
		return fmt.Errorf("FAILED Upserting Watch")
	}
	return nil
}

func (mdb *MockDB) DeleteWatch(w string) error {
	if w == "FAIL" {
		return fmt.Errorf("FAILED Deleting Watch")
	}
	return nil
}

// GetAllTickers - Mocked Function to Retrieve a list of Ticker Watches
func (mdb *MockDB) GetAllTickers() ([]*giles.Ticker, error) {
	tic := make([]*giles.Ticker, 0)
	tic = append(tic, &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 14, 04, 06, 58235543, time.UTC), Price: 121.57})
	tic = append(tic, &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 15, 04, 05, 42984843, time.UTC), Price: 121.34})
	tic = append(tic, &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 16, 04, 06, 22754343, time.UTC), Price: 121.61})
	tic = append(tic, &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 17, 04, 06, 52345543, time.UTC), Price: 121.75})
	tic = append(tic, &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 18, 04, 05, 72345543, time.UTC), Price: 121.89})
	tic = append(tic, &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 19, 04, 07, 82324343, time.UTC), Price: 122.01})
	return tic, nil
}

// SaveTicker - Mocked Fuction for Saving Ticker Watch
func (mdb *MockDB) SaveTicker(tic *giles.Ticker) error {
	if tic.Symbol == "" {
		return fmt.Errorf("Missing Ticker Symbol")
	}
	return nil
}

// GetTicker - Get Ticker Watcher
func (mdb *MockDB) GetTicker(tic string) (*giles.Ticker, error) {
	if tic == "FAIL" {
		return nil, fmt.Errorf("There was an error getting %v", tic)
	}
	return &giles.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 11, 21, 42, 19, 23465435, time.UTC), Price: 125.43}, nil
}

// DeleteTicker - Delete Ticker Watcher
func (mdb *MockDB) DeleteTicker(tic string) error {
	if tic == "FAIL" {
		return fmt.Errorf("There was an error deleting %v", tic)
	}
	return nil
}
