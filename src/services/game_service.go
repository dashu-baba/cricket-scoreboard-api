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
	SeriesRepository  *repositories.SeriesRepository
	TeamRepository    *repositories.TeamRepository
	MatchRepository   *repositories.MatchRepository
	PlayerRepository  *repositories.PlayerRepository
	InningsRepository *repositories.InningsRepository
	BattingRepository *repositories.BattingRepository
	BowlingRepository *repositories.BowlingRepository
	OverRepository    *repositories.OverRepository
}

//NewGameService returns a new GameService.
func NewGameService(seriesRepository *repositories.SeriesRepository,
	teamRepository *repositories.TeamRepository,
	matchRepository *repositories.MatchRepository,
	playerRepository *repositories.PlayerRepository,
	inningsRepository *repositories.InningsRepository,
	battingRepository *repositories.BattingRepository,
	bowlingRepository *repositories.BowlingRepository,
	overRepository *repositories.OverRepository) *GameService {
	return &GameService{
		SeriesRepository:  seriesRepository,
		TeamRepository:    teamRepository,
		PlayerRepository:  playerRepository,
		MatchRepository:   matchRepository,
		BattingRepository: battingRepository,
		InningsRepository: inningsRepository,
		OverRepository:    overRepository,
		BowlingRepository: bowlingRepository,
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
			teamIds := []primitive.ObjectID{}
			if len(val.Participants) > 0 {
				for _, p := range val.Participants {
					teamID, err := primitive.ObjectIDFromHex(p)
					if err != nil {
						panic(err)
					}

					teamIds = append(teamIds, teamID)
				}
			}

			max++
			match := domains.Match{
				ID:          primitive.NewObjectID(),
				MatchType:   val.MatchType,
				Number:      max,
				OverLimit:   val.OverLimit,
				SeriesID:    series.ID,
				MatchStatus: models.NotStarted,
				Team1: domains.MatchParticipant{
					TeamID: teamIds[0],
				},
				Team2: domains.MatchParticipant{
					TeamID: teamIds[1],
				},
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

//UpdateMatchStatus godoc
// @Summary update match status
func (service *GameService) UpdateMatchStatus(ctx context.Context, id string, model requestmodels.UpdateSeriesStatusModel) responsemodels.ErrorModel {

	match := service.MatchRepository.GetByID(ctx, id)

	if match.ID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The match you are tried to modified is not exists",
		}
	}

	if match.MatchStatus > model.Status {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Can't downgrade the series status",
		}
	}

	updates := map[string]interface{}{}
	updates["matchstatus"] = model.Status

	service.MatchRepository.Update(ctx, match.ID.Hex(), updates)

	return responsemodels.ErrorModel{}
}

//UpdateMatchPlayingSquad godoc
// @Summary update match playing squad updates squad of a team by adding players
func (service *GameService) UpdateMatchPlayingSquad(ctx context.Context, id string, matchID string, model requestmodels.MatchPlayingSquadModel) responsemodels.ErrorModel {

	match := service.MatchRepository.GetByID(ctx, matchID)
	series := service.SeriesRepository.GetByID(ctx, id)

	if series.ID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The series does not exists",
		}
	}

	if match.ID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The match you are tried to modified is not exists",
		}
	}

	if series.ID != match.SeriesID {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "The match id does not belong in the series",
		}
	}

	// if match.Team1.TeamID.Hex() == model.TeamID || match.Team2.TeamID.Hex() == model.TeamID {
	// 	return responsemodels.ErrorModel{
	// 		ErrorCode: http.StatusBadRequest,
	// 		Message:   "The the team info already updated",
	// 	}
	// }

	if match.MatchStatus != models.NotStarted {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modified match",
		}
	}

	idList := append(model.Players, model.Extras...)

	players := service.PlayerRepository.GetAllByIds(ctx, idList)

	if len(idList) != len(players) {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Players does not exists",
		}
	}

	var teamSquad = domains.SeriesParticipant{}
	for _, team := range series.Teams {
		if team.TeamID.Hex() == model.TeamID {
			teamSquad = team
			break
		}
	}

	if teamSquad.TeamID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "This team does not belong to the series",
		}
	}

	for _, p := range idList {
		var exists = false
		for _, p1 := range teamSquad.SquadPlayers {
			if p == p1.Hex() {
				exists = true
				break
			}
		}

		if !exists {
			return responsemodels.ErrorModel{
				ErrorCode: http.StatusBadRequest,
				Message:   "Players does not included in the team squad",
			}
		}
	}

	teamID, err := primitive.ObjectIDFromHex(model.TeamID)
	if err != nil {
		panic(err)
	}
	playerIds := []primitive.ObjectID{}
	extraIds := []primitive.ObjectID{}

	for _, val := range model.Players {
		playerID, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		playerIds = append(playerIds, playerID)
	}

	for _, val := range model.Extras {
		extraID, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		extraIds = append(extraIds, extraID)
	}

	matchTeam := domains.MatchParticipant{
		TeamID:       teamID,
		PlayingSquad: playerIds,
		Extras:       extraIds,
	}

	updates := map[string]interface{}{}

	if match.Team1.TeamID == matchTeam.TeamID {
		updates["team1"] = matchTeam
	} else {
		updates["team2"] = matchTeam
	}

	service.MatchRepository.Update(ctx, match.ID.Hex(), updates)

	return responsemodels.ErrorModel{}
}

