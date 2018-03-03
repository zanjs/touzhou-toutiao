package app

import (
	"anla.io/hound/config"
	"anla.io/hound/handler"
	"anla.io/hound/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var (
	appConf = config.Config.APP
	jwtConf = config.Config.JWT
)

// InitApp is
func InitApp() {
	app := iris.New()
	app.Use(crs)
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.Before)

	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(config.NewLogFile())

	app.Options("/*", handler.OptionsHandler)

	app.Get("/", handler.IndexHand)
	app.Get("/init", handler.InitTable)
	app.Get("/suuid", handler.UUID{}.Create)

	v1 := app.Party("/api/v1")
	{
		v1.Post("/login", handler.PostLogin)
		v1.Post("/register", handler.Register{}.Add)
		v1.Get("/", jwtHandler.Serve, handler.Controller{}.JWTHandler)
	}

	Au := v1.Party("/a")
	Op := v1.Party("/o")
	Au.Use(jwtHandler.Serve)

	AuArticle := Au.Party("/article")
	{
		AuArticle.Post("/", handler.Article{}.Create)
	}
	OpAriticle := Op.Party("/article")
	{
		OpAriticle.Get("/", handler.Article{}.All)
	}

	AuCategory := Au.Party("/category")
	{
		AuCategory.Post("/", handler.Category{}.Create)
	}

	AuUpload := Au.Party("/upload")
	{
		AuUpload.Post("/file", handler.UploadFile)
	}

	// navigate to defafult config http://localhost:8080
	if err := app.Run(iris.Addr(":"+appConf.Port), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}
