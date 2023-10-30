package parameter

type UserPaginationParam struct {
	Limit     uint64 `json:"limit,omitempty"`
	Page      uint64 `json:"page,omitempty"`
	IsDeleted int64  `json:"is_deleted,omitempty"`
}

type UserGetParam struct {
	ID string `json:"id,omitempty"`
}

type UserCreateParam struct {
	Name    string             `json:"name,omitempty"`
	Email   string             `json:"email,omitempty"`
	Age     int                `json:"age,omitempty"`
	Hobbies []string           `json:"hobbies,omitempty"`
	Address AddressCreateParam `json:"address,omitempty"`
}

type UserUpdateParam struct {
	ID      string             `json:"id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Email   string             `json:"email,omitempty"`
	Age     int                `json:"age,omitempty"`
	Hobbies []string           `json:"hobbies,omitempty"`
	Address AddressUpdateParam `json:"address,omitempty"`
}
