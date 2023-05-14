package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"hackathon/internal/data/conf"
	"hackathon/internal/data/config"
	badgeH "hackathon/internal/handler/badge"
	userH "hackathon/internal/handler/user"
	"hackathon/internal/model/rds"
	badgeS "hackathon/internal/service/badge"
	userS "hackathon/internal/service/user"
)

const (
	confDir = "CONF_DIR"
)

var (
	appConf *conf.AppConf

	rdsConn *rds.Rds

	badgeModel          *rds.BadgeModel
	userBadgeAssetModel *rds.UserBadgeAssetModel

	badgeService *badgeS.Service
	userService  *userS.Service

	badgeHandler *badgeH.Handler
	userHandler  *userH.Handler
)

// export CONF_DIR=configs/local
// go run cmd/server/main.go
func initConf() {
	if e, f := os.LookupEnv(confDir); f {
		log.Printf("load configs from %s\n", e)
		appConf = &conf.AppConf{}
		config.NewConfig(e, "hackathon", appConf)
		log.Printf("load conf success, %+v\n", appConf)
	} else {
		panic("Error on loading configs, please feed CONF_DIR environment!")
	}
}

func initDB() {
	var err error
	rdsConn, err = rds.New(appConf.MySQL.DSN())
	if err != nil {
		panic(err)
	}

	badgeModel = rds.NewBadgeModel(rdsConn)
	userBadgeAssetModel = rds.NewUserBadgeAssetModel(rdsConn)
}

func initService() {
	badgeService = badgeS.NewService(badgeModel)
	userService = userS.NewService(rdsConn, badgeModel, userBadgeAssetModel)
}

func initHandler() {
	badgeHandler = badgeH.NewHandler(badgeService)
	userHandler = userH.NewHandler(userService)
}

func initRoute() {
	r := gin.New()

	r.GET("/api/v1/badge", badgeHandler.List)
	r.POST("/api/v1/badge", badgeHandler.Register)
	r.PUT("/api/v1/badge", badgeHandler.Change)
	r.GET("/api/v1/badge_category", badgeHandler.ListBadgeCategories)
	r.GET("/api/v1/badge_trigger_event", badgeHandler.ListBadgeTriggerEvents)

	r.POST("/api/v1/user_behavior", userHandler.HandleBehavior)
	r.GET("/api/v1/user_badge_asset/:user_id", userHandler.ListBadgeAssets)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func main() {
	initConf()
	initDB()
	initService()
	initHandler()
	initRoute()
}
