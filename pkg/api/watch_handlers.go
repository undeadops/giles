package api

import "net/http"

// ListWatch - Http Handler for Listing Watches
func (s Server) ListWatch(w http.ResponseWriter, r *http.Request) {
	// _, err := s.DB.Get('watch', nil)
	// if err != nil {
	// 	w.Write([]byte("There was an error in listing tickers"))
	// }
	// msg := `{"alive": true }`
	// w.Write([]byte(msg))
}

// Create - Http Handler for Creating Ticker Watches
func (s Server) CreateWatch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa create"))
}
