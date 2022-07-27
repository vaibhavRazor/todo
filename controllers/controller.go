package controllers

import (
	"fmt"
	"todo/database"
	"todo/model"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Repo struct {
	DbMap *gorp.DbMap
}

func init() {
	//prometheus.MustRegister(counter_new)
}

func New() *Repo {
	db := database.InitDb()
	return &Repo{DbMap: db}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func (repo *Repo) ListTasks(c *gin.Context) {
	counter_new.Inc()
	c.JSON(200, "hello")
	// task, err := model.ListTasks(c, repo.DbMap)
	// if err == nil {
	// 	c.JSON(200, task)
	// } else {
	// 	c.JSON(200, err)
	// }

}

func (repo *Repo) AddTask(c *gin.Context) {
	var task model.Task
	c.BindJSON(&task)
	id, err := model.AddTask(c, task, repo.DbMap)
	if err != nil {
		c.JSON(500, "Someone error occurred")
	}
	var msg = fmt.Sprintf("record inserted successfully with id : %d", id)
	c.JSON(200, msg)
}

func UpdateTask(c *gin.Context) {

}

func DeleteTask(c *gin.Context) {

}

var counter_new = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "default",
	Name:      "leagues_request_counter",
	Help:      "Number of requests",
})
