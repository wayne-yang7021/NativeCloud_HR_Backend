package model

import "time"

type CheckinRecord struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	CheckinAt time.Time `json:"checkin_at"`
	SiteClass string    `json:"site_class"` // A, B, C...
	Site      string    `json:"site"`       // 儲位名稱
	Note      string    `json:"note"`       // optional: 上班/下班
}
