//Package services defines the repository items.
package services

import (
	"context"
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
func (service *TeamService) GetAllTeam(ctx context.Context) []responsemodels.Team {

	teams := service.TeamRepository.GetAll(ctx)
	responses := []responsemodels.Team{}
	for _, team := range teams {
		response := responsemodels.Team{
			ID:   team.ID.Hex(),
			Logo: team.Logo,
			Name: team.Name,
		}
		response.Players = []responsemodels.Player{}
		for _, player := range team.Players {
			response.Players = append(response.Players, responsemodels.Player{
				ID:         player.ID.Hex(),
				Name:       player.Name,
				PlayerType: player.PlayerType,
				TeamID:     player.TeamID.Hex(),
			})
		}

		responses = append(responses, response)
	}

	return responses
}

//GetTeam returns the team by id
func (service *TeamService) GetTeam(ctx context.Context, id string) responsemodels.Team {

	team := service.TeamRepository.GetByID(ctx, id)
	response := responsemodels.Team{
		ID:   team.ID.Hex(),
		Logo: team.Logo,
		Name: team.Name,
	}
	response.Players = []responsemodels.Player{}
	for _, player := range team.Players {
		responsePlayer := responsemodels.Player{
			ID:         player.ID.Hex(),
			Name:       player.Name,
			PlayerType: player.PlayerType,
			TeamID:     player.TeamID.Hex(),
		}
		response.Players = append(response.Players, responsePlayer)
	}

	return response
}

//CreateTeam insert a team item
func (service *TeamService) CreateTeam(ctx context.Context, model requestmodels.TeamCreateModel) {
	team := domains.Team{
		Logo: model.Logo,
		Name: model.Name,
	}

	team = service.TeamRepository.Insert(ctx, team)
	players := []domains.Player{}
	for _, val := range model.Players {
		player := domains.Player{
			Name:       val.Name,
			PlayerType: val.PlayerType,
			TeamID:     team.ID,
		}
		players = append(players, player)
	}

	players = service.PlayerRepository.InsertMany(ctx, players)

	if len(players) > 0 {
		updates := map[string]interface{}{}
		updates["players"] = players

		service.TeamRepository.Update(ctx, team.ID.Hex(), updates)
	}
}

//UpdateTeam update a team item
func (service *TeamService) UpdateTeam(ctx context.Context, id string, model requestmodels.TeamUpdateModel) {
	updates := map[string]interface{}{}
	updates["name"] = model.Name
	service.TeamRepository.Update(ctx, id, updates)
}

//UpdatePlayer update a player item
func (service *TeamService) UpdatePlayer(ctx context.Context, id string, teamID string, model requestmodels.PlayerUpdateModel) {
	updates := map[string]interface{}{}
	updates["name"] = model.Name
	updates["playertype"] = model.PlayerType
	service.PlayerRepository.Update(ctx, id, updates)

	team := service.TeamRepository.GetByID(ctx, teamID)
	team.Players = service.PlayerRepository.GetAll(ctx, teamID)

	updates = map[string]interface{}{}
	updates["players"] = team.Players

	service.TeamRepository.Update(ctx, team.ID.Hex(), updates)
}

//CreatePlayer insert a player item
func (service *TeamService) CreatePlayer(ctx context.Context, teamID string, model requestmodels.PlayerCreateModel) responsemodels.Player {
	teamObjID, err := primitive.ObjectIDFromHex(teamID)
	if err != nil {
		panic(err)
	}

	player := domains.Player{
		Name:       model.Name,
		PlayerType: model.PlayerType,
		TeamID:     teamObjID,
	}

	player = service.PlayerRepository.Insert(ctx, player)

	team := service.TeamRepository.GetByID(ctx, teamID)
	team.Players = append(team.Players, player)

	updates := map[string]interface{}{}
	updates["players"] = team.Players

	service.TeamRepository.Update(ctx, team.ID.Hex(), updates)

	return responsemodels.Player{
		ID:         player.ID.Hex(),
		Name:       player.Name,
		PlayerType: player.PlayerType,
		TeamID:     player.TeamID.Hex(),
	}
}

//RemovePlayer remove a player from team
func (service *TeamService) RemovePlayer(ctx context.Context, teamID string, playerID string) {
	playerObjID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		panic(err)
	}

	service.PlayerRepository.Remove(ctx, playerObjID)

	team := service.TeamRepository.GetByID(ctx, teamID)

	position := -1

	for index, val := range team.Players {
		if val.ID == playerObjID {
			position = index
			break
		}
	}

	if position != -1 {
		team.Players[position] = team.Players[len(team.Players)-1] // Copy last element to index i.
		// Erase last element (write zero value).
		team.Players = team.Players[:len(team.Players)-1]
	}

	updates := map[string]interface{}{}
	updates["players"] = team.Players

	service.TeamRepository.Update(ctx, team.ID.Hex(), updates)
}
