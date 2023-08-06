package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type FollowController struct {
	FollowUsecase domain.FollowUsecase
	Env           *bootstrap.Env
}

// TODO:レスポンス変更　=> accounts
// CreateFollow	godoc
// @Summary			Create Follow
// @Description		Follow!!!!!!!!
// @Tags			Accounts
// @Produce			json
// @Param			from_user_id 				path 		string						true	"Accounts API wildcard"
// @Param			CreateFollowRequestBody 	body 		CreateFollowRequestBody		true	"create Follow Request Body"
// @Success			200							{array}		db.Follows					"succsss response"
// @Failure 		400							{object}	ErrorResponse				"error response"
// @Failure 		500							{object}	ErrorResponse				"error response"
// @Router       	/accounts/:from_user_id/follow			[post]
func (fc *FollowController) CreateFollow(ctx *gin.Context) {
	var (
		reqURI  domain.AccountRequestWildCard
		reqBody domain.CreateFollowRequestBody
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := fc.FollowUsecase.CreateFollow(ctx, db.CreateFollowParams{ToUserID: reqBody.ToUserID, FromUserID: reqURI.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// TODO:レスポンス修正
// RemoveFollow	godoc
// @Summary			Remove follow
// @Description		Unfollow
// @Tags			Accounts
// @Produce			json
// @Param			from_user_id 				path 		string						true	"Accounts API wildcard"
// @Param			RemoveFollowRequestQueries 	formData 	CreateFollowRequestBody		true	"Remove Follow Request Body"
// @Success			200							{object}	DeleteResponse				"succsss response"
// @Failure 		400							{object}	ErrorResponse				"error response"
// @Failure 		500							{object}	ErrorResponse				"error response"
// @Router       	/accounts/:from_user_id/follow			[delete]
func (fc *FollowController) RemoveFollow(ctx *gin.Context) {
	var (
		reqURI   AccountRequestWildCard
		reqQuery domain.RemoveFollowRequestQueries
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := fc.FollowUsecase.RemoveFollow(ctx, db.RemoveFollowParams{ToUserID: reqQuery.ToUserID, FromUserID: reqURI.ID}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, DeleteResponse{Result: "Delete Successful"})
}

func (fc *FollowController) GetFollow(ctx *gin.Context) {
	var (
		reqURI   AccountRequestWildCard
		reqQuery domain.GetFollowRequestQueries
		result   []domain.FollowResponse
		err      error
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO:　ToFollowからの取得と FromFollowからの取得　両方作る
	if reqQuery.Mode {
		result, err = fc.FollowUsecase.GetFollowByToID(ctx, reqURI.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		// Fromの取得
	}
	ctx.JSON(http.StatusOK, result)
}
