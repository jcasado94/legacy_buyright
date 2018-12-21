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
	collectionName = "testCol"
)

func TestServices(t *testing.T) {
	t.Run("IUserService", iUserService)
	t.Run("UsageService", usageService)
}

func iUserService(t *testing.T) {
	t.Run("CreateIUser", createIUser_should_insert_user_into_mongo)
}

func createIUser_should_insert_user_into_mongo(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer finishTest(session)
	iUserService := service.NewIUserService(session.Copy(), dbName, collectionName)

	testId, testUsername := "1111", "super_username"
	user := entity.IUser{
		ID:       testId,
		Username: testUsername,
	}

	err = iUserService.CreateUser(&user)
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	results := make([]entity.IUser, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].Username != user.Username {
		t.Errorf("Wrong username. Expected %s, got %s", testUsername, results[0].Username)
	}
}

func usageService(t *testing.T) {
	t.Run("CreateUsage", createUsage_should_insert_usage_into_mongo)
	t.Run("GetAllUsages", getAllUsages_should_return_all_usages_in_mongo)
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

func getAllUsages_should_return_all_usages_in_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	usageService := service.NewUsageService(session.Copy(), dbName, collectionName)

	usages, err := usageService.GetAllUsages()
	if err != nil {
		t.Error("Coulddd not retrieve usages.")
	}
	if len(usages) != 0 {
		t.Errorf("Retrieved wrong number of usages. Expected %d, but got %d", 0, len(usages))
	}

	testId1, testUsageID1, testUsageName1 := "1", "this_usageID1", "this_usageName1"
	testId2, testUsageID2, testUsageName2 := "2", "this_usageID2", "this_usageName2"
	usage1 := &entity.Usage{
		ID:        testId1,
		UsageID:   testUsageID1,
		UsageName: testUsageName1,
	}
	usage2 := &entity.Usage{
		ID:        testId2,
		UsageID:   testUsageID2,
		UsageName: testUsageName2,
	}
	usageService.CreateUsage(usage1)
	usageService.CreateUsage(usage2)

	usages, err = usageService.GetAllUsages()
	if err != nil {
		t.Error("Could not retreive usages.")
	}
	if len(usages) != 2 {
		t.Errorf("Retrieved wrong number of usages. Expected %d, but got %d", 0, len(usages))
	}
	if usages[0].UsageID != testUsageID1 {
		t.Errorf("Wrong usage retrieved. Expected id %s, but got id %s", testUsageID1, usages[0].UsageID)
	}
	if usages[1].UsageID != testUsageID2 {
		t.Errorf("Wrong usage retrieved. Expected id %s, but got id %s", testUsageID2, usages[1].UsageID)
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
