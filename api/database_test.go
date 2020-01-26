package api

import (
	"os"
	"testing"

	"go_chatible/env"
)

func TestMain(m *testing.M) {
	env.Load("../")
	code := m.Run()
	os.Exit(code)
}

func TestConnectDB(t *testing.T) {
	ConnectDB()
	res, err := DB.Exec("SELECT * FROM pg_catalog.pg_tables;")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.RowsReturned())
}
