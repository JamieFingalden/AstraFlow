package jwt

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明结构
// 定义JWT Token中包含的用户信息和标准声明
type Claims struct {
	UserID               int64  `json:"user_id"`  // 用户ID
	Username             string `json:"username"` // 用户名
	Role                 string `json:"role"`     // 用户角色 (Key)
	TenantID             *int64 `json:"tenant_id,omitempty"`
	jwt.RegisteredClaims        // JWT标准声明
}

// GenerateToken 生成JWT Token
// 根据用户信息生成访问令牌，过期时间为24小时
func GenerateToken(user *model.User) (string, error) {
	// 设置过期时间（24小时）
	expirationTime := time.Now().Add(24 * time.Hour)

	roleKey := ""
	if user.Role != nil {
		roleKey = user.Role.Key
	}

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     roleKey,
		TenantID: user.TenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "astraflow", // 签发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.Server.JwtSecret))
}

// ParseToken 解析JWT Token
// 验证并解析JWT Token，返回用户声明信息
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Server.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// GenerateRefreshToken 生成刷新Token
// 根据用户信息生成刷新令牌，过期时间为7天
func GenerateRefreshToken(user *model.User) (string, error) {
	// 设置过期时间（7天）
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	roleKey := ""
	if user.Role != nil {
		roleKey = user.Role.Key
	}

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     roleKey,
		TenantID: user.TenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "astraflow_refresh", // 刷新令牌签发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.Server.JwtSecret))
}