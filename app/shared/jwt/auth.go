package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

type Claims struct {
	UserID   int64     `json:"user_id"`
	DeviceID string    `json:"device_id"`
	Platform string    `json:"platform"`
	IAT      time.Time `json:"iat"`
	Openid   string    `json:"openid"`
}

func NewWithClaims(userID int64, deviceID, platform, openid string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,   // 用户id
		"device_id": deviceID, // 设备id
		"platform":  platform,
		"iat":       time.Now().Unix(), // 签发时间
		"openid":    openid,
	})
}

func GetClaims(token *jwt.Token) *Claims {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &Claims{}
	}
	return &Claims{
		UserID:   cast.ToInt64(claims["user_id"]),
		DeviceID: cast.ToString(claims["device_id"]),
		Platform: cast.ToString(claims["platform"]),
		IAT:      time.Unix(cast.ToInt64(claims["iat"]), 0),
		Openid:   cast.ToString(claims["openid"]),
	}
}

func Parse(token, secret string) (*jwt.Token, error) {
	token = strings.TrimPrefix(token, "Bearer ")
	return jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
