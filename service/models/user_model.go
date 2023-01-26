package models

import (
	"app/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type AppUser struct {
	gorm.Model
	Username        string `json:"userName" gorm:"index;comment:用户名"`
	Password        string `json:"password" gorm:"comment:密码"`
	Address         string `json:"address" gorm:"comment:钱包地址"`
	Salt            string `json:"salt" gorm:"comment:加密串"`
	Chain           string `json:"chain" gorm:"comment:安全密码"`
	Amount          int    `json:"amount" gorm:"comment:用户余额"`
	PayU            int    `json:"payU" gorm:"comment:开启机器人投资Usdt余额"`
	AmountUsdt      int    `json:"amountUsdt" gorm:"comment:lp添池Usdt余额"`
	AmountLp        int    `json:"amountLp" gorm:"comment:开启机器人Lp量"`
	Level           int    `json:"level" gorm:"comment:等级"`
	RecommendUserId int    `json:"recommendUserId" gorm:"comment:推荐人id"`
	IsBroker        int    `json:"isBroker" gorm:"comment:是否矿工"`
	IsMember        int    `json:"isMember" gorm:"comment:是否会员"`
	Status          int    `json:"status" gorm:"comment:状态"`
}

type SelectInfo struct {
	Id       string
	Username string
	Address  string
}

type List struct {
	Limit  int `json:"limit" default:"10"`
	Offset int `json:"offset" default:"0"`
}

type SelectStruct struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Address  string `json:"address"`
}

func (AppUser) TableName() string {
	return "app_users"
}

// QueryUserFromUserName 查询用户
func QueryUserFromUserName(username string) (appUser *AppUser, err error) {
	user := AppUser{}
	err = utils.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, err
}
func QueryUserFromID(id string) (appUser *AppUser, err error) {
	user := AppUser{}
	err = utils.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	fmt.Println(user)
	return &user, err
}
func QueryUserFromAddress(address string) (appUser *AppUser, err error) {
	user := AppUser{}
	err = utils.DB.Where("address = ?", address).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, err
}

func QueryUserInfo(config SelectInfo) (appUser *AppUser, err error) {
	user := AppUser{}
	err = utils.DB.Where("address = ?", config.Address).Or("id = ?", config.Id).Or("username = ?", config.Username).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, err
}

// CreateUser 创建用户
func CreateUser(user AppUser) *gorm.DB {
	return utils.DB.Create(&user)
}

// GetUserList 获取用户列表
func GetUserList(p List) (appUser []AppUser, err error) {
	var user []AppUser
	err = utils.DB.Limit(p.Limit).Offset(p.Offset).Find(&user).Error
	if err != nil {
		return nil, errors.New("用户列表查询错误")
	}
	return user, nil
}
