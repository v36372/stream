package handler

import (
	"clip/repo"
	"stream/app/entity"
	"stream/config"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func InitEngine(conf *config.Config) *gin.Engine {
	if conf.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())

	if conf.App.Debug {
		r.Use(gin.Logger())
	}

	// ----------------------   INIT STATIC
	r.Static("static", "./public")

	templateConfig := gintemplate.TemplateConfig{
		Root:      "public",
		Extension: ".html",
	}

	r.HTMLRender = gintemplate.New(templateConfig)

	// ----------------------   INIT HANDLER

	streamHandler := streamHandler{
		clip: entity.NewClip(repo.Clip),
	}

	// ----------------------   INIT ROUTE

	groupIndex := r.Group("")
	groupIndex.GET("/", streamHandler.Index)

	return r
}

func GET(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	group.GET(relativePath, f)
}

func POST(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	group.POST(relativePath, f)
}

func PUT(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	group.PUT(relativePath, f)
}

func DELETE(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	group.DELETE(relativePath, f)
}
