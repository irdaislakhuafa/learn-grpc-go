package parameter

type UserPaginationParam struct {
	Limit     uint64 `json:"limit,omitempty"`
	Page      uint64 `json:"page,omitempty"`
	IsDeleted int64  `json:"is_deleted,omitempty"`
}

type UserGetParam struct {
	ID string `json:"id,omitempty"`
}