package angelco_testing

import (
    "os"
    "github.com/aircto/angelgo/angelco"
)

var CLIENT_ID = os.Getenv("CLIENT_ID")
var CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")

var api *angelco.AngelApi

func init() {
    api = angelco.New(CLIENT_ID, CLIENT_SECRET)
    api.SetAccessToken(ACCESS_TOKEN)
}
