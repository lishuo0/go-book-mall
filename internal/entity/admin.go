package entity

type AdminLoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminLoginResp struct {
	Token string `json:"token"`
}
