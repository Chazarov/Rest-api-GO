package project

type AdvertItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ImagesPath  string `json:"imagesPath"`
	UserId      int    `json:"userId" binding:"required"`
}
