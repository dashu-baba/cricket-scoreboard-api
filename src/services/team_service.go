//Package services defines the repository items.
package services

import (
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"
)

//TeamService defines the service instance
type TeamService struct {
	TeamRepository   *repositories.TeamRepository
	PlayerRepository *repositories.PlayerRepository
}

//NewTeamService returns a new TeamService.
func NewTeamService(TeamRepository *repositories.TeamRepository,
	PlayerRepository *repositories.PlayerRepository) *TeamService {
	return &TeamService{
		TeamRepository:   TeamRepository,
		PlayerRepository: PlayerRepository,
	}
}

//GetAllTeam returns the collection of all team
func (service *TeamService) GetAllTeam() []responsemodels.Team {

	teams := service.TeamRepository.GetAll()
	responses := []responsemodels.Team{}
	for _, team := range teams {
		response := responsemodels.Team{
			ID:   team.ID.String(),
			Logo: team.Logo,
			Name: team.Name,
		}
		responses = append(responses, response)
	}

	return responses
}

//CreateTeam insert a team item
func (service *TeamService) CreateTeam(model requestmodels.TeamCreateModel) {
	team := domains.Team{
		Logo: model.Logo,
		Name: model.Name,
	}

	team = service.TeamRepository.Insert(team)
	players := []domains.Player{}
	for _, val := range model.Players {
		player := domains.Player{
			Name:       val.Name,
			PlayerType: val.PlayerType,
			TeamID:     team.ID,
		}
		players = append(players, player)
	}

	players = service.PlayerRepository.InsertMany(players)

	service.TeamRepository.Update(team, players)
}
