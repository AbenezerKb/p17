package initiator

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	//"sso/platform/logger"
	log "github.com/sirupsen/logrus"
	"sms-gateway/platform/pgxAdapter"
)

//InitEnforcer  It initializes casbin service. path is the directory of the config file and
//conn is the pgx adapter.
//It returns the initialized casbin enforcer
func InitEnforcer(path string, conn *pgxpool.Pool) *casbin.Enforcer {
	adapter, err := pgxAdapter.NewAdapterWithDB(conn)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to create adapter: %v", err))
	}

	enforcer, err := casbin.NewEnforcer(path, adapter)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to create enforcer: %v", err))
	}

	return enforcer
}
