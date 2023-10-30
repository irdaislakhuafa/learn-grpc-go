package parameter

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

type UserDeleteParam struct {
	ID string `json:"id,omitempty"`
}
