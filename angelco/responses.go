package angelco

type UserResponse struct {
    *User
}

type StartupRolesReponse struct {
    *Pagination
    StartupRoles []StartupRole `json:"startup_roles"`
}

type StatusUpdateResponse struct {
    *StatusUpdate
}

type StatusUpdatesResponse struct {
    *Pagination
    StatusUpdates []StatusUpdate `json:"status_updates"`
}

type JobResponse struct {
    *Job
}

type JobsListResponse struct {
    *Pagination
    Jobs []Job
}

type StartupResponse struct {
    *Startup
}

type AccessTokenResponse struct {
    *AccessToken
}

type ErrorResponse struct {
    ErrorJson `json:"error"`
    Success bool
}

type AccessError struct {
    ErrorType string `json:"error"`
    ErrorDescription string `json:"error_description"`
}

type ErrorJson struct {
    Type string
    Message string
}

type Pagination struct {
	Total    int64
	Page     int64
	PerPage  int64 `json:"per_page"`
	LastPage int64 `json:"last_page"`
}

