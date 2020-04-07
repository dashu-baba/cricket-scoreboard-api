//Package services defines the business logics.
package services

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GameService defines the service instance
type GameService struct {
	SeriesRepository *repositories.SeriesRepository
	TeamRepository   *repositories.TeamRepository
	MatchRepository  *repositories.MatchRepository
	PlayerRepository *repositories.PlayerRepository
}

//NewGameService returns a new GameService.
func NewGameService(seriesRepository *repositories.SeriesRepository,
	teamRepository *repositories.TeamRepository,
	matchRepository *repositories.MatchRepository,
	playerRepository *repositories.PlayerRepository) *GameService {
	return &GameService{
		SeriesRepository: seriesRepository,
		TeamRepository:   teamRepository,
		PlayerRepository: playerRepository,
		MatchRepository:  matchRepository,
	}
}

//CreateSeries crates a series
func (service *GameService) CreateSeries(ctx context.Context, model requestmodels.SeriesCreateModel) (responsemodels.Series, responsemodels.ErrorModel) {
	series := domains.Series{
		Name:     model.Name,
		GameType: model.GameType,
	}

	teamIds := []string{}
	playerIds := []string{}
	if len(model.Teams) > 0 {
		for _, val := range model.Teams {
			teamIds = append(teamIds, val.TeamID)
			playerIds = append(playerIds, val.SquadPlayers...)
		}
	}

	teams := service.TeamRepository.GetAllByIds(ctx, teamIds)
	players := service.PlayerRepository.GetAllByIds(ctx, playerIds)

	if len(teams) != len(teamIds) {
		return responsemodels.Series{}, responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "Team not found.",
		}
	}

	if len(playerIds) != len(players) {
		return responsemodels.Series{}, responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "Player not found.",
		}
	}

	for i := range model.Teams {
		oid, err := primitive.ObjectIDFromHex(model.Teams[i].TeamID)
		if err != nil {
			panic(err)
		}

		participant := domains.SeriesParticipant{
			TeamID: oid,
		}

		for j := range model.Teams[i].SquadPlayers {
			pOid, err := primitive.ObjectIDFromHex(model.Teams[i].SquadPlayers[j])
			if err != nil {
				panic(err)
			}

			participant.SquadPlayers = append(participant.SquadPlayers, pOid)
		}

		series.Teams = append(series.Teams, participant)
	}

	series = service.SeriesRepository.Insert(ctx, series)

	res := responsemodels.Series{
		ID:     series.ID.Hex(),
		Type:   series.GameType,
		Name:   series.Name,
		Teams:  []responsemodels.Team{},
		Status: series.Status,
	}

	if len(series.Teams) > 0 {
		for _, val := range teams {
			team := responsemodels.Team{
				ID:   val.ID.Hex(),
				Name: val.Name,
				Logo: val.Logo,
			}
			res.Teams = append(res.Teams, team)
		}
	}

	return res, responsemodels.ErrorModel{}
}

//GetSeries returns the team by id
func (service *GameService) GetSeries(ctx context.Context, id string) responsemodels.Series {

	series := service.SeriesRepository.GetByID(ctx, id)
	res := responsemodels.Series{
		ID:     series.ID.Hex(),
		Type:   series.GameType,
		Name:   series.Name,
		Teams:  []responsemodels.Team{},
		Status: series.Status,
	}

	if len(series.Teams) > 0 {
		oids := []primitive.ObjectID{}
		for i := range series.Teams {
			oids = append(oids, series.Teams[i].TeamID)
		}

		teams := service.TeamRepository.GetAllByObjIds(ctx, oids)
		for _, val := range teams {
			team := responsemodels.Team{
				ID:   val.ID.Hex(),
				Name: val.Name,
				Logo: val.Logo,
			}
			res.Teams = append(res.Teams, team)
		}
	}

	return res
}

