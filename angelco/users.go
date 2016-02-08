package angelgo

import (
    "fmt"
    "net/url"
)

func (api *AngelApi) Me() (res *UserResponse, err error){
    res = new(UserResponse)
    err = api.get(fmt.Sprintf("/me/"), url.Values{}, res)
    return
}

func (api *AngelApi) GetUser(userId int64) (res *UserResponse, err error) {
    res = new(UserResponse)
    err = api.get(fmt.Sprintf("/users/%d", userId), url.Values{}, res)
    return
}

func (api *AngelApi) GetUserStartupRoles(userId int64) (res *StartupRolesReponse, err error) {
    res = new(StartupRolesReponse)
    err = api.get(fmt.Sprintf("/users/%d/roles", userId), url.Values{}, res)
    return
}
