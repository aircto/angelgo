package angelco

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
    "bytes"
    // "io/ioutil"
)

var (
	scheme           = "https"
	version          = "1"
	apiEndpoint      = scheme + "://api.angel.co" + "/" + version
	oauthEndpoint    = scheme + "://angel.co/api"
	accessTokenUrl   = "/oauth/token"
	authorizationUrl = "/oauth/authorize"
)

type AngelApi struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
}

type AngelcoInterface interface {
    AuthUrl() string
    GetAccessToken(code string) (*AccessTokenResponse, error)
    SetAccessToken(access_token string)
    Me()(*UserResponse, error)
    User(userId int64) (*UserResponse, error)
    UserStartupRoles(userId int64) (*StartupRolesReponse, error)

    // Status message apis
    MyStatusList() (*StatusUpdatesResponse, error)
    UserStatusList(userId int64) (*StatusUpdatesResponse, error)

    // Will give a list of status messages if you are team member of the startup else it will be emptyn
    StartupStatusList(startupId int64) (*StatusUpdatesResponse, error)
    PostMyStatus(message string) (*StatusUpdateResponse, error)
    PostStartupStatus(startupId int64, message string) error
    RemoveStatus(statusId int64) error

    // Jobs api
    JobsList() (*JobsListResponse, error)
    Job(jobId int64) (*JobResponse, error)
    JobsOfStartup(startupId int64) ([]Job , error)
    JobsOfTag(tagId int64) (*JobsListResponse, error)

    //Startups api
    Startup(startupId int64)
    
}

func New(clientId string, clientSecret string) *AngelApi {
	if clientId == "" || clientSecret == "" {
		panic("ClientId and ClientSecret should be given to create an AngelCo Api")
	}

	return &AngelApi{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}

func (api *AngelApi) AuthUrl() string {
    u, err := url.Parse(oauthEndpoint + authorizationUrl)
    if err != nil {
        panic(err)
    }
    v := url.Values{}

    if api.ClientId == "" {
        panic("Cannot use AuthUrl without setting ClientId")
    }
    
    v.Set("client_id", api.ClientId)
    v.Set("scope", "email")
    v.Set("response_type", "code")
    u.RawQuery = v.Encode()
    return u.String()
}

func (api *AngelApi) tokenUrl() string {
    return oauthEndpoint + accessTokenUrl
}

func (api *AngelApi) GetAccessToken(code string) (AccessTokenResponse,  error) {
    token := new(AccessTokenResponse)

    u, err := url.Parse(api.tokenUrl())
    if err != nil {
        return *token, err
    }

    res, err := http.PostForm(u.String(), url.Values{"client_id": {api.ClientId}, "client_secret": {api.ClientSecret}, "code": {code}, "grant_type": {"authorization_code"}})
    if err != nil {
        return *token, err
    }

    if res.StatusCode != 200 {
        token_err := new(AccessError)
        err = decodeResponse(res.Body, token_err)
        if err != nil {
            panic(err)
        }
        return *token, token_err
    }
    err = decodeResponse(res.Body, token)
    if err != nil {
        panic(err)
    }
    return *token, nil
}

func (api *AngelApi) SetAccessToken(accessToken string) {
	if accessToken == "" {
		panic("Cannot set empty Access Token")
	}
	api.AccessToken = accessToken
	return
}

// Create actuall http Request based on type GET, POST, DELETE

func (api *AngelApi) get(path string, params url.Values, r interface{}) error {
	params = api.extendParams(params)
	req, err := buildGetRequest(urlify(apiEndpoint, path), params)

	if err != nil {
		return err
	}
	return api.do(req, r)
}

func (api *AngelApi) post(path string, params url.Values, data url.Values, r interface{}) error {
    params = api.extendParams(params)
    req, err := buildPostRequest(urlify(apiEndpoint, path), params, data)

    if err != nil {
        return err
    }
    return api.do(req, r)
}

func (api *AngelApi) delete(path string, params url.Values, r interface{}) error {
    params = api.extendParams(params)
    req, err := buildDeleteRequest(urlify(apiEndpoint, path), params)

    if err != nil {
        return err
    }
    return api.do(req, r)
}

func urlify(base_url, path string) string {
	return base_url + path
}


func (m *ErrorResponse) Error() string {
	return fmt.Sprintf("Error making api call:  %s %s", m.ErrorJson.Type, m.ErrorJson.Message)
}

func (m *AccessError) Error() string {
	return fmt.Sprintf("Error making api call:  %s %s", m.ErrorType, m.ErrorDescription)
}

func (api *AngelApi) extendParams(p url.Values) url.Values {
	if p == nil {
		return url.Values{}
	}
	if api.AccessToken == "" {
		panic("Please set Acess Token by calling api.SetAccessToken before making request")
	}
	p.Set("access_token", api.AccessToken)
	return p
}

func buildGetRequest(urlStr string, params url.Values) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if params != nil {
		if u.RawQuery != "" {
			return nil, fmt.Errorf("Please remove any query params from urlString")
		}
		u.RawQuery = params.Encode()
	}
    fmt.Println(u.String())
	return http.NewRequest("GET", u.String(), nil)
}

func buildPostRequest(urlStr string, params url.Values, data url.Values) (r *http.Request, err error) {
    u, err := url.Parse(urlStr)
    if err != nil {
        return nil, err
    }
    
    if params != nil {
		if u.RawQuery != "" {
			return nil, fmt.Errorf("Please remove any query params from urlString")
		}
		u.RawQuery = params.Encode()
	}
    
    return http.NewRequest("POST", u.String(), bytes.NewBufferString(data.Encode()))
}

func buildDeleteRequest(urlStr string, params url.Values) (r *http.Request, err error) {
    u, err := url.Parse(urlStr)
    if err != nil {
        return nil, err
    }
    if params != nil {
		if u.RawQuery != "" {
			return nil, fmt.Errorf("Please remove any query params from urlString")
		}
		u.RawQuery = params.Encode()
	}
    return http.NewRequest("DELETE", u.String(), nil)
}

func (api *AngelApi) do(req *http.Request, r interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return apiError(resp)
	}
	return decodeResponse(resp.Body, r)
}

func apiError(resp *http.Response) error {
	m := new(ErrorResponse)
	if err := decodeResponse(resp.Body, m); err != nil {
		return err
	}
	err := ErrorResponse(*m)
	return &err
}

func decodeResponse(body io.Reader, to interface{}) error {
	err := json.NewDecoder(body).Decode(to)

	if err != nil {
		return fmt.Errorf("angelco: error decoding body: %s", err.Error())
	}
	return nil
}
