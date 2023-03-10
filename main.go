package main

import (
	"github.com/gin-gonic/gin"
	"loadtest-locust/cluster"
	"loadtest-locust/contrib"
	"loadtest-locust/dao"
	"loadtest-locust/db"
	"loadtest-locust/user"
	"loadtest-locust/util"

	"github.com/gin-contrib/pprof"
)

func init() {
	go util.InitLog()
	go contrib.InitLocust()
	db.RedisInit()
	db.Connect()
	go cluster.HealthCheck()
	go dao.GetHostName()
	//db.EtcdInit()
}

func main() {
	g := gin.Default()
	g.Use(gin.Recovery())
	pprof.Register(g)
	g.POST("/create_task", user.CreateTask)
	g.POST("/start_task", user.StartTask)
	g.POST("/stop_task", user.StopTask)
	g.GET("/task_list", user.TaskList)
	g.GET("/test", user.TestApi)
	_ = g.Run(":9999")
}
