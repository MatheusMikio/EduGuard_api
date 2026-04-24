package schemas

import "gorm.io/gorm"

type CameraStatus string

const (
	CameraStatusOnline      CameraStatus = "ONLINE"
	CameraStatusOffLine     CameraStatus = "OFFLINE"
	CameraStatusManutenence CameraStatus = "MANUTENÇÃO"
)

type Camera struct {
	gorm.Model
	Label     string       `gorm:"not null"`
	Location  string       `gorm:"not null;size=100"`
	StreamURL string       `gorm:"not null"`
	Status    CameraStatus `gorm:"not null;default:'ONLINE'"`
	SchoolID  uint         `gorm:"not null"`
	School    *School      `gorm:"foreignKey:SchoolID"`
	RoomID    uint         `gorm:"not null"`
	Room      *Room        `gorm:"foreignKey:RoomID"`
}
