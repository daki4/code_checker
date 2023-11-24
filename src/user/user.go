package user

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user User) CheckPassword(password string) bool {
	return true
}
