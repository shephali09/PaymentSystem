package entity

type User struct {
	UserId    string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}
