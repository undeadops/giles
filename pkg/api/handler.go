package api

import (
	"net/http"
)

// TickerWatcherList - Http Handler for Listing Tickers Watches
func (s Server) TickerWatcherList(w http.ResponseWriter, r *http.Request) {
	_, err := s.DB.GetAllTickers()
	if err != nil {
		w.Write([]byte("There was an error in listing tickers"))
	}
	msg := `{"alive": true }`
	w.Write([]byte(msg))
}

// Create - Http Handler for Creating Ticker Watches
func (s Server) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa create"))
}

// Get - Http Handler for Getting Ticker
func (s Server) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa get"))
}

// Update - Http Handler for Updating Ticker Watcher
func (s Server) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa update"))
}

// Delete - Http Handler for Deleting Ticker Watcher
func (s Server) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa delete"))
}