//AddTeam add teams into series
func (service *GameService) AddTeam(ctx context.Context, id string, model requestmodels.TeamsAddModel) (responsemodels.Series, responsemodels.ErrorModel) {

	series := service.SeriesRepository.GetByID(ctx, id)

	if series.GameType == models.Bilateral && (len(model.Teams)+len(series.Teams)) > 2 {
		return responsemodels.Series{}, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Invalid data, bilateral series only have max 2 teams",
		}
	}

	if series.Status != models.NotStarted {
		return responsemodels.Series{}, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modified series",
		}
	}

	if len(model.Teams) > 0 {
		teamIds := []string{}
		playerIds := []string{}
		if len(model.Teams) > 0 {
			for _, val := range model.Teams {
				teamIds = append(teamIds, val.TeamID)
				playerIds = append(playerIds, val.SquadPlayers...)
			}
		}

		teams := service.TeamRepository.GetAllByIds(ctx, teamIds)
		players := service.PlayerRepository.GetAllByIds(ctx, playerIds)

		if len(teams) != len(teamIds) {
			return responsemodels.Series{}, responsemodels.ErrorModel{
				ErrorCode: http.StatusNotFound,
				Message:   "Team not found.",
			}
		}

		if len(playerIds) != len(players) {
			return responsemodels.Series{}, responsemodels.ErrorModel{
				ErrorCode: http.StatusNotFound,
				Message:   "Player not found.",
			}
		}

		for i := range model.Teams {
			oid, err := primitive.ObjectIDFromHex(model.Teams[i].TeamID)
			if err != nil {
				panic(err)
			}

			participant := domains.SeriesParticipant{
				TeamID: oid,
			}

			for j := range model.Teams[i].SquadPlayers {
				pOid, err := primitive.ObjectIDFromHex(model.Teams[i].SquadPlayers[j])
				if err != nil {
					panic(err)
				}

				participant.SquadPlayers = append(participant.SquadPlayers, pOid)
			}

			series.Teams = append(series.Teams, participant)
		}

		updates := map[string]interface{}{}
		updates["teams"] = series.Teams

		service.SeriesRepository.Update(ctx, series.ID.Hex(), updates)
	}

	return service.GetSeries(ctx, id), responsemodels.ErrorModel{}
}

//RemoveTeam add teams into series
func (service *GameService) RemoveTeam(ctx context.Context, id string, model requestmodels.TeamsRemoveModel) (responsemodels.Series, responsemodels.ErrorModel) {

	series := service.SeriesRepository.GetByID(ctx, id)

	if series.Status != models.NotStarted {
		return responsemodels.Series{}, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modified series",
		}
	}

	if len(model.Teams) > 0 {
		teams := []primitive.ObjectID{}
		for _, val := range model.Teams {
			team, err := primitive.ObjectIDFromHex(val)
			if err != nil {
				panic(err)
			}

			teams = append(teams, team)
		}

		for i := 0; i < len(series.Teams); i++ {
			element := series.Teams[i].TeamID
			length1 := len(teams)
			for j := 0; j < length1; j++ {
				if teams[j] == element {
					series.Teams = append(series.Teams[:i], series.Teams[i+1:]...)
					i--
					break
				}
			}
		}

		updates := map[string]interface{}{}
		updates["teams"] = series.Teams

		service.SeriesRepository.Update(ctx, series.ID.Hex(), updates)
	}

	return service.GetSeries(ctx, id), responsemodels.ErrorModel{}
}

