package responsemodels

import "cricket-scoreboard-api/src/models"

// InningsSummary godoc
// @Summary Define InningsSummary model
type InningsSummary struct {
	ID          string             `json:"id" form:"id" xml:"id"`
	Status      models.SeriesState `json:"status" form:"status" xml:"status"`
	MatchID     string             `json:"matchId" form:"matchid" xml:"matchId"`
	Batsmans    []InningsBatsman   `json:"batsmans" form:"batsmans" xml:"batsmans"`
	Bowlers     []InningsBowler    `json:"bowlers" form:"bowlers" xml:"bowlers"`
	CurrentOver Over               `json:"currentOver" form:"currentover" xml:"currentOver"`
}

// InningsBatsman godoc
// @Summary Define InningsBatsman model
type InningsBatsman struct {
	ID         string `json:"id" form:"id" xml:"id"`
	PlayerID   string `json:"playerId" form:"playerid" xml:"playerId"`
	Run        int    `json:"run" form:"run" xml:"run"`
	Ball       int    `json:"ball" form:"ball" xml:"ball"`
	IsInStrike bool   `json:"inInStrike" form:"ininstrike" xml:"inInStrike"`
	Four       int    `json:"four" form:"four" xml:"four"`
	Six        int    `json:"six" form:"six" xml:"six"`
}

// InningsBowler godoc
// @Summary Define InningsBowler model
type InningsBowler struct {
	ID       string `json:"id" form:"id" xml:"id"`
	PlayerID string `json:"playerId" form:"playerid" xml:"playerId"`
	Run      int    `json:"run" form:"run" xml:"run"`
	Over     int    `json:"over" form:"over" xml:"over"`
	IsActive bool   `json:"isActive" form:"isactive" xml:"isactive"`
	Wickets  int    `json:"wickets" form:"wickets" xml:"wickets"`
}

// Over godoc
// @Summary Define Over model
type Over struct {
	ID       string `json:"id" form:"id" xml:"id"`
	Sequence string `json:"sequence" form:"sequence" xml:"sequence"`
	Ball     int    `json:"ball" form:"ball" xml:"ball"`
}
