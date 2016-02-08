package angelco_testing

import (
    "testing"
)

func TestMeEndpoint(t *testing.T) {
    me, err := api.Me()
    if err != nil {
        t.Error(err)
    }
    if me.Id == 0 {
        t.Error("Error in getting my details")
    }

    // Check email scope
    if me.Email == "" {
        t.Error("Wrong scope.")
    }
}


func TestGetUserEndpoint(t *testing.T) {
    user, err := api.GetUser(kaviraj_userId)

    if err != nil {
        t.Error(err)
    }

    if user.Id == 0 {
        t.Error("Error in fetching user detail")
    }
}
