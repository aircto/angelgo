package angelgo

type UserResponse struct {
    User
}

type MetaResponse struct {
    Meta
}

type Meta struct {
    ErrorType string `json:"error"`
    ErrorDescription string `json:"error_description"`
}
