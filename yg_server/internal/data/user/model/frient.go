package model

import "yug_server/internal/data"

// Friend 表的结构定义
type Friend struct {
	data.BaseModel
	UserID   uint64 `gorm:"type:bigint;not null" json:"user_id"`   // 当前用户 ID
	FriendID uint64 `gorm:"type:bigint;not null" json:"friend_id"` // 好友用户 ID
	Status   int    `gorm:"type:tinyint;default:0" json:"status"`  // 好友状态（0: 待接受, 1: 已接受, 2: 拒绝）
	User     User   `gorm:"foreignKey:UserID;references:ID"`       // 当前用户
	Friend   User   `gorm:"foreignKey:FriendID;references:ID"`     // 好友用户
}

// 表名
func (Friend) TableName() string {
	return "friends"
}
