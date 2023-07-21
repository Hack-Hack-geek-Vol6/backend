package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateRoomTxParams struct {
	// ルーム登録部分
	Rooms
	// RoomsAccounts登録部分
	UserID string
}
type RoomTechTags struct {
	TechTag TechTags `json:"tech_tag"`
	Count   int32    `json:"count"`
}
type RoomFramework struct {
	Framework Frameworks `json:"framework"`
	Count     int32      `json:"count"`
}
type RoomHackathonData struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
type NowRoomAccounts struct {
	UserID  string `json:"user_id"`
	Icon    string `json:"icon"`
	IsOwner bool   `json:"is_owner"`
}

type CreateRoomTxResult struct {
	Rooms
	Hackathon  RoomHackathonData `json:"hackathon"`
	Accounts   []NowRoomAccounts `json:"accounts"`
	TechTags   []RoomTechTags    `json:"techtags"`
	Frameworks []RoomFramework   `json:"frameworks"`
}

// TechTagの配列にマージする
func MargeTechTagArray(roomTechTags []RoomTechTags, techtag TechTags) []RoomTechTags {
	for _, roomTechTag := range roomTechTags {
		if roomTechTag.TechTag == techtag {
			roomTechTag.Count++
		}
	}
	roomTechTags = append(roomTechTags, RoomTechTags{
		TechTag: techtag,
		Count:   1,
	})

	return roomTechTags
}

// フレームワークの配列にマージする
func MargeFrameworkArray(roomFramework []RoomFramework, framework Frameworks) []RoomFramework {
	for _, roomFramework := range roomFramework {
		if roomFramework.Framework == framework {
			roomFramework.Count++
		}
	}
	roomFramework = append(roomFramework, RoomFramework{
		Framework: framework,
		Count:     1,
	})

	return roomFramework
}

func (store *SQLStore) CreateRoomTx(ctx context.Context, arg CreateRoomTxParams) (CreateRoomTxResult, error) {
	var result CreateRoomTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// ハッカソンデータを送る
		hackathon, err := q.GetHackathonByID(ctx, arg.HackathonID)
		if err != nil {
			return err
		}
		result.Hackathon = RoomHackathonData{
			ID:   hackathon.HackathonID,
			Name: hackathon.Name,
			Icon: hackathon.Icon.String,
		}

		// ルームを登録する
		result.Rooms, err = q.CreateRoom(ctx, CreateRoomParams{
			RoomID:      arg.RoomID,
			HackathonID: hackathon.HackathonID,
			Title:       arg.Title,
			Description: arg.Description,
			MemberLimit: arg.MemberLimit,
			IsDelete:    false,
		})
		if err != nil {
			return err
		}

		// ルームのオーナーを登録する
		_, err = q.CreateRoomsAccounts(ctx, CreateRoomsAccountsParams{
			UserID:  arg.UserID,
			RoomID:  result.RoomID,
			IsOwner: true,
		})

		if err != nil {
			return err
		}
		accounts, err := q.GetRoomsAccountsByRoomID(ctx, result.RoomID)
		if err != nil {
			return err
		}

		for _, account := range accounts {
			result.Accounts = append(result.Accounts, NowRoomAccounts{
				UserID:  account.UserID.String,
				Icon:    account.Icon.String,
				IsOwner: account.IsOwner,
			})
		}

		// ルーム内のユーザをもとにユーザの持つ技術タグとフレームワークタグを配列に落とし込む（力業
		for _, account := range result.Accounts {
			techTags, err := q.ListAccountTagsByUserID(ctx, account.UserID)
			if err != nil {
				return err
			}
			for _, techTag := range techTags {
				result.TechTags = MargeTechTagArray(result.TechTags, TechTags{
					TechTagID: techTag.TechTagID.Int32,
					Language:  techTag.Language.String,
				})
			}

			frameworks, err := q.ListAccountFrameworksByUserID(ctx, account.UserID)
			if err != nil {
				return err
			}
			for _, framework := range frameworks {
				result.Frameworks = MargeFrameworkArray(result.Frameworks, Frameworks{
					FrameworkID: framework.FrameworkID.Int32,
					TechTagID:   framework.TechTagID.Int32,
					Framework:   framework.Framework.String,
				})
			}
		}
		return nil
	})
	return result, err
}

type ListRoomTxParam struct {
}

type ListRoomTxRoomInfo struct {
	RoomID      uuid.UUID `json:"room_id"`
	Title       string    `json:"title"`
	MemberLimit int32     `json:"member_limit"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListRoomTxHackathonInfo struct {
	HackathonID   int32  `json:"hackathon_id"`
	HackathonName string `json:"hackathon_name"`
	Icon          string `json:"icon"`
}
type ListRoomTxResult struct {
	Rooms             ListRoomTxRoomInfo      `json:"rooms"`
	Hackathon         ListRoomTxHackathonInfo `json:"hackathon"`
	NowMember         []NowRoomAccounts       `json:"now_member"`
	MembersTechTags   []RoomTechTags          `json:"members_tech_tags"`
	MembersFrameworks []RoomFramework         `json:"members_frameworks"`
}

func (store *SQLStore) ListRoomTx(ctx context.Context, arg ListRoomTxParam) ([]ListRoomTxResult, error) {
	var result []ListRoomTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// ルーム一覧を取得してくる
		rooms, err := q.ListRoom(ctx, 100)
		if err != nil {
			return err
		}
		// それぞれのルームの確認
		for _, room := range rooms {
			var oneRoomInfos ListRoomTxResult
			oneRoomInfos.Rooms = ListRoomTxRoomInfo{
				RoomID:      room.RoomID,
				Title:       room.Title,
				MemberLimit: room.MemberLimit,
				CreatedAt:   room.CreateAt,
			}
			hackathon, err := q.GetHackathonByID(ctx, room.HackathonID)
			if err != nil {
				return err
			}
			// ハッカソンの追加
			oneRoomInfos.Hackathon = ListRoomTxHackathonInfo{
				HackathonID:   hackathon.HackathonID,
				HackathonName: hackathon.Name,
				Icon:          hackathon.Icon.String,
			}

			members, err := q.GetRoomsAccountsByRoomID(ctx, room.RoomID)
			if err != nil {
				return err
			}
			// アカウントの追加
			for _, account := range members {
				oneRoomInfos.NowMember = append(oneRoomInfos.NowMember, NowRoomAccounts{
					UserID:  account.UserID.String,
					Icon:    account.Icon.String,
					IsOwner: account.IsOwner,
				})
				// タグの追加
				techTags, err := q.ListAccountTagsByUserID(ctx, account.UserID.String)
				if err != nil {
					return err
				}
				for _, techTag := range techTags {
					oneRoomInfos.MembersTechTags = MargeTechTagArray(oneRoomInfos.MembersTechTags, TechTags{
						TechTagID: techTag.TechTagID.Int32,
						Language:  techTag.Language.String,
					})
				}
				// FWの追加
				frameworks, err := q.ListAccountFrameworksByUserID(ctx, account.UserID.String)
				if err != nil {
					return err
				}
				for _, framework := range frameworks {
					oneRoomInfos.MembersFrameworks = MargeFrameworkArray(oneRoomInfos.MembersFrameworks, Frameworks{
						FrameworkID: framework.FrameworkID.Int32,
						TechTagID:   framework.TechTagID.Int32,
						Framework:   framework.Framework.String,
					})
				}
			}
			result = append(result, oneRoomInfos)
		}
		return err
	})
	return result, err
}
