package mgtest

import (
	"testing"

	"github.com/cuijxin/my-kubeapps/pkg/common/datastore"
	"github.com/cuijxin/my-kubeapps/pkg/dbutils"
	"github.com/cuijxin/my-kubeapps/pkg/dbutils/dbutilstest"
)

const (
	// EnvvarMongoTests enables test that run against a local mongo db
	EnvvarMongoTests = "ENABLE_MONGO_INTEGRATION_TESTS"
)

func SkipIfNoDB(t *testing.T) {
	if !dbutilstest.IsEnvVarTrue(t, EnvvarMongoTests) {
		t.Skipf("skipping mongodb tests as %q not to be true", EnvvarMongoTests)
	}
}

func OpenTestManager(t *testing.T) *dbutils.MongodbAssetManager {
	manager := dbutils.NewMongoDBManager(datastore.Config{
		URL:      "localhost:27017",
		Username: "root",
		Password: "testpassword",
	}, dbutilstest.KubeappsTestNamespace)

	err := manager.Init()
	if err != nil {
		t.Fatalf("%s+v", err)
	}

	return manager
}

// GetInitializedManager returns an initialized mongodb manager ready for testing.
func GetInitializedManager(t *testing.T) (*dbutils.MongodbAssetManager, func()) {
	manager := OpenTestManager(t)
	cleanup := func() { manager.Close() }

	err := manager.InvalidateCache()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	return manager, cleanup
}
