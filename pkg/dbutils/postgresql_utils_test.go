package dbutils

import (
	"testing"

	"github.com/cuijxin/my-kubeapps/pkg/common/datastore"
)

func Test_NewPGManager(t *testing.T) {
	config := datastore.Config{URL: "10.11.12.13:5432", Database: "assets", Username: "postgres", Password: "123"}
	m, err := NewPGManager(config, "kubeapps")
	if err != nil {
		t.Errorf("Found error %v", err)
	}
	expectedConnStr := "10.11.12.13 port=5432 user=postgres password=123 dbname=assets sslmode=disable"
	if m.connStr != expectedConnStr {
		t.Errorf("Expected %s got %s", expectedConnStr, m.connStr)
	}
}
