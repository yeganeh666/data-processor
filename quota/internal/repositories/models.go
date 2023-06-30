package repositories

import (
	"IofIPOS/shared/gormext"
	"github.com/google/uuid"
	"sync"
)

type User struct {
	sync.Mutex
	gormext.UniversalModel
}
type UserQuota struct {
	sync.Mutex
	gormext.UniversalModel
	UserID                  uuid.UUID `gorm:"type:uuid REFERENCES users(id);not null"`
	RequestsPerMinute       int
	CurrentDataVolume       int64
	TotalDataVolumePerMonth int64
}

type Object struct {
	gormext.UniversalModel
	Key       string    `gorm:"index"`
	UploadKey string    `gorm:"unique;not null"`
	UserID    uuid.UUID `gorm:"type:uuid REFERENCES users(id);not null"`
}
