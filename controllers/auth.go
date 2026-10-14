package controllers

import (
	"net/http"

	"studyplanner/config"
	"studyplanner/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// สมัครสมาชิก
func Register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	confirm := c.PostForm("confirm")

	if username == "" || password == "" || confirm == "" {
		c.String(http.StatusBadRequest, "กรุณากรอกข้อมูลให้ครบ")
		return
	}

	if password != confirm {
		c.String(http.StatusBadRequest, "รหัสผ่านไม่ตรงกัน")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.String(http.StatusInternalServerError, "เข้ารหัสรหัสผ่านไม่ได้")
		return
	}

	user := models.User{
		Username: username,
		Password: string(hashPassword),
		Role:     "user",
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.String(http.StatusBadRequest, "Username นี้มีอยู่แล้ว")
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

// เข้าสู่ระบบ
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User

	result := config.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.String(http.StatusUnauthorized, "ไม่พบผู้ใช้")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.String(http.StatusUnauthorized, "รหัสผ่านไม่ถูกต้อง")
		return
	}

	session := sessions.Default(c)

	session.Set("userID", user.ID)
	session.Set("username", user.Username)
	session.Set("role", user.Role)

	err = session.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, "ไม่สามารถบันทึก Session ได้")
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
