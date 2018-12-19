package mongo_test

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/service"
	"log"
	"testing"
)

const (
	mongoUrl       = "localhost:27017"
	dbName         = "testDb"
	collectionName = "user"
)

func TestServices(t *testing.T) {
	t.Run("UserService", userService)
	t.Run("UsageService", usageService)
}

func usageService(t *testing.T) {
	t.Run("CreateUsage", createUsage_should_insert_usage_into_mongo)
}

func userService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer finishTest(session)
	userService := service.NewUserService(session.Copy(), dbName, collectionName)

	testId, testUsername := "1111", "super_username"
	user := entity.User{
		ID:       testId,
		Username: testUsername,
	}

	err = userService.CreateUser(&user)
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	results := make([]entity.User, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].Username != user.Username {
		t.Errorf("Wrong username. Expected %s, got %s", testUsername, results[0].Username)
	}
}

func createUsage_should_insert_usage_into_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	usageService := service.NewUsageService(session.Copy(), dbName, collectionName)

	testId, testUsageID, testUsageName := "1111", "this_usageID", "this_usageName"
	usage := entity.Usage{
		ID:        testId,
		UsageID:   testUsageID,
		UsageName: testUsageName,
	}

	err := usageService.CreateUsage(&usage)
	if err != nil {
		t.Errorf("Unable to create usage: %s", err)
	}

	results := make([]entity.Usage, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].UsageID != usage.UsageID {
		t.Errorf("Wrong usageID. Expected %s, got %s", testUsageID, results[0].UsageID)
	}
	if results[0].UsageName != usage.UsageName {
		t.Errorf("Wrong UsageName. Expected %s, got %s", testUsageName, results[0].UsageName)
	}

}

func connect() *mongo.Session {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo %s", err)
	}
	return session
}

func finishTest(s *mongo.Session) {
	s.DropDatabase(dbName)
	s.Close()
}
