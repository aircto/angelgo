package angelco

type UserResponse struct {
    User
}

type StartupRolesReponse struct {
    StartupRolesPagination
    StartupRoles []StartupRole `json:"startup_roles"`
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

type StartupRolesPagination struct {
    Pagination
}
