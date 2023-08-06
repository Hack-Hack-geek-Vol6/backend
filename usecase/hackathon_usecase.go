package usecase

import (
	"context"
	"database/sql"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type hackathonUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewHackathonUsercase(store db.Store, timeout time.Duration) domain.HackathonUsecase {
	return &hackathonUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (hu *hackathonUsecase) CreateHackathon(ctx context.Context, body domain.CreateHackathonParams) (result domain.HackathonResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.contextTimeout)
	defer cancel()

	Icon, err := hu.store.UploadImage(ctx, body.Image)
	if err != nil {
		return
	}

	hackathon, err := hu.store.CreateHackathon(ctx, db.CreateHackathonParams{
		Name: body.Name,
		Icon: sql.NullString{
			String: Icon,
			Valid:  true,
		},
		Description: body.Description,
		Link:        body.Link,
		Expired:     body.Expired,
		StartDate:   body.StartDate,
		Term:        body.Term,
	})
	if err != nil {
		return
	}

	for _, statusTag := range body.CreateHackathonRequestBody.StatusTags {
		_, err := hu.store.CreateHackathonStatusTag(ctx, db.CreateHackathonStatusTagParams{HackathonID: hackathon.HackathonID, StatusID: statusTag})
		if err != nil {
			return domain.HackathonResponses{}, err
		}
	}

	statusTags, err := hu.store.GetHackathonStatusTagsByHackathonID(ctx, hackathon.HackathonID)
	if err != nil {
		return
	}

	var tags []db.StatusTags
	for _, statusTag := range statusTags {
		tag, err := hu.store.GetStatusTagByStatusID(ctx, statusTag.StatusID)
		if err != nil {
			return domain.HackathonResponses{}, err
		}
		tags = append(tags, tag)
	}

	result = domain.HackathonResponses{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  tags,
	}
	return
}

func (hu *hackathonUsecase) ListHackathons(ctx context.Context, query domain.ListHackathonsParams) (result []domain.ListHackathonsResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.contextTimeout)
	defer cancel()

	var expired time.Time
	if query.Expired {
		expired = time.Now().Add(time.Hour * 24 * 30 * 6)
	} else {
		expired = time.Now()
	}

	hackathons, err := hu.store.ListHackathons(ctx, db.ListHackathonsParams{
		Expired: expired,
		Limit:   query.PageSize,
		Offset:  (query.PageId - 1) * query.PageSize,
	})
	if err != nil {
		return
	}

	for _, hackathon := range hackathons {
		var tags []db.StatusTags

		statusTags, err := hu.store.GetHackathonStatusTagsByHackathonID(ctx, hackathon.HackathonID)
		if err != nil {
			return nil, err
		}

		for _, statusTag := range statusTags {
			tag, err := hu.store.GetStatusTagByStatusID(ctx, statusTag.StatusID)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}

		result = append(result, domain.ListHackathonsResponses{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			StatusTags:  tags,
		})
	}

	return
}

func (hu *hackathonUsecase) GetHackathon(ctx context.Context, id int32) (result domain.HackathonResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.contextTimeout)
	defer cancel()

	hackathon, err := hu.store.GetHackathonByID(ctx, id)
	if err != nil {
		return domain.HackathonResponses{}, err
	}

	statusTags, err := hu.store.GetHackathonStatusTagsByHackathonID(ctx, hackathon.HackathonID)
	if err != nil {
		return
	}

	var tags []db.StatusTags
	for _, statusTag := range statusTags {
		tag, err := hu.store.GetStatusTagByStatusID(ctx, statusTag.StatusID)
		if err != nil {
			return domain.HackathonResponses{}, err
		}
		tags = append(tags, tag)
	}

	result = domain.HackathonResponses{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  tags,
	}
	return
}