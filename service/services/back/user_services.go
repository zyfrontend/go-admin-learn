package back

import (
	"app/common"
	"app/models"
	"app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserService struct {
}

type FormJson struct {
	UserName string `json:"userName"`
	Password string `json:"Password"`
}

type Res struct {
	Token    string         `json:"token"`
	UserInfo models.AppUser `json:"userInfo"`
}

// Register 注册
func (u *UserService) Register(c *gin.Context) {
	user := models.AppUser{}
	formJson := FormJson{}
	err := c.ShouldBindJSON(&formJson)
	if err != nil {
		common.FailMsg(err.Error(), c)
		return
	}
	fmt.Println(formJson.UserName)
	// 查询用户是否存在
	res, err := models.QueryUserFromUserName(formJson.UserName)
	if res != nil {
		common.FailMsg("用户已存在", c)
		return
	}
	// 用户不存在，创建用户
	user.Username = formJson.UserName
	user.Password = formJson.Password
	models.CreateUser(user)
	common.SuccessDataMsg(user, "注册成功", c)

}

// Login 登录
func (u *UserService) Login(c *gin.Context) {
	formJson := FormJson{}
	info := Res{}
	err := c.ShouldBindJSON(&formJson)
	if err != nil {
		common.FailMsg(err.Error(), c)
		return
	}
	user := &models.AppUser{}
	// 查询是否存在
	user, err = models.QueryUserFromUserName(formJson.UserName)
	if err != nil {
		common.FailMsg(err.Error(), c)
		return
	}
	// 绑定推荐人
	// 派发token
	j := &utils.JWT{SigningKey: []byte(utils.SigningKey)} // 唯一签名
	claims := j.CreateClaims(utils.BaseClaims{
		ID:          user.ID,
		NickName:    user.Username,
		Username:    user.Username,
		AuthorityId: user.ID,
	})
	token, err := j.CreateToken(claims)
	err = utils.SetRedisJwt(token, formJson.UserName)
	if err != nil {
		return
	}
	info.UserInfo = *user
	info.Token = token
	common.SuccessDataMsg(info, "登录成功", c)
}

// QueryUserInfo 用户信息
func (u *UserService) QueryUserInfo(c *gin.Context) {
	data := &models.AppUser{}
	var selectConfig = models.SelectInfo{}
	err := c.ShouldBindJSON(&selectConfig)
	if err != nil {
		common.FailMsg(err.Error(), c)
		return
	}
	data, err = models.QueryUserInfo(selectConfig)
	if err != nil {
		common.FailMsg("查找失败", c)
		return
	}
	common.SuccessDataMsg(data, "查找成功", c)
}

// GetUserList 用户列表
func (u *UserService) GetUserList(c *gin.Context) {
	var err error
	data := make([]models.AppUser, 10)
	var p = models.List{}
	p.Limit, _ = strconv.Atoi(c.Query("limit"))
	p.Offset, _ = strconv.Atoi(c.Query("offset"))
	data, err = models.GetUserList(p)
	if err != nil {
		common.FailMsg(err.Error(), c)
		return
	}
	fmt.Println("data", data)
	common.SuccessDataMsg(data, "查询成功", c)
}
