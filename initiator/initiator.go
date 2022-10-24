package initiator

import (
	"github.com/swaggo/swag/example/basic/docs"
	"net/http"
	"os"
	"sms-gateway/initiator/utils"
	"sms-gateway/internal/adapter/http/middleware"

	"github.com/gravitational/trace"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"

	//docs "crbt/cmd/rest/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

const (
	authModel = "config/rbac_with_domain_model.conf"
)

func Initialize() {
	DATABASE_URL, err := utils.DbConnectionString()

	lg, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error while initializing logger: %v", zap.Error(err))
	}
	logger := lg.Sugar()
	logger.Info("logger initialized")

	if err != nil {
		log.Fatal("database connection failed!")
	}
	hook, herr := trace.NewUDPHook()
	if err != nil {
		panic(herr)
	}
	log.AddHook(hook)

	common, er := utils.GetUtils(DATABASE_URL, authModel)
	if err != nil {
		log.Fatal(er)
	}
	//casbin initialization
	log.Info("initializing casbin")
	enforcer := InitEnforcer(authModel, common.Conn)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes := router.Group("/v1")
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// mw := middleware(common)
	routes.Use(middleware.ErrorHandling())
	// routes.Use(middleware.NewAuthorizer(logger))

	// Register domains start
	UserDomainInit(routes, common)
	ClientDomainInit(enforcer, routes, common, logger)
	//SubscriptionDomainInit(routes, common)
	log.WithFields(log.Fields{
		"host": os.Getenv("SERVER_HOST"),
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_HOST"), router))
}
