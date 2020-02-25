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

const teamCollectionName = "teams"

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
	collections := repo.DB.Database.Collection(teamCollectionName)
	team.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, team)

	if err != nil {
		panic(err)
	}

	return team
}

//Update update a team object into db
//and return that updated item.
func (repo *TeamRepository) Update(team domains.Team, players []domains.Player) domains.Team {
	collections := repo.DB.Database.Collection(teamCollectionName)

	filter := bson.M{"id": team.ID}
	update := bson.M{"$set": bson.M{"players": players}}
	_, err := collections.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		panic(err)
	}

	return team
}

//toDoc converts object to bson document
func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

//GetAll retrieves all team objects from db
//and return that collection.
func (repo *TeamRepository) GetAll() []domains.Team {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection(teamCollectionName)
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

//GetAllByIds retrieves team objects from db by ids
//and return that collection.
func (repo *TeamRepository) GetAllByIds(ids []string) []domains.Team {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection(teamCollectionName)

	oids := []primitive.ObjectID{}
	for _, val := range ids {
		oid, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		oids = append(oids, oid)
	}
	cursor, err := collections.Find(ctx, bson.M{"id": bson.M{"in": oids}})

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
