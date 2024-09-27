package project

type User struct {
	Id       int    `json:"~"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Merchant struct {
	User
	MerchantID int `json:"merchantId"`
}
