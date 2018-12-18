package mongo_test

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"log"
	"testing"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "testDb"
	userCollectionName = "user"
)

func TestGfUserService(t *testing.T) {
	t.Run("CreateGfUser", createGfUser_should_insert_user_into_mongo)
}

func createGfUser_should_insert_user_into_mongo(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer finishTest(session)
	userService := mongo.NewGfUserService(session.Copy(), dbName, userCollectionName)

	testId, testUsername := "1111", "super_username"
	user := entity.GfUser{
		ID:       testId,
		Username: testUsername,
	}

	err = userService.CreateUser(&user)

	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}
	results := make([]entity.GfUser, 0)
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}
	if results[0].Username != user.Username {
		t.Errorf("Wrong username. Expected %s, got %s", testUsername, results[0].Username)
	}
}

func finishTest(s *mongo.Session) {
	s.DropDatabase(dbName)
	s.Close()
}
