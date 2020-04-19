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

//InningsService defines the service instance
type InningsService struct {
	SeriesRepository  *repositories.SeriesRepository
	TeamRepository    *repositories.TeamRepository
	MatchRepository   *repositories.MatchRepository
	PlayerRepository  *repositories.PlayerRepository
	InningsRepository *repositories.InningsRepository
	BattingRepository *repositories.BattingRepository
	BowlingRepository *repositories.BowlingRepository
	OverRepository    *repositories.OverRepository
}

//NewInningsService returns a new InningsService.
func NewInningsService(seriesRepository *repositories.SeriesRepository,
	teamRepository *repositories.TeamRepository,
	matchRepository *repositories.MatchRepository,
	playerRepository *repositories.PlayerRepository,
	inningsRepository *repositories.InningsRepository,
	battingRepository *repositories.BattingRepository,
	bowlingRepository *repositories.BowlingRepository,
	overRepository *repositories.OverRepository) *InningsService {
	return &InningsService{
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

//StartInnings godoc
// @Summary start a innings
func (service *InningsService) StartInnings(ctx context.Context, id string, matchID string, inningsID string,
	model requestmodels.StartInningsModel) responsemodels.ErrorModel {

	match := service.MatchRepository.GetByID(ctx, matchID)
	series := service.SeriesRepository.GetByID(ctx, id)
	innings := service.InningsRepository.GetByID(ctx, inningsID)

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

	if innings.ID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The innings you are tried to modified is not exists",
		}
	}

	if series.ID != match.SeriesID {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "The match id does not belong in the series",
		}
	}

	if innings.MatchID != match.ID {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "The innings does not belong in the match",
		}
	}

	if innings.InningsStatus != models.NotStarted {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot modify innings",
		}
	}

	battingTeam := domains.MatchParticipant{}
	bowlingTeam := domains.MatchParticipant{}

	if match.Team1.TeamID == innings.BattingTeamID {
		battingTeam = match.Team1
		bowlingTeam = match.Team2
	} else {
		battingTeam = match.Team2
		bowlingTeam = match.Team1
	}

	exist1 := false
	exist2 := false
	for _, val := range battingTeam.PlayingSquad {
		if exist1 && exist2 {
			break
		}
		if model.StrikeBatsmanID == val.Hex() && exist1 {
			exist1 = true
			continue
		}
		if model.NonStrikeBatsmanID == val.Hex() && exist2 {
			exist2 = true
			continue
		}
	}

	if !(exist1 && exist2) {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Batsman does not included in the squad",
		}
	}

	exist1 = false
	for _, val := range bowlingTeam.PlayingSquad {
		if model.BowlerID == val.Hex() {
			exist1 = true
			break
		}
	}

	if !(exist1) {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Bowler does not included in the squad",
		}
	}

	batsman1ID, err := primitive.ObjectIDFromHex(model.StrikeBatsmanID)
	if err != nil {
		panic(err)
	}

	batsman2ID, err := primitive.ObjectIDFromHex(model.NonStrikeBatsmanID)
	if err != nil {
		panic(err)
	}

	bowlerID, err := primitive.ObjectIDFromHex(model.BowlerID)
	if err != nil {
		panic(err)
	}

	batsman1 := domains.Batting{
		ID:         primitive.NewObjectID(),
		InningsID:  innings.ID,
		IsInCrease: true,
		IsInStrike: true,
		PlayerID:   batsman1ID,
		Ball:       0,
	}

	batsman2 := domains.Batting{
		ID:         primitive.NewObjectID(),
		InningsID:  innings.ID,
		IsInCrease: true,
		IsInStrike: true,
		PlayerID:   batsman2ID,
		Ball:       0,
	}

	bowler := domains.Bowling{
		ID:        primitive.NewObjectID(),
		InningsID: innings.ID,
		PlayerID:  bowlerID,
		IsCurrent: true,
	}

	over := domains.Over{
		ID:         primitive.NewObjectID(),
		InningsID:  innings.ID,
		IsRunning:  true,
		BowlerID:   bowlerID,
		OverNumber: 1,
		Ball:       0,
	}

	list := []domains.Batting{}
	list = append(list, batsman1, batsman2)
	service.BattingRepository.InsertMany(ctx, list)

	bowlerList := []domains.Bowling{}
	bowlerList = append(bowlerList, bowler)
	service.BowlingRepository.InsertMany(ctx, bowlerList)

	overList := []domains.Over{}
	overList = append(overList, over)
	service.OverRepository.InsertMany(ctx, overList)

	updates := map[string]interface{}{}
	updates["inningsstatus"] = models.OnGoing
	service.InningsRepository.Update(ctx, inningsID, updates)

	updates = map[string]interface{}{}
	updates["matchstatus"] = models.OnGoing
	service.MatchRepository.Update(ctx, matchID, updates)

	updates = map[string]interface{}{}
	updates["status"] = models.OnGoing
	service.SeriesRepository.Update(ctx, id, updates)

	return responsemodels.ErrorModel{}
}

