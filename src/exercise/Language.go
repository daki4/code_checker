package exercise

type Language struct {
	Name    string `json:"name" gorm:"primaryKey"`
	Version string `json:"version"`
}
