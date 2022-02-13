package models

// DB structures
type User struct {
	Id        int32  `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// Json structures
type JsonUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonResetUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonUserToken struct {
	Token string `json:"token"`
}

type JsonUserGet struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonUserDelete struct {
	Id int32 `db:"id"`
}

type JsonUserUpdate struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonUser struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
