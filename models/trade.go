package models

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TradeModel ...
type TradeModel struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	BuyDate        time.Time          `bson:"buydate" json:"buydate"`
	Shares         float64            `bson:"shares" json:"shares"`
	BuyPrice       float64            `bson:"buyprice" json:"buyprice"`
	AvePrice       float64            `bson:"aveprice" json:"aveprice"`
	BuyComm        float64            `bson:"buycomm" json:"buycomm"`
	SellPrice      float64            `bson:"sellprice" json:"sellprice"`
	SellComm       float64            `bson:"sellcomm" json:"sellcomm"`
	Profit         float64            `bson:"profit" json:"profit"`
	ROI            float32            `bson:"roi" json:"roi"`
	BreakSellPrice float64            `bson:"breaksellprice" json:"breaksellprice"`
	SellDate       time.Time          `bson:"selldate" json:"selldate"`
	SecuritiesID   primitive.ObjectID `bson:"securities_id" json:"securities_id"`
	StockID        primitive.ObjectID `bson:"stock_id" json:"stock_id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"user_id"`
	mu             sync.Mutex
}

// Calculate break-even selling price
func (trade *TradeModel) Calculate() {

}
