//Package services defines the repository items.
package services

import (
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

//CreatePlayer insert a player item
func (service *TeamService) CreatePlayer(model requestmodels.PlayerCreateModel) responsemodels.Player {
	teamID, err := primitive.ObjectIDFromHex(model.TeamID)
	if err != nil {
		panic(err)
	}

	player := domains.Player{
		Name:       model.Name,
		PlayerType: model.PlayerType,
		TeamID:     teamID,
	}

	player = service.PlayerRepository.Insert(player)

	team := service.TeamRepository.GetByID(model.TeamID)
	team.Players = append(team.Players, player)
	service.TeamRepository.Update(team, team.Players)

	return responsemodels.Player{
		ID:         player.ID.String(),
		Name:       player.Name,
		PlayerType: player.PlayerType,
		TeamID:     player.TeamID.String(),
	}
}

//RemovePlayer remove a player from team
func (service *TeamService) RemovePlayer(teamID string, playerID string) {
	playerObjID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		panic(err)
	}

	service.PlayerRepository.Remove(playerObjID)

	team := service.TeamRepository.GetByID(teamID)

	position := -1

	for index, val := range team.Players {
		if val.ID == playerObjID {
			position = index
			break
		}
	}

	if position != -1 {
		team.Players = append(team.Players[:position], team.Players[:position+1]...)
	}

	service.TeamRepository.Update(team, team.Players)
}
