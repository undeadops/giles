package mockdb

import (
	"fmt"
	"time"

	ticker "github.com/undeadops/giles"
)

// GetAllTickers - Mocked Function to Retrieve a list of Ticker Watches
func (mdb *MockDB) GetAllTickers() ([]*ticker.Ticker, error) {
	tic := make([]*ticker.Ticker, 0)
	tic = append(tic, &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 14, 04, 06, 58235543, time.UTC), Price: 121.57})
	tic = append(tic, &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 15, 04, 05, 42984843, time.UTC), Price: 121.34})
	tic = append(tic, &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 16, 04, 06, 22754343, time.UTC), Price: 121.61})
	tic = append(tic, &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 17, 04, 06, 52345543, time.UTC), Price: 121.75})
	tic = append(tic, &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 18, 04, 05, 72345543, time.UTC), Price: 121.89})
	tic = append(tic, &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 9, 19, 04, 07, 82324343, time.UTC), Price: 122.01})
	return tic, nil
}

// SaveTicker - Mocked Fuction for Saving Ticker Watch
func (mdb *MockDB) SaveTicker(tic *ticker.Ticker) error {
	if tic.Symbol == "" {
		return fmt.Errorf("Missing Ticker Symbol")
	}
	return nil
}

// GetTicker - Get Ticker Watcher
func (mdb *MockDB) GetTicker(tic string) (*ticker.Ticker, error) {
	if tic == "FAIL" {
		return nil, fmt.Errorf("There was an error getting %v", tic)
	}
	return &ticker.Ticker{Symbol: "MDB", Date: time.Date(2019, 11, 11, 21, 42, 19, 23465435, time.UTC), Price: 125.43}, nil
}

// DeleteTicker - Delete Ticker Watcher
func (mdb *MockDB) DeleteTicker(tic string) error {
	if tic == "FAIL" {
		return fmt.Errorf("There was an error deleting %v", tic)
	}
	return nil
}
