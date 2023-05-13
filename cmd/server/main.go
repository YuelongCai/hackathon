package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"hackson/internal/data/conf"
	"hackson/internal/data/config"
	badgetH "hackson/internal/handler/badget"
	userH "hackson/internal/handler/user"
	"hackson/internal/model/rds"
	badgetS "hackson/internal/service/badget"
	userS "hackson/internal/service/user"
)

const (
	confDir = "CONF_DIR"
)

var (
	appConf *conf.AppConf

	rdsConn *rds.Rds

	badgetModel          *rds.BadgetModel
	userBadgetAssetModel *rds.UserBadgetAssetModel

	badgetService *badgetS.Service
	userService   *userS.Service

	badgetHandler *badgetH.Handler
	userHandler   *userH.Handler
)

// export CONF_DIR=configs/local
// go run cmd/server/main.go
func initConf() {
	if e, f := os.LookupEnv(confDir); f {
		log.Printf("load configs from %s\n", e)
		appConf = &conf.AppConf{}
		config.NewConfig(e, "hackson", appConf)
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

	badgetModel = rds.NewBadgetModel(rdsConn)
	userBadgetAssetModel = rds.NewUserBadgetAssetModel(rdsConn)
}

func initService() {
	badgetService = badgetS.NewService(badgetModel)
	userService = userS.NewService(rdsConn, badgetModel, userBadgetAssetModel)
}

func initHandler() {
	badgetHandler = badgetH.NewHandler(badgetService)
	userHandler = userH.NewHandler(userService)
}

func initRoute() {
	r := gin.New()

	r.GET("/api/v1/badget", badgetHandler.List)
	r.POST("/api/v1/badget", badgetHandler.Register)
	r.PUT("/api/v1/badget", badgetHandler.Change)
	r.GET("/api/v1/badget_category", badgetHandler.ListBadgetCategories)
	r.GET("/api/v1/badget_trigger_event", badgetHandler.ListBadgetTriggerEvents)

	r.POST("/api/v1/user_behavior", userHandler.HandleBehavior)
	r.GET("/api/v1/user_badget_asset/:user_id", userHandler.ListBadgetAssets)

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
