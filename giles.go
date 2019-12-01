package giles

import (
	"time"
)

// {
//   "_id" : ObjectId("5b4690e047f49a04be523cbe"),
//   "p" : 56.58,
//   "symbol" : "MDB",
//   "d" : ISODate("2018-06-30T00:00:02Z")
// },

// Ticker - Document structure for StockTicker Collection
type Ticker struct {
	//ID     bson.TypeObjectID `json:"id" bson:"_id"`
	Symbol string    `json:"symbol" bson:"symbol"`
	Date   time.Time `json:"date" bson:"date"`
	Price  float32   `json:"price" bson:"price"`
}

// Watch - Document structure to Configure What Tickers to Watch
type Watch struct {
	Symbol      string    `json:"symbol" bson:"symbol"`
	CompanyName string    `json:"company_name" bson:"company_name"`
	ExtraArgs   []string  `json:"extra_args" bson:"extra_args"`
	CreateDate  time.Time `json:"create_date" bson:"create_date"`
	ModDate     time.Time `json:"mod_date" bson:"mod_date"`
}
