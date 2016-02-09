package angelco

import (
    // "fmt"
    "strconv"
    "net/url"
)

func (api *AngelApi) MyStatusList() (res *StatusUpdatesResponse, err error) {
    res = new(StatusUpdatesResponse)
    err = api.get("/status_updates/", url.Values{}, res)
    return 
}

func (api *AngelApi) UserStatusList(userId int64) (res *StatusUpdatesResponse, err error) {
    res = new(StatusUpdatesResponse)
    err = api.get("/status_updates/", url.Values{"user_id": {strconv.FormatInt(userId, 10)}}, res)
    return
}
