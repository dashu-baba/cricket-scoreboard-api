package models

// PlayerType godoc
// @Summary Define different types of palayer
type PlayerType int

const (
	// Batsman godoc
	Batsman PlayerType = 0
	// Bowler godoc
	Bowler PlayerType = 1
	// AllRouner godoc
	AllRouner PlayerType = 3
)