//CreateInnings godoc
// @Summary create a new innings
func (service *GameService) CreateInnings(ctx context.Context, id string, matchID string, model requestmodels.CreateInningsModel) (string, responsemodels.ErrorModel) {

	match := service.MatchRepository.GetByID(ctx, matchID)
	series := service.SeriesRepository.GetByID(ctx, id)

	if series.ID.String() == "" {
		return "", responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The series does not exists",
		}
	}

	if match.ID.String() == "" {
		return "", responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The match you are tried to modified is not exists",
		}
	}

	if series.ID != match.SeriesID {
		return "", responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "The match id does not belong in the series",
		}
	}

	if match.MatchStatus != models.NotStarted {
		return "", responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modified match",
		}
	}

	idList := []string{}
	idList = append(idList, model.BattingTeamID, model.BowlingTeamID)

	players := service.PlayerRepository.GetAllByIds(ctx, idList)

	if len(idList) != len(players) {
		return "", responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Players does not exists",
		}
	}

	if (match.Team1.TeamID.Hex() != model.BattingTeamID && match.Team1.TeamID.Hex() == model.BowlingTeamID) ||
		(match.Team2.TeamID.Hex() != model.BattingTeamID && match.Team2.TeamID.Hex() == model.BowlingTeamID) {
		return "", responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "This team does not belong to the match",
		}
	}

	battingTeamID, err := primitive.ObjectIDFromHex(model.BattingTeamID)
	if err != nil {
		panic(err)
	}

	bowlingTeamID, err := primitive.ObjectIDFromHex(model.BowlingTeamID)
	if err != nil {
		panic(err)
	}

	tossWinningTeamID, err := primitive.ObjectIDFromHex(model.TossWinningTeamID)
	if err != nil {
		panic(err)
	}

	number := service.InningsRepository.GetLastInningsNumber(ctx)
	wicketLimit := 0
	if match.Team1.TeamID == battingTeamID {
		wicketLimit = len(match.Team1.PlayingSquad)
	} else {
		wicketLimit = len(match.Team2.PlayingSquad)
	}

	innings := domains.Innings{
		BattingTeamID: battingTeamID,
		BowlingTeamID: bowlingTeamID,
		ID:            primitive.NewObjectID(),
		MatchID:       match.ID,
		Number:        number + 1,
		OverLimit:     match.OverLimit,
		TossResult:    tossWinningTeamID,
		InningsStatus: models.NotStarted,
		Run:           0,
		Wicket:        0,
		WicketLimit:   wicketLimit,
	}

	list := []domains.Innings{}
	list = append(list, innings)

	service.InningsRepository.InsertMany(ctx, list)

	return innings.ID.Hex(), responsemodels.ErrorModel{}
}
