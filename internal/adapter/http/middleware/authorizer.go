package middleware

import (
	"github.com/casbin/casbin/v2"
	// "crbt/internal/constant/model"
	errors "sms-gateway/internal/constant/rest/error_types"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func newAuthorizer(enforcer *casbin.Enforcer, logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// JWT Authentication
		const bearerSchema = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= len(bearerSchema) {

			err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
			logger.Error(zap.Error(err))
			_ = ctx.Error(err)
			ctx.Abort()
			return
		}

		tokenString := authHeader[len(bearerSchema):]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
				return nil, errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
			}
			return []byte("secretKey"), nil
		})

		if err != nil || !token.Valid {
			err = errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
			logger.Error(zap.Error(err))
			_ = ctx.Error(err)
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("userID", claims["Phone"])
		logger.Infof("New login: userID: %s", claims["Phone"])

		// Casbin Authorization
		object := ctx.Request.URL
		action := ctx.Request.Method
		ok, err := enforcer.Enforce(claims["Phone"], object, action)
		if !ok || err != nil {
			err = errors.ErrForbiddenMethod.New(errors.ErrorUnauthorizedError)
			logger.Error(zap.Error(err))
			_ = ctx.Error(err)
		}

	}
}
