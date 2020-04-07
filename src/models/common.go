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
	AllRouner PlayerType = 2
)

// OutType godoc
// @Summary Define different types of wicket
type OutType int

const (
	// Bowled godoc
	Bowled OutType = 0
	// Caught godoc
	Caught OutType = 1
	// RunOut godoc
	RunOut OutType = 2
	// Stamped godoc
	Stamped OutType = 3
	// HitWicket godoc
	HitWicket OutType = 3
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

// SeriesState godoc
// @Summary Define different state of series
type SeriesState int

const (
	// NotStarted godoc
	NotStarted SeriesState = 0
	// OnGoing godoc
	OnGoing SeriesState = 1
	// Finished godoc
	Finished SeriesState = 2
)

// ResultType godoc
// @Summary Define different types of game result
type ResultType int

const (
	// Completed godoc
	Completed ResultType = 0
	// Abandoned godoc
	Abandoned ResultType = 1
	// Drawn godoc
	Drawn ResultType = 2
)

// WinLoseType godoc
// @Summary Define different types of win lose type
type WinLoseType int

const (
	// ByRun godoc
	ByRun WinLoseType = 0
	// ByWicket godoc
	ByWicket WinLoseType = 1
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
