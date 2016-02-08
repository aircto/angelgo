package angelgo

type UserResponse struct {
    User
}

type StartupRolesReponse struct {
    Pagination *StartupRolesPagination
    StartupRoles []StartupRole
}

type MetaResponse struct {
    Meta
}

type Meta struct {
    ErrorType string `json:"error"`
    ErrorDescription string `json:"error_description"`
}

type Pagination struct {
	Total    int64
	Page     int64
	PerPage  int64 `json:"per_page"`
	LastPage int64 `json:"last_page"`
}

type StartupRolesPagination struct {
    *Pagination
}
