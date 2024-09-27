package project

type Advert struct {
	Id          int    `json:"~"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AdvertsList struct {
	Id         int
	MerchantID int
	AdvertID   int
}
