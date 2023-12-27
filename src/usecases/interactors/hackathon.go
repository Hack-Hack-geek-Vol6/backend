package interactors

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/google/uuid"
)

const (
	// HACKATHON_IMAGE_DIR はハッカソンの画像を保存するディレクトリ
	HACKATHON_IMAGE_DIR = "hackathon/"
)

type HackathonInteractor struct {
	Hackathon       dai.HackathonDai
	HackathonStatus dai.HackathonStatusDai
	FileStore       dai.FileStore
	HackathonOutput ports.HackathonOutputBoundary
}

func NewHackathonInteractor(hackathonDai dai.HackathonDai, HackathonStatus dai.HackathonStatusDai, filestoreDai dai.FileStore, hackathonOutput ports.HackathonOutputBoundary) ports.HackathonInputBoundary {
	return &HackathonInteractor{
		Hackathon:       hackathonDai,
		HackathonStatus: HackathonStatus,
		FileStore:       filestoreDai,
		HackathonOutput: hackathonOutput,
	}
}

func (hi *HackathonInteractor) CreateHackathon(ctx context.Context, in *ports.InputCreatehackathonData) (int, *response.CreateHackathon) {
	// 画像があるときは画像を保存してLinkを追加
	// 画像がないときは初期画像をLinkに追加
	var (
		imageLinks  string = config.Config.Server.DefaultHackathonImage
		hackathonID string = uuid.New().String()
	)

	if in.ImageFile != nil {
		src, err := in.ImageFile.Open()
		if err != nil {
			return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
				Error:    err,
				Response: nil,
			})
		}
		defer src.Close()

		data, err := io.ReadAll(src)
		if err != nil {
			return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
				Error:    err,
				Response: nil,
			})
		}

		// 画像を保存してLinkを追加
		log.Println("uploading image")
		links, err := hi.FileStore.UploadFile(ctx, data, fmt.Sprintf("%s%s.%s", HACKATHON_IMAGE_DIR, hackathonID, in.ImageFile.Filename))
		if err != nil {
			return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
				Error:    err,
				Response: nil,
			})
		}
		imageLinks = links
	}

	// ハッカソンを作成
	if err := hi.Hackathon.Create(ctx, &models.Hackathon{
		HackathonID: hackathonID,
		Name:        in.Name,
		Link:        in.Link,
		Expired:     in.Expired,
		StartDate:   in.StartDate,
		Term:        in.Term,
		Icon:        imageLinks,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}); err != nil {
		return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	// ステータスを作成
	statusTags, err := hi.craeteStatues(ctx, hackathonID, in.Statuses)
	if err != nil {
		return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	// TODO:ハッカソンを取得？

	return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
		Error: nil,
		Response: &response.CreateHackathon{
			HackathonID: hackathonID,
			StatusTags:  statusTags,
		},
	})
}

func (hi *HackathonInteractor) craeteStatues(ctx context.Context, hackathonID string, statuses []int64) ([]*response.StatusTag, error) {
	if err := hi.HackathonStatus.Create(ctx, hackathonID, statuses); err != nil {
		return nil, err
	}

	result, err := hi.HackathonStatus.FindAll(ctx, []string{hackathonID})
	if err != nil {
		return nil, err
	}

	var statusTags []*response.StatusTag
	for _, status := range result {
		statusTags = append(statusTags, &response.StatusTag{
			ID:     status.StatusID,
			Status: status.Status,
		})
	}
	return statusTags, nil
}