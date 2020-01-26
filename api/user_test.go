package api

import (
	"net/http"
	"testing"

	"go_chatible/model/user"
)

func TestGetUserData(t *testing.T) {
	fetcher := NewUserFetcherPool()
	defer fetcher.Close()
	t.Log(fetcher.urlFormatString)
	client := http.Client{}
	usr, err := getUserData(&client, "2191897550831658", fetcher.urlFormatString)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(usr)
	}

}

func TestUserFetcherPool_FetchUserData(t *testing.T) {
	fetcher := NewUserFetcherPool()
	defer fetcher.Close()
	usr := user.User{
		ID: "2191897550831658",
	}
	if err := fetcher.FetchUserData(&usr); err != nil {
		t.Error(err)
	}
	usr = user.User{
		ID: "123",
	}
	if err := fetcher.FetchUserData(&usr); err == nil {
		t.Fail()
	} else {
		t.Log("Expected error", err)
	}
}
