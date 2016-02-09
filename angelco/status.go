package angelco

import (
    "fmt"
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

func (api *AngelApi) StartupStatusList(startupId int64) (res *StatusUpdatesResponse, err error) {
    res = new(StatusUpdatesResponse)
    err = api.get("/status_updates/", url.Values{"startup_id": {strconv.FormatInt(startupId, 10)}}, res)
    return
}

func (api *AngelApi) PostMyStatus(message string) (res *StatusUpdateResponse, err error) {
    res = new(StatusUpdateResponse)
    err = api.post("/status_updates/", url.Values{}, url.Values{"message": {message}}, res)
    return
}

func (api *AngelApi) PostStartupStatus(startupId int64, message string) (res *StatusUpdateResponse, err error) {
    res = new(StatusUpdateResponse)
    err = api.post("/status_updates/", url.Values{"startup_id": {strconv.FormatInt(startupId, 10)}}, url.Values{"message": {message}}, res)
    return
}

func (api *AngelApi) RemoveStatus(statusId int64) (res *StatusUpdateResponse, err error) {
    res = new(StatusUpdateResponse)
    err = api.delete(fmt.Sprintf("/status_updates/%d/", statusId), url.Values{}, res)
    return
}
