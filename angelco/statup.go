package angelco

import (
    "fmt"
    "net/url"
)

func (api *AngelApi) Startup(startupId int64) (res *StartupResponse, err error) {
    res = new(StartupResponse)
    err = api.get(fmt.Sprintf("/startup/%d/", startupId), url.Values{}, res)
    return
}
