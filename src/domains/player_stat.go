//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlayerStat godoc
// @Summary Define PlayerStat model
type PlayerStat struct {
	ID          primitive.ObjectID
	LimitedStat PlayerStatDetails
	TestStat    PlayerStatDetails
	PlayerID    primitive.ObjectID
}

// PlayerStatDetails godoc
// @Summary Define PlayerStatDetails model
type PlayerStatDetails struct {
	PlayerBattingStatDetails PlayerBattingStatDetails
	PlayerBowlingStatDetails PlayerBowlingStatDetails
}

// PlayerBattingStatDetails godoc
// @Summary Define PlayerBattingStatDetails model
type PlayerBattingStatDetails struct {
	InningsPlayed  int
	NumberOfNotOut int
	Run            int
	BallFaced      int
	Four           int
	Six            int
	StrikeRate     float64
	Average        float64
	Fifties        int
	Centuries      int
	Highest        int
}

// PlayerBowlingStatDetails godoc
// @Summary Define PlayerBowlingStatDetails model
type PlayerBowlingStatDetails struct {
	InningsPlayed int
	Wicket        int
	BallDelivered int
	Over          int
	Wide          int
	NoBall        int
	Average       float64
	Economy       float64
	BestFigure    string
	Five          int
	Ten           int
}
