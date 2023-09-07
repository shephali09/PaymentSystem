package entity

type Notification struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Date      string `json:"date"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
