package model

import "time"

type Model struct {
	ID        int       `json:"id"`
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}
