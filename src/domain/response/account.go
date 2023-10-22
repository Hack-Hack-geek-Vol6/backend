package response

import "github.com/hackhack-Geek-vol6/backend/src/repository"

type Account struct {
	AccountID       string `json:"account_id"`
	Username        string `json:"username"`
	Icon            string `json:"icon"`
	ExplanatoryText string `json:"explanatory_text"`
	Rate            int32  `json:"rate"`
	Email           string `json:"email"`
	Locate          string `json:"locate"`
	GithubLink      string `json:"github_link"`
	TwitterLink     string `json:"twitter_link"`
	DiscordLink     string `json:"discord_link"`

	ShowLocate bool `json:"show_locate"`
	ShowRate   bool `json:"show_rate"`

	TechTags   []repository.TechTag   `json:"tech_tags"`
	Frameworks []repository.Framework `json:"frameworks"`
}

type AccountRate struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Icon      string `json:"icon"`
	Rate      int32  `json:"rate"`
}

type GetJoinRoom struct {
	RoomID string `json:"room_id"`
	Title  string `json:"title"`
}