//UpdateOver godoc
// @Summary start a innings
func (service *InningsService) UpdateOver(ctx context.Context, inningsID string, overID string,
	model requestmodels.OverUpdateModel) responsemodels.ErrorModel {

	over := service.OverRepository.GetByID(ctx, overID)
	innings := service.InningsRepository.GetByID(ctx, inningsID)
	batsmans := service.BattingRepository.GetCurrentBatsman(ctx, inningsID)
	bowler := service.BowlingRepository.GetCurrentBowler(ctx, inningsID)

	if innings.ID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The innings you are tried to modified is not exists",
		}
	}

	if over.ID.String() == "" {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The over you are tried to modified is not exists",
		}
	}

	if !over.IsRunning {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusNotFound,
			Message:   "The over already finished",
		}
	}

	match := service.MatchRepository.GetByID(ctx, innings.MatchID.Hex())

	if innings.InningsStatus != models.OnGoing {
		return responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   "Cannot innings not started",
		}
	}

	batsman := domains.Batting{}
	nonStrikebatsman := domains.Batting{}
	for _, val := range batsmans {
		if val.IsInStrike {
			batsman = val
		} else {
			nonStrikebatsman = val
		}
	}

	overUpdates := map[string]interface{}{}
	batsmanUpdates := map[string]interface{}{}
	nonStrikeUpdates := map[string]interface{}{}
	bowlerUpdates := map[string]interface{}{}
	inningsUpdates := map[string]interface{}{}

	if model.Wicket != (requestmodels.WicketDetailsModel{}) {
		battingTeam := domains.MatchParticipant{}
		bowlingTeam := domains.MatchParticipant{}

		if match.Team1.TeamID == innings.BattingTeamID {
			battingTeam = match.Team1
			bowlingTeam = match.Team2
		} else {
			battingTeam = match.Team2
			bowlingTeam = match.Team1
		}

		exist1 := false
		exist2 := false
		for _, val := range bowlingTeam.PlayingSquad {
			if exist1 && exist2 {
				break
			}
			if model.Wicket.BowlerID == val.Hex() && exist1 {
				exist1 = true
				continue
			}
			if model.Wicket.SupportID == "" || (model.Wicket.SupportID == val.Hex() && exist2) {
				exist2 = true
				continue
			}
		}

		if !(exist1 && exist2) {
			return responsemodels.ErrorModel{
				ErrorCode: http.StatusBadRequest,
				Message:   "Bowler and/or supporting wicket-taker does not included in the squad",
			}
		}

		exist1 = false
		for _, val := range battingTeam.PlayingSquad {
			if model.Wicket.BatsmanID == val.Hex() {
				exist1 = true
				break
			}
		}

		if !(exist1) {
			return responsemodels.ErrorModel{
				ErrorCode: http.StatusBadRequest,
				Message:   "Batsman does not included in the squad",
			}
		}

		exist1 = false
		for _, val := range batsmans {
			if model.Wicket.BatsmanID == val.PlayerID.Hex() {
				exist1 = true
				break
			}
		}

		if !(exist1) {
			return responsemodels.ErrorModel{
				ErrorCode: http.StatusBadRequest,
				Message:   "Batsman does not exists in the crease",
			}
		}

		batsmanID, err := primitive.ObjectIDFromHex(model.Wicket.BatsmanID)
		if err != nil {
			panic(err)
		}

		bowlerID, err := primitive.ObjectIDFromHex(model.Wicket.BowlerID)
		if err != nil {
			panic(err)
		}

		wicket := domains.Wicket{
			BatsmanID: batsmanID,
			OutType:   model.Wicket.WicketType,
		}

		if model.Wicket.SupportID != "" {
			supportID, err := primitive.ObjectIDFromHex(model.Wicket.SupportID)
			if err != nil {
				panic(err)
			}
			wicket.SupportID = supportID
			batsmanUpdates["supportedby"] = supportID
		}

		over.Wickets = append(over.Wickets, wicket)
		overUpdates["wickets"] = over.Wickets

		if batsman.PlayerID == batsmanID {
			batsmanUpdates["outtype"] = model.Wicket.WicketType
			batsmanUpdates["wicketby"] = bowlerID
			batsmanUpdates["isincrease"] = false
		} else {
			nonStrikeUpdates["outtype"] = model.Wicket.WicketType
			nonStrikeUpdates["wicketby"] = bowlerID
			nonStrikeUpdates["isincrease"] = false
		}

		innings.Wicket++
		inningsUpdates["wicket"] = innings.Wicket
	}

	if model.Extra != "" {
		overUpdates["ball"] = over.Ball + 1
	}

	run := 0
	batsmanUpdates["ball"] = batsman.Ball + 1
	batsmanUpdates, nonStrikeUpdates = ChangeCrease(batsmanUpdates, nonStrikeUpdates, model)

	if model.NB {
		over.Noball++
		overUpdates["noball"] = over.Noball
	}

	if model.Extra == "" {
		overUpdates, batsmanUpdates, run = UpdateRun(&over, overUpdates, &batsman, batsmanUpdates, run, model)
	} else {
		overUpdates, run = UpdateExtra(&over, overUpdates, run, model)
	}

	inningsUpdates["run"] = innings.Run + run

	if over.Ball == 6 {
		overUpdates["isrunning"] = false
	}

	over = service.OverRepository.Update(ctx, overID, overUpdates)

	if !over.IsRunning {
		bowler.Overs = append(bowler.Overs, over)
		bowlerUpdates["over"] = bowler.Overs
		service.BowlingRepository.Update(ctx, bowler.ID.Hex(), bowlerUpdates)
		batsmanUpdates["isinstrike"] = !batsmanUpdates["isinstrike"].(bool)
		nonStrikeUpdates["isinstrike"] = !nonStrikeUpdates["isinstrike"].(bool)
		innings.OverPlayed++
		inningsUpdates["overplayed"] = innings.OverPlayed
	}

	if innings.OverPlayed == innings.OverLimit || innings.WicketLimit == innings.Wicket+1 {
		inningsUpdates["inningsstatus"] = models.Finished
	}

	service.BattingRepository.Update(ctx, batsman.ID.Hex(), batsmanUpdates)
	service.BattingRepository.Update(ctx, nonStrikebatsman.ID.Hex(), nonStrikeUpdates)
	service.InningsRepository.Update(ctx, innings.ID.Hex(), inningsUpdates)

	return responsemodels.ErrorModel{}
}

