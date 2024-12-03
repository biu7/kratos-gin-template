package middleware

import (
	"errors"
	"github.com/biu7/gokit/ctxvalue"
	"github.com/biu7/gokit/ginutils"
	"github.com/biu7/gokit/ginutils/response"
	"github.com/biu7/gokit/log"
	"kratos-gin-template/app/server/internal/constants"
	"kratos-gin-template/app/shared/jwt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrMissingToken = errors.New("missing token")
	ErrInvalidToken = errors.New("invalid token")
	ErrNeedLogin    = errors.New("need login")
)

type Auth struct {
	jwtSecret string
	log       log.Logger
}

// Anonymous 不强制要求有 token，只是为了解析 token 中的信息
func (a *Auth) Anonymous() gin.HandlerFunc {
	return func(c *gin.Context) {
		if tokenStr := ginutils.GetToken(c); tokenStr != "" {
			token, _ := jwt.Parse(tokenStr, a.jwtSecret)
			if token != nil {
				injectRequestData(c, jwt.GetClaims(token))
			}
		}

		c.Next()
	}
}

func (a *Auth) Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := ginutils.GetToken(c)
		token, err := jwt.Parse(tokenStr, a.jwtSecret)
		if err != nil {
			a.log.Ctx(c).Error("parse token failed", "err", err, "token", c.Request.Header.Get("Authorization"))
			response.AuthFail(c, ErrInvalidToken)
			c.Abort()
			return
		}
		if !token.Valid {
			a.log.Ctx(c).Error("parse token success, but not valid", "err", err, "token", c.Request.Header.Get("Authorization"))
			response.AuthFail(c, ErrInvalidToken)
			c.Abort()
			return
		}

		injectRequestData(c, jwt.GetClaims(token))

		c.Next()
	}
}

func (a *Auth) User() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := ginutils.GetToken(c)
		token, err := jwt.Parse(tokenStr, a.jwtSecret)
		if err != nil {
			a.log.Ctx(c).Error("parse token failed", "err", err, "token", c.Request.Header.Get("Authorization"))
			response.AuthFail(c, ErrInvalidToken)
			c.Abort()
			return
		}
		if !token.Valid {
			a.log.Ctx(c).Error("parse token success, but not valid", "err", err, "token", c.Request.Header.Get("Authorization"))
			response.AuthFail(c, ErrInvalidToken)
			c.Abort()
			return
		}
		claims := jwt.GetClaims(token)
		if claims.UserID <= 0 {
			response.AuthFail(c, ErrNeedLogin)
			c.Abort()
			return
		}

		injectRequestData(c, claims)

		c.Next()
	}
}

func injectRequestData(c *gin.Context, claims *jwt.Claims) {
	ctxvalue.SetUserID(c, claims.UserID)
	ctxvalue.SetDeviceID(c, claims.DeviceID)
	ctxvalue.SetPlatform(c, claims.Platform)

	// 从请求头中获取版本号
	ctxvalue.SetContextVersion(c, unpackAppVersion(c))

	// 从请求头中获取微信 openid
	openid := c.GetHeader(constants.APPHeaderWechatOpenid)
	if openid == "" && claims.Openid != "" {
		openid = claims.Openid
	}
	ctxvalue.SetContextWechatOpenID(c, openid)
}

func unpackAppVersion(c *gin.Context) int32 {
	version := c.GetHeader(constants.APPHeaderVersion)
	version = strings.TrimPrefix(version, "0")
	v, _ := strconv.ParseInt(version, 10, 32)
	return int32(v)
}
