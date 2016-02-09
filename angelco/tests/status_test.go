package angelco_testing

import (
//    "fmt"
    "testing"
)


func TestMyStatusMessages(t *testing.T) {
    _, err := api.MyStatusList()
    if err != nil {
        t.Error(err)
    }
}

func TestUserStatusMessages(t *testing.T) {
    _, err := api.UserStatusList(kaviraj_userId)
    if err != nil {
        t.Error(err)
    }
}

func TestUserStatusMessagesError(t *testing.T) {
    _, err := api.UserStatusList(int64(0))
    if err == nil {
        t.Error("Fail")
    }
}

func TestStartupStatusMessages(t *testing.T) {
    _, err := api.StartupStatusList(launchyardId)
    if err != nil {
        t.Error(err)
    }
}

func TestStartupStatusMessagesError(t *testing.T) {
    _, err := api.StartupStatusList(int64(0))
    if err == nil {
        t.Error("Fail")
    }
}

func TestPostMyStatus(t *testing.T) {
    expected := "Guess what?.. Im posting this via Golang"
    _, err := api.PostMyStatus(expected)

    if err != nil {
        t.Error(err)
    }
}

func TestPostStartupStatus(t *testing.T) {
    expected := "Guess what?.. Im posting this via Golang"
    res, err := api.PostStartupStatus(launchyardId, expected)

    if err != nil {
        panic(err)
    }

    if err == nil {
        if res.Message != expected {
            t.Error("Status updated wrongly")
        }
    }
}

func TestRemoveStatus(t *testing.T) {
    res, err := api.PostMyStatus("testing")
    if err != nil {
        t.Error(err)
    }
    id := res.Id
    res, err = api.RemoveStatus(id)
    if err != nil {
        t.Error(err)
    }
}
