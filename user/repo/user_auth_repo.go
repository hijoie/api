package repo

type UserAuth struct {
	Phone    string `json:"phone,omitempty"`
	VrifCode string `json:"vrif_code,omitempty"`
}
