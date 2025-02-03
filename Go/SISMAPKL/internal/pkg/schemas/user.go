package schemas

type LoginInputMahasiswa struct {
	NIM      string `json:"nim"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type ChangeProfileRequest struct {
}
