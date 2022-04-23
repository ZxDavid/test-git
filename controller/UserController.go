package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"web.com/ginGormJwt/common"
	"web.com/ginGormJwt/dto"
	"web.com/ginGormJwt/model"
	"web.com/ginGormJwt/response"
	"web.com/ginGormJwt/util"
)

func Register(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须是11位")
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"手机号码必须是11位",
		})*/
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"密码不能小于6位",
		})*/
		return
	}
	//如果名称没有传，随机一个10位的自付串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判断手机是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"用户已经存在",
		})*/
		return
	}
	//创建用户
	//加密用户的密码，不能明文
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		/*c.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"msg":"加密错误",
		})*/
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	log.Println("你好"+name, telephone, password)
	//返回结果
	response.Success(c, nil, "注册成功")
	/*c.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"注册成功",
	})*/
}

func Login(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//数据验证
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"密码不能小于6位",
		})*/
		return
	}

	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须是11位")
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"手机号码必须是11位",
		})*/
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"用户不存在",
		})*/
		return
	}

	//判断密码是否正确 ,失败会有一个err，否则为nil
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, "密码错误", nil)
		/*c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"密码错误",
		})*/
		return
	}

	//发放tiken
	//token := "11"
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		/*c.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"mssage":"系统异常",
		})*/
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
	/*c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"message":"登录成功",
	})*/

}

//获取用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"user": dto.ToUserDto(user.(model.User))}, "登录成功")
	/*c.JSON(http.StatusOK,gin.H{
	"code":200,
	//"data":gin.H{"user":user},
	"data":gin.H{
		"user":dto.ToUserDto(user.(model.User))}})*/
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
