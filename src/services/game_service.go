//Package services defines the business logics.
package services

import (
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/requestmodels"
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

//CreateGame creates a game item
func (service *GameService) CreateGame(model requestmodels.GameCreateModel) {
	game := domains.Game{
		GameType: model.GameType,
		Teams:    service.TeamRepository.GetAllByIds(model.Teams),
		Name:     model.Name,
	}

	game = service.GameRepository.Insert(game)
}
