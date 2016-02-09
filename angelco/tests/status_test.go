package angelco_testing

import (
    // "fmt"
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
