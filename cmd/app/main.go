package main

import (
	"fmt"
	"os"

	"github.com/gimmickless/keat-kit-service/internal/app"
	egdb "github.com/gimmickless/keat-kit-service/internal/transport/egress/db"
	inhttp "github.com/gimmickless/keat-kit-service/internal/transport/ingress/http"
	applog "github.com/gimmickless/keat-kit-service/pkg/logging"
	"github.com/gofiber/fiber/v2"
	httplogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var (
	port = os.Getenv("HTTP_PORT")
)

func main() {
	logger := applog.NewLogger()

	// Load .env file (only expected on local workstations)
	if err := godotenv.Load(".env"); err != nil {
		logger.Debugf(".env file could not be loaded (only harmful when running as standalone on local workstations).")
	}

	// Initialize the db connection
	db, cancel, disconnect := initdb()
	defer cancel()
	defer disconnect()

	// Init and bind projects
	var (
		catgRepo   = egdb.NewCategoryRepository(logger, db)
		ingredRepo = egdb.NewIngredientRepository(logger, db)
		kitsRepo   = egdb.NewKitRepository(logger, db)
		catgSrv    = app.NewCategoryService(logger, catgRepo)
		ingredSrv  = app.NewIngredientService(logger, ingredRepo)
		kitSrv     = app.NewKitService(logger, kitsRepo)
		handler    = inhttp.NewHTTPHandler(logger, catgSrv, ingredSrv, kitSrv)
	)

	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(httplogger.New())

	inhttp.Register(app, handler)

	app.Listen(fmt.Sprintf(":%s", port))
}
