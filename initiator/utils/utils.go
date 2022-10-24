package utils

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	init_const "sms-gateway/internal/constant/init"
	"sms-gateway/platform/pgxAdapter"

	// model "crbt/internal/constant/model/db"
	"log"
	"os"
	"strconv"
	"time"

	"sms-gateway/platform/httpclient"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func GetUtils(dbUrl string, authModel string) (init_const.Utils, error) {

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      true,        // Disable color
	// 	},
	// )
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}
	conn, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}

	trans, validate, err := GetValidation()
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}

	duration, _ := strconv.Atoi(os.Getenv("timeout"))
	timeoutContext := time.Duration(duration) * time.Second

	httpClient := httpclient.InitHttpClient()
	enforcer := NewEnforcer(conn, authModel)

	return init_const.Utils{
		Timeout:     timeoutContext,
		Translator:  trans,
		GoValidator: validate,
		Conn:        conn,
		Enforcer:    enforcer,
		HttpClient:  httpClient,
	}, nil
}

func NewEnforcer(conn *pgxpool.Pool, authModel string) *casbin.Enforcer {

	adapter, err := pgxAdapter.NewAdapterWithDB(conn)
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}

	enforcer, err := casbin.NewEnforcer(authModel, adapter)
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}

	enforcer.EnableAutoSave(true)
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}
	return enforcer
}
