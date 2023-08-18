// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package repository

import (
	"database/sql"
	"time"
)

type Account struct {
	AccountID       string         `json:"account_id"`
	Username        string         `json:"username"`
	Icon            sql.NullString `json:"icon"`
	ExplanatoryText sql.NullString `json:"explanatory_text"`
	LocateID        int32          `json:"locate_id"`
	Rate            int32          `json:"rate"`
	Character       sql.NullInt32  `json:"character"`
	ShowLocate      bool           `json:"show_locate"`
	ShowRate        bool           `json:"show_rate"`
	CreateAt        time.Time      `json:"create_at"`
	UpdateAt        time.Time      `json:"update_at"`
	IsDelete        bool           `json:"is_delete"`
	Email           string         `json:"email"`
}

type AccountFramework struct {
	AccountID   string `json:"account_id"`
	FrameworkID int32  `json:"framework_id"`
}

type AccountPastWork struct {
	Opus      int32  `json:"opus"`
	AccountID string `json:"account_id"`
}

type AccountTag struct {
	AccountID string `json:"account_id"`
	TechTagID int32  `json:"tech_tag_id"`
}

type AccountsAchievment struct {
	AccountID    string    `json:"account_id"`
	AchievmentID int32     `json:"achievment_id"`
	CreateAt     time.Time `json:"create_at"`
}

type Achievment struct {
	AchievmentID int32     `json:"achievment_id"`
	Achievment   string    `json:"achievment"`
	Description  string    `json:"description"`
	Icon         string    `json:"icon"`
	Conditions   string    `json:"conditions"`
	CreateAt     time.Time `json:"create_at"`
	IsDelete     bool      `json:"is_delete"`
}

type Award struct {
	AwardID int32  `json:"award_id"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
}

type AwardDatum struct {
	AwardDataID int32 `json:"award_data_id"`
	AwardID     int32 `json:"award_id"`
	HackathonID int32 `json:"hackathon_id"`
}

type Follow struct {
	ToAccountID   string    `json:"to_account_id"`
	FromAccountID string    `json:"from_account_id"`
	CreateAt      time.Time `json:"create_at"`
}

type Framework struct {
	FrameworkID int32  `json:"framework_id"`
	TechTagID   int32  `json:"tech_tag_id"`
	Framework   string `json:"framework"`
}

type Hackathon struct {
	HackathonID int32          `json:"hackathon_id"`
	Name        string         `json:"name"`
	Icon        sql.NullString `json:"icon"`
	Description string         `json:"description"`
	Link        string         `json:"link"`
	Expired     time.Time      `json:"expired"`
	StartDate   time.Time      `json:"start_date"`
	Term        int32          `json:"term"`
}

type HackathonStatusTag struct {
	HackathonID int32 `json:"hackathon_id"`
	StatusID    int32 `json:"status_id"`
}

type Like struct {
	Opus      int32     `json:"opus"`
	AccountID string    `json:"account_id"`
	CreateAt  time.Time `json:"create_at"`
	IsDelete  bool      `json:"is_delete"`
}

type Locate struct {
	LocateID int32  `json:"locate_id"`
	Name     string `json:"name"`
}

type PastWork struct {
	Opus            int32         `json:"opus"`
	Name            string        `json:"name"`
	ThumbnailImage  string        `json:"thumbnail_image"`
	ExplanatoryText string        `json:"explanatory_text"`
	AwardDataID     sql.NullInt32 `json:"award_data_id"`
	CreateAt        time.Time     `json:"create_at"`
	UpdateAt        time.Time     `json:"update_at"`
	IsDelete        bool          `json:"is_delete"`
}

type PastWorkFramework struct {
	Opus        int32 `json:"opus"`
	FrameworkID int32 `json:"framework_id"`
}

type PastWorkTag struct {
	Opus      int32 `json:"opus"`
	TechTagID int32 `json:"tech_tag_id"`
}

type RateEntity struct {
	AccountID string    `json:"account_id"`
	Rate      int32     `json:"rate"`
	CreateAt  time.Time `json:"create_at"`
}

type Role struct {
	RoleID int32  `json:"role_id"`
	Role   string `json:"role"`
}

type Room struct {
	RoomID      string    `json:"room_id"`
	HackathonID int32     `json:"hackathon_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	MemberLimit int32     `json:"member_limit"`
	IncludeRate bool      `json:"include_rate"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	IsDelete    bool      `json:"is_delete"`
}

type RoomsAccount struct {
	AccountID string        `json:"account_id"`
	RoomID    string        `json:"room_id"`
	Role      sql.NullInt32 `json:"role"`
	IsOwner   bool          `json:"is_owner"`
	CreateAt  time.Time     `json:"create_at"`
}

type StatusTag struct {
	StatusID int32  `json:"status_id"`
	Status   string `json:"status"`
}

type TechTag struct {
	TechTagID int32  `json:"tech_tag_id"`
	Language  string `json:"language"`
	Icon      string `json:"icon"`
}

type Tutor struct {
	TutorID     string         `json:"tutor_id"`
	Title       string         `json:"Title"`
	Description sql.NullString `json:"description"`
	CreateAt    time.Time      `json:"create_at"`
	UpdateAt    time.Time      `json:"update_at"`
	IsDelete    bool           `json:"is_delete"`
}