//CreateMatches godoc
// @Summary This method create collection of matches under a series.
func (service *GameService) CreateMatches(ctx context.Context, id string, model requestmodels.MatchCreateModel) responsemodels.ErrorModel {
	max := service.MatchRepository.GetLastMatchNumber(ctx)
	series := service.SeriesRepository.GetByID(ctx, id)

	if series.Status != models.NotStarted {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modified series",
		}
	}

	allTeams := []string{}
	for i := range model.Matches {
		allTeams = append(allTeams, model.Matches[i].Participants...)
	}

	for i := range allTeams {
		isExists := false
		for j := range series.Teams {
			if allTeams[i] == series.Teams[j].TeamID.Hex() {
				isExists = true
				break
			}
		}
		if !isExists {
			return responsemodels.ErrorModel{
				ErrorCode: http.StatusNotFound,
				Message:   "Team does not exists in the series",
			}
		}
	}

	teams := service.TeamRepository.GetAllByIds(ctx, allTeams)
	if len(allTeams) != len(teams) {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "Team does not exists in the series",
		}
	}

	matches := []domains.Match{}
	if len(model.Matches) > 0 {
		for _, val := range model.Matches {
			max++
			match := domains.Match{
				ID:           primitive.NewObjectID(),
				MatchType:    val.MatchType,
				Number:       max,
				OverLimit:    val.OverLimit,
				SeriesID:     series.ID,
				Participants: []primitive.ObjectID{},
			}

			if len(val.Participants) > 0 {
				teamIds := []primitive.ObjectID{}
				for _, p := range val.Participants {
					teamID, err := primitive.ObjectIDFromHex(p)
					if err != nil {
						panic(err)
					}

					teamIds = append(teamIds, teamID)
				}
				match.Participants = teamIds
			}

			matches = append(matches, match)
		}

	}

	service.MatchRepository.InsertMany(ctx, matches)

	return responsemodels.ErrorModel{}
}

//UpdateSquad godoc
// @Summary update squad updates squad of a team by adding or removing player
func (service *GameService) UpdateSquad(ctx context.Context, id string, teamID string, model requestmodels.UpdateSquadModel) responsemodels.ErrorModel {

	series := service.SeriesRepository.GetByID(ctx, id)

	if series.Status != models.NotStarted {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modified series",
		}
	}

	players := service.PlayerRepository.GetAllByIds(ctx, model.AddedPlayer)

	if len(model.AddedPlayer) > 0 && len(model.AddedPlayer) != len(players) {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "Team does not exists in the series",
		}
	}

	teamIndex := 0
	for i, val := range series.Teams {
		if val.TeamID.Hex() == teamID {
			teamIndex = i
		}
	}

	if len(model.RemovedPlayer) > 0 {
		exists := false
		for _, val := range model.RemovedPlayer {
			for i := 0; i < len(series.Teams[teamIndex].SquadPlayers); i++ {
				if val == series.Teams[teamIndex].SquadPlayers[i].Hex() {
					series.Teams[teamIndex].SquadPlayers = append(series.Teams[teamIndex].SquadPlayers[:i], series.Teams[teamIndex].SquadPlayers[i+1:]...)
					i--
					exists = true
					break
				}
			}

			if !exists {
				return responsemodels.ErrorModel{
					ErrorCode: http.StatusNotFound,
					Message:   "Player does not exists in the squad to remove",
				}
			}
		}
	}

	if len(model.AddedPlayer) > 0 {
		for _, val := range model.AddedPlayer {
			oid, err := primitive.ObjectIDFromHex(val)
			if err != nil {
				panic(err)
			}

			series.Teams[teamIndex].SquadPlayers = append(series.Teams[teamIndex].SquadPlayers, oid)
		}
	}

	updates := map[string]interface{}{}
	updates["teams"] = series.Teams

	service.SeriesRepository.Update(ctx, series.ID.Hex(), updates)

	return responsemodels.ErrorModel{}
}

//UpdateSeriesStatus godoc
// @Summary update series status
func (service *GameService) UpdateSeriesStatus(ctx context.Context, id string, model requestmodels.UpdateSeriesStatusModel) responsemodels.ErrorModel {

	series := service.SeriesRepository.GetByID(ctx, id)

	if series.Status > model.Status {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Can't downgrade the series status",
		}
	}

	updates := map[string]interface{}{}
	updates["status"] = model.Status

	service.SeriesRepository.Update(ctx, series.ID.Hex(), updates)

	return responsemodels.ErrorModel{}
}
