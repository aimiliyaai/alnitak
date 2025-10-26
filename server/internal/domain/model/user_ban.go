package model

import (
	"time"

	"gorm.io/gorm"
)

type UserBan struct {
	gorm.Model
	Uid      uint      `gorm:"comment:用户ID;not null;index"`
	EndTime  time.Time `gorm:"comment:封禁截止时间"`
	Reason   string    `gorm:"type:varchar(255);comment:封禁理由"`
	Status   int       `gorm:"default:0;comment:状态:0封禁中、1管理员解封、2自动解封、3永久封禁、4封禁撤销"`
	Operator uint      `gorm:"default:0;comment:操作人ID:0系统"`
}

func (UserBan) TableName() string {
	return "user_ban"
}
