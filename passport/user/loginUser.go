package user

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/gongyao1992/go-util/passport/jwt"
	"time"
)

// CachedInfo ...
type CachedInfo struct {
	User User   `json:"user"`
	Sign string `json:"sign"`

	ExpTime    int64 `json:"exp_time"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

// User ...
type User struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Mobile string   `json:"mobile"`
	Role   []string `json:"role"`

	Status      string `json:"status"`
	EnterStatus int64  `json:"enter_status"`

	IsEnable string `json:"is_enable"`
	IsAdmin  int64  `json:"is_admin"`
}

const USER_AUTH = "user-auth"

// GetLoginUser ..
func GetLoginUser(rdb *redis.Client, token, confHeader, confKey string) (user User, err error) {
	j, err := jwt.Decode(token, confHeader, confKey)
	if err != nil {
		return
	}
	key := USER_AUTH
	byt, err := rdb.HGet(key, j.ID).Bytes()
	if err != nil {
		return user, errors.New("用户信息错误")
	}
	var info CachedInfo
	if err = json.Unmarshal(byt, &info); err != nil {
		return
	}
	now := time.Now().Unix()
	if info.ExpTime < now {
		return user, errors.New("登录信息超时")
	}
	return info.User, nil
}
