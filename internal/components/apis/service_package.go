package apis

import (
	_ "airplane/docs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"runtime"
	"time"
)

func HealthPackage(group *gin.RouterGroup) {
	// health check
	group.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "health: "+time.Now().Format(time.RFC3339Nano))
	})
}

// PrometheusPackage is a handler for prometheus metrics
func PrometheusPackage(group *gin.RouterGroup) {
	prometheusHandler := promhttp.Handler()
	group.GET("/metrics", func(ctx *gin.Context) {
		prometheusHandler.ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func PprofPackage(group *gin.RouterGroup) {
	// pprof
	pprofGroup := group.Group("")
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	pprof.Register(pprofGroup, "/debug")
}

func ApiDocPackage(group *gin.RouterGroup) {
	// swagger
	group.GET("/redoc", func(ctx *gin.Context) {
		const redocIndexHTML = `<!DOCTYPE html>
        <html>
          <head>
            <!-- needed for adaptive design -->
            <meta charset="utf-8"/>
            <meta name="viewport" content="width=device-width, initial-scale=1">
            <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
        
            <!--
            ReDoc doesn't change outer page styles
            -->
            <style>
              body {
                margin: 0;
                padding: 0;
              }
            </style>
          </head>
          <body>
            <redoc spec-url='/_/docs/doc.json'></redoc>
            <script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
          </body>
        </html>`
		_, _ = ctx.Writer.Write([]byte(redocIndexHTML))
	})
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
