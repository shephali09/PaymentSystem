package entity

type Notification struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Date      string `json:"date"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
