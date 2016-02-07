package angelgo

import (
    "fmt"
    "net/url"
)

func (api *AngelApi) Me() (res *UserResponse, err error){
    res = new(UserResponse)
    err = api.get(fmt.Sprintf("/1/me/"), url.Values{}, res)
    return
}
