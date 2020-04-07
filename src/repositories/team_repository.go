//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"fmt"

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
func (repo *TeamRepository) Insert(ctx context.Context, team domains.Team) domains.Team {
	collections := repo.DB.Database.Collection(teamCollectionName)
	team.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, team)

	if err != nil {
		panic(err)
	}

	return team
}

//Update update a player object into db
//and return that updated item.
func (repo *TeamRepository) Update(ctx context.Context, id string, updates map[string]interface{}) domains.Team {
	collections := repo.DB.Database.Collection(teamCollectionName)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"id": objID}
	updatedValue := bson.M{}
	updatedValue = updates

	update := bson.M{"$set": updatedValue}
	updateResult, err := collections.UpdateOne(
		ctx,
		filter,
		update,
	)

	fmt.Println(updateResult)

	if err != nil {
		panic(err)
	}

	return repo.GetByID(ctx, id)
}

//GetAll retrieves all team objects from db
//and return that collection.
func (repo *TeamRepository) GetAll(ctx context.Context) []domains.Team {
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
func (repo *TeamRepository) GetAllByIds(ctx context.Context, ids []string) []domains.Team {
	oids := []primitive.ObjectID{}
	for _, val := range ids {
		oid, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		oids = append(oids, oid)
	}

	return repo.GetAllByObjIds(ctx, oids)
}

//GetAllByObjIds retrieves team objects from db by ids
//and return that collection.
func (repo *TeamRepository) GetAllByObjIds(ctx context.Context, oids []primitive.ObjectID) []domains.Team {
	collections := repo.DB.Database.Collection(teamCollectionName)

	cursor, err := collections.Find(ctx, bson.M{"id": bson.M{"$in": oids}})

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

//GetByID retrieves team object from db by id
//and return that object.
func (repo *TeamRepository) GetByID(ctx context.Context, id string) domains.Team {
	collections := repo.DB.Database.Collection(teamCollectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"id": objID})

	if err := findResult.Err(); err != nil {
		panic(err)
	}

	team := domains.Team{}
	err = findResult.Decode(&team)
	if err != nil {
		panic(err)
	}

	return team
}
