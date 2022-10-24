package const_init

import (
	"github.com/casbin/casbin/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"

	"sms-gateway/platform/httpclient"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type Utils struct {
	Conn        *pgxpool.Pool
	GoValidator *validator.Validate
	Translator  ut.Translator
	Enforcer    *casbin.Enforcer
	Timeout     time.Duration
	HttpClient  httpclient.HttpClient
}
