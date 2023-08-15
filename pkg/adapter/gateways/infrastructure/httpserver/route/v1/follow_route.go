package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

// フォローのルーティングを定義する
func NewFollowRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group gin.IRoutes) {
	followController := controller.FollowController{
		FollowUsecase: usecase.NewFollowUsercase(store, timeout),
		Env:           env,
	}

	group.GET("/accounts/:account_id/follow", followController.GetFollow)
	group.POST("/accounts/:account_id/follow", followController.CreateFollow)
	group.DELETE("/acccounts/:account_id/follow	", followController.RemoveFollow)
}
