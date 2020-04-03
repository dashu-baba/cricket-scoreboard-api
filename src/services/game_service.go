//Package services defines the business logics.
package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/responsemodels"
	"cricket-scoreboard-api/src/requestmodels"
	"context"
	"cricket-scoreboard-api/src/repositories"
)

//GameService defines the service instance
type GameService struct {
	SeriesRepository   *repositories.SeriesRepository
	TeamRepository   *repositories.TeamRepository
}

//NewGameService returns a new GameService.
func NewGameService(seriesRepository *repositories.SeriesRepository,
	teamRepository   *repositories.TeamRepository) *GameService {
	return &GameService{
		SeriesRepository:   seriesRepository,
		TeamRepository: teamRepository,
	}
}

//CreateSeries crates a series
func (service *GameService) CreateSeries(ctx context.Context, model requestmodels.SeriesCreateModel) responsemodels.Series {
	series := domains.Series{
		Name: model.Name,
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
		ID : series.ID.Hex(),
		Type : series.GameType,
		Name : series.Name,
		Teams : []responsemodels.Team{},
	}

	if len(series.Teams) > 0 {

		teams := service.TeamRepository.GetAllByObjIds(ctx, series.Teams)
		for _, val := range teams {
			team := responsemodels.Team{
				ID : 	val.ID.Hex(),
				Name:   val.Name,
				Logo:   val.Logo,
			}
			res.Teams = append(res.Teams, team)
		}
	}

	return res;
}

//GetSeries returns the team by id
func (service *GameService) GetSeries(ctx context.Context, id string) responsemodels.Series {

	series := service.SeriesRepository.GetByID(ctx, id)
	res := responsemodels.Series{
		ID : series.ID.Hex(),
		Type : series.GameType,
		Name : series.Name,
		Teams : []responsemodels.Team{},
	}

	if len(series.Teams) > 0 {

		teams := service.TeamRepository.GetAllByObjIds(ctx, series.Teams)
		for _, val := range teams {
			team := responsemodels.Team{
				ID : 	val.ID.Hex(),
				Name:   val.Name,
				Logo:   val.Logo,
			}
			res.Teams = append(res.Teams, team)
		}
	}

	return res
}
