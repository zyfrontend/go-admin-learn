package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/sync/singleflight"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token 已过期")
	TokenNotValidYet = errors.New("token 尚未激活")
	TokenMalformed   = errors.New("这甚至不是一个 Token")
	TokenInvalid     = errors.New("无法处理此 Token：")
)

var (
	SigningKey  = "da5295ca-49b6-492d-9fee-490c1af10534" // jwt签名
	ExpiresTime = "7d"                                   // 过期时间
	BufferTime  = "1d"                                   // 缓冲时间
	Issuer      = "qmPlus"
)
var GVA_Concurrency_Control = &singleflight.Group{}

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}

func NewJWT(key string) *JWT {
	return &JWT{
		[]byte(key),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := ParseDuration(BufferTime)
	ep, _ := ParseDuration(ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(), // 过期时间 7天  配置文件
			Issuer:    Issuer,                    // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func SetRedisJwt(jwt string, address string) (err error) {
	dr, err := ParseDuration(ExpiresTime)
	if err != nil {
		return err
	}
	err = Redis.Set(context.Background(), address, jwt, dr).Err()
	return err
}
