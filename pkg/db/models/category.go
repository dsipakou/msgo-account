package models

// DB structures
type Category struct {
	Id        int32  `db:"id"`
	Name      string `db:"name"`
	Parent    string `db:"parent"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// Json structures
type JsonCategoryCreate struct {
	Name   string `json:"name"`
	Parent string `json:"parentName"`
}

type JsonCategoryUpdate struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parentName"`
}

type JsonCategoryDelete struct {
	Id int32 `db:"id"`
}

type JsonCategoryResponse struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	Parent    string `json:"parentName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