//UpdateRun godoc
func UpdateRun(over *domains.Over, overUpdate map[string]interface{},
	batsman *domains.Batting, batsmanUpdate map[string]interface{}, run int,
	model requestmodels.OverUpdateModel) (map[string]interface{}, map[string]interface{}, int) {
	switch model.Run {
	case 0:
		{
			over.Zero++
			overUpdate["zero"] = over.Zero
			batsman.Zero++
			batsmanUpdate["zero"] = batsman.Zero
			return overUpdate, batsmanUpdate, run
		}
	case 1:
		{
			over.One++
			overUpdate["one"] = over.One
			batsman.One++
			batsmanUpdate["one"] = batsman.One
			batsmanUpdate["run"] = batsman.Run + 1
			run++
			return overUpdate, batsmanUpdate, run
		}
	case 2:
		{
			over.Two++
			overUpdate["two"] = over.Two
			batsman.Two++
			batsmanUpdate["two"] = batsman.Two
			batsmanUpdate["run"] = batsman.Run + 2
			run += 2
			return overUpdate, batsmanUpdate, run
		}
	case 3:
		{
			over.Three++
			overUpdate["three"] = over.Three
			batsman.Three++
			batsmanUpdate["three"] = batsman.Three
			batsmanUpdate["run"] = batsman.Run + 3
			run += 3
			return overUpdate, batsmanUpdate, run
		}
	case 4:
		{
			over.Four++
			overUpdate["four"] = over.Four
			batsman.Four++
			batsmanUpdate["four"] = batsman.Four
			batsmanUpdate["run"] = batsman.Run + 4
			run += 4
			return overUpdate, batsmanUpdate, run
		}
	case 5:
		{
			over.Five++
			overUpdate["five"] = over.Five
			batsman.Five++
			batsmanUpdate["five"] = batsman.Five
			batsmanUpdate["run"] = batsman.Run + 5
			run += 5
			return overUpdate, batsmanUpdate, run
		}
	case 6:
		{
			over.Six++
			overUpdate["six"] = over.Six
			batsman.Six++
			batsmanUpdate["six"] = batsman.Six
			batsmanUpdate["run"] = batsman.Run + 6
			run += 6
			return overUpdate, batsmanUpdate, run
		}
	}

	return overUpdate, batsmanUpdate, run
}

//ChangeCrease godoc
func ChangeCrease(strikerUpdate map[string]interface{}, nonStrikerUpdate map[string]interface{},
	model requestmodels.OverUpdateModel) (map[string]interface{}, map[string]interface{}) {
	switch model.Run {
	case 1:
	case 3:
	case 5:
		{
			strikerUpdate["isinstrike"] = false
			nonStrikerUpdate["isinstrike"] = true
			return strikerUpdate, nonStrikerUpdate
		}
	}
	return strikerUpdate, nonStrikerUpdate
}

//UpdateExtra godoc
func UpdateExtra(over *domains.Over, overUpdate map[string]interface{},
	run int, model requestmodels.OverUpdateModel) (map[string]interface{}, int) {
	switch model.Extra {
	case "wd":
		{
			over.Wide++
			overUpdate["wide"] = over.Wide
			over.Bye += model.Run
			overUpdate["bye"] = over.Bye
			return overUpdate, run
		}
	case "b":
		{
			over.Bye += model.Run
			overUpdate["bye"] = over.Bye
			return overUpdate, run
		}
	case "lb":
		{
			over.LB += model.Run
			overUpdate["lb"] = over.LB
			return overUpdate, run
		}
	}

	return overUpdate, run
}
