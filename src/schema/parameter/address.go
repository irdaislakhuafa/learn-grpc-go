package parameter

type AddressCreateParam struct {
	Country     string `json:"country,omitempty"`
	Province    string `json:"province,omitempty"`
	Regency     string `json:"regency,omitempty"`
	SubDistrict string `json:"sub_district,omitempty"`
}
