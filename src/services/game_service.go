//Package services defines the business logics.
package services

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GameService defines the service instance
type GameService struct {
	SeriesRepository *repositories.SeriesRepository
	TeamRepository   *repositories.TeamRepository
}

//NewGameService returns a new GameService.
func NewGameService(seriesRepository *repositories.SeriesRepository,
	teamRepository *repositories.TeamRepository) *GameService {
	return &GameService{
		SeriesRepository: seriesRepository,
		TeamRepository:   teamRepository,
	}
}

//CreateSeries crates a series
func (service *GameService) CreateSeries(ctx context.Context, model requestmodels.SeriesCreateModel) responsemodels.Series {
	series := domains.Series{
		Name:     model.Name,
		GameType: model.GameType,
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
	}

	series = service.SeriesRepository.Insert(ctx, series)

	res := responsemodels.Series{
		ID:    series.ID.Hex(),
		Type:  series.GameType,
		Name:  series.Name,
		Teams: []responsemodels.Team{},
	}

	if len(series.Teams) > 0 {

		teams := service.TeamRepository.GetAllByObjIds(ctx, series.Teams)
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

//GetSeries returns the team by id
func (service *GameService) GetSeries(ctx context.Context, id string) responsemodels.Series {

	series := service.SeriesRepository.GetByID(ctx, id)
	res := responsemodels.Series{
		ID:    series.ID.Hex(),
		Type:  series.GameType,
		Name:  series.Name,
		Teams: []responsemodels.Team{},
	}

	if len(series.Teams) > 0 {

		teams := service.TeamRepository.GetAllByObjIds(ctx, series.Teams)
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
func (service *GameService) AddTeam(ctx context.Context, id string, model requestmodels.TeamsAddRemoveModel) (responsemodels.Series, error) {

	series := service.SeriesRepository.GetByID(ctx, id)

	if series.GameType == models.Bilateral && (len(model.Teams)+len(series.Teams)) > 2 {
		return responsemodels.Series{}, errors.New("Invalid data, bilateral series only have max 2 teams")
	}

	if len(model.Teams) > 0 {
		for _, val := range model.Teams {
			team, err := primitive.ObjectIDFromHex(val)
			if err != nil {
				panic(err)
			}

			series.Teams = append(series.Teams, team)
		}

		updates := map[string]interface{}{}
		updates["teams"] = series.Teams

		service.SeriesRepository.Update(ctx, series.ID.Hex(), updates)
	}

	return service.GetSeries(ctx, id), nil
}

//RemoveTeam add teams into series
func (service *GameService) RemoveTeam(ctx context.Context, id string, model requestmodels.TeamsAddRemoveModel) (responsemodels.Series, error) {

	series := service.SeriesRepository.GetByID(ctx, id)

	if len(model.Teams) > 0 {
		teams := []primitive.ObjectID{}
		for _, val := range model.Teams {
			team, err := primitive.ObjectIDFromHex(val)
			if err != nil {
				panic(err)
			}

			teams = append(teams, team)
		}

		length := len(series.Teams)
		for i := 0; i < length; i++ {
			element := series.Teams[i]
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

	return service.GetSeries(ctx, id), nil
}
