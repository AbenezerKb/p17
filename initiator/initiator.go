package initiator

import (
	"github.com/gravitational/trace"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
	"os"
	"sms-gateway/initiator/utils"

	//docs "crbt/cmd/rest/docs"
	"github.com/gin-gonic/gin"
)

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

	router := gin.Default()

	routes := router.Group("/v1")

	// Register domains start
	UserDomainInit(routes, common)

	log.WithFields(log.Fields{
		"host": os.Getenv("SERVER_HOST"),
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_HOST"), router))
}
