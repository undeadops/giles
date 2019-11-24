package ticker

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
