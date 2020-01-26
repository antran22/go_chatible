package api

import (
	"testing"

	"go_chatible/model/user"
)

func TestUserFetcherPool_FetchUser(t *testing.T) {
	fetcher := NewUserFetcherPool()
	defer fetcher.Close()
	usr := user.User{
		ID: "1785417844831933",
	}
	if err := fetcher.FetchUser(&usr); err != nil {
		t.Error(err)
	}
	usr = user.User{
		ID: "123",
	}
	if err := fetcher.FetchUser(&usr); err == nil {
		t.Fail()
	} else {
		t.Log("Expected error", err)
	}
}
