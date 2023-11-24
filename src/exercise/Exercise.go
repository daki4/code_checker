package exercise

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	Id                       uuid.UUID     `json:"id" gorm:"primaryKey"`
	Author                   string        `json:"user" gorm:"foreignKey:username"`
	CreationTime             time.Time     `json:"creationTime"`
	LastEditTime             time.Time     `json:"lastEditTime"`
	Description              string        `json:"description"`
	AllowedLanguages         []Language    `json:"allowedLanguages" gorm:"many2many:exercise_languages;"`
	AllowedExecutionDuration int64         `json:"allowedExecutionDuration"`
	AllowedMemory            int64         `json:"allowedMemory"`
	Tests                    []ControlTest `json:"tests" gorm:"foreignKey:ExerciseId"`
}
