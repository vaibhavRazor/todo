package mapping

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Router *gin.Engine
var repo = controllers.New()
var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// v1 of the API
	{
		Router.GET("/task/list", repo.ListTasks)

		Router.POST("/task/add", repo.AddTask)
		Router.POST("/task/update", controllers.UpdateTask)
		Router.PUT("/task/delete", controllers.DeleteTask)
		Router.GET("/metrics", prometheusHandler())

	}

}
func prometheusHandler() gin.HandlerFunc {
	cpuTemp.Set(65.3)
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
