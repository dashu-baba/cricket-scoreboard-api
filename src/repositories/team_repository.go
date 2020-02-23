//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TeamRepository defines the instance
type TeamRepository struct {
	DB *driver.DB
}

//NewTeamRepository returns a new TeamRepository.
func NewTeamRepository(DB *driver.DB) *TeamRepository {
	return &TeamRepository{
		DB: DB,
	}
}

//Insert insert a team object into db
//and return that inserted item.
func (repo *TeamRepository) Insert(team domains.Team) domains.Team {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection("teams")
	team.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, team)

	if err != nil {
		panic(err)
	}

	return team
}

//GetAll retrieves all team objects from db
//and return that collection.
func (repo *TeamRepository) GetAll() []domains.Team {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	collections := repo.DB.Database.Collection("teams")
	cursor, err := collections.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	teams := []domains.Team{}
	for cursor.Next(ctx) {
		team := domains.Team{}
		cursor.Decode(&team)
		teams = append(teams, team)
	}

	return teams
}
