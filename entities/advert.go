package project

type AdvertItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagesPath  string `json:"imagesPath"`
	MerchantID  int    `json:"merchantId"`
}

type MerchantAdvert struct {
	Id         int `json:"id"`
	MerchantID int `json:"merchantId"`
	AdvertID   int `json:"advertId"`
}
