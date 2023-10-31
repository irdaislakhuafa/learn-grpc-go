package parameter

type RoleGetParam struct {
	ID string `json:"id,omitempty"`
}

type RoleCreateParam struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type RoleUpdateParam struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type RoleDeleteParam struct {
	ID string `json:"id,omitempty"`
}
