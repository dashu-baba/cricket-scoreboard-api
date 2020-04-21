package responsemodels

import "cricket-scoreboard-api/src/models"

// MatchSummary godoc
// @Summary Define MatchSummary model
type MatchSummary struct {
	ID               string             `json:"id" form:"id" xml:"id"`
	MatchType        models.MatchType   `json:"matchType" form:"matchtype" xml:"matchType"`
	Status           models.SeriesState `json:"status" form:"status" xml:"status"`
	CurrentInningsID string             `json:"currentInningsId" form:"currentinningsid" xml:"currentInningsId"`
	SeriesID         string             `json:"seriesId" form:"seriesid" xml:"seriesId"`
}
