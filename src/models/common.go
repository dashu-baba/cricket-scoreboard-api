package models

// PlayerType godoc
// @Summary Define different types of player
type PlayerType int

const (
	// Batsman godoc
	Batsman PlayerType = 0
	// Bowler godoc
	Bowler PlayerType = 1
	// AllRouner godoc
	AllRouner PlayerType = 3
)

// GameType godoc
// @Summary Define different types of game
type GameType int

const (
	// Tournament godoc
	Tournament GameType = 0
	// Bilateral godoc
	Bilateral GameType = 1
)

// MatchType godoc
// @Summary Define different types of match
type MatchType int

const (
	// LimitedOver godoc
	LimitedOver MatchType = 0
	// Test godoc
	Test MatchType = 1
)
