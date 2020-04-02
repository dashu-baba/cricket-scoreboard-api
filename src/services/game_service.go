//Package services defines the business logics.
package services

import (
	"cricket-scoreboard-api/src/repositories"
)

//GameService defines the service instance
type GameService struct {
	TeamRepository   *repositories.TeamRepository
	PlayerRepository *repositories.PlayerRepository
	GameRepository   *repositories.GameRepository
}

//NewGameService returns a new GameService.
func NewGameService(teamRepository *repositories.TeamRepository,
	playerRepository *repositories.PlayerRepository,
	gameRepository *repositories.GameRepository) *GameService {
	return &GameService{
		TeamRepository:   teamRepository,
		PlayerRepository: playerRepository,
		GameRepository:   gameRepository,
	}
}
