package angelco_testing

import (
    // "fmt"
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


func TestGetUserErrorResponse(t *testing.T) {
    _, err := api.GetUser(int64(0))
    if err == nil {
        t.Error("Fail")
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

func TestGetUserStarupRoles(t *testing.T) {
    roles, err := api.GetUserStartupRoles(kaviraj_userId)
    if err != nil {
        t.Error(err)
    }
    id := roles.StartupRoles[0].Id
    if id == 0 {
        t.Error("Failed to get users Startup roles")
    }
}

func TestGetUserStarupRolesError(t *testing.T) {
    _, err := api.GetUserStartupRoles(int64(0))

    if err == nil {
        t.Error("Fail")
    }
}
