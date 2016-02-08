package angelco

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
    // "io/ioutil"
)

var (
	scheme           = "https"
	version          = "/1/"
	apiEndpoint      = scheme + "://api.angel.co" + version
	oauthEndpoint    = scheme + "://angel.co/api" + version
	accessTokenUrl   = "/oauth/token"
	authorizationUrl = "/oauth/authorize"
)

type AngelApi struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
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

func urlify(base_url, path string) string {
	return base_url + path
}


func (m *ErrorResponse) Error() string {
	return fmt.Sprintf("Error making api call:  %s %s", m.ErrorJson.Type, m.ErrorJson.Message)
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

	return http.NewRequest("GET", u.String(), nil)
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
