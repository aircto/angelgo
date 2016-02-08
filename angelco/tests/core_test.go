package angelco_testing

import (
    "fmt"
    "testing"
)

func TestGetAuthUrl(t *testing.T) {
    u := api.AuthUrl()
    eUrl := "https://angel.co/api/oauth/authorize"
    expected := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=email", eUrl, api.ClientId)
    if u != expected {
        t.Error("Fail")
    }
}

func TestGetAccessTokenError(t *testing.T) {
    _, err := api.GetAccessToken("d9dc7c0236660f3faecd7d6e4af8fbd2398ec5e70a8b4971")
    if err == nil {
        t.Error("Fail")
    }
}
