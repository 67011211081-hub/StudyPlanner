package controllers

import (
	"net/http"

	"studyplanner/config"
	"studyplanner/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// หน้า Admin
func AdminPage(c *gin.Context) {

	session := sessions.Default(c)

	role := session.Get("role")

	if role == nil || role != "admin" {
		c.String(http.StatusForbidden, "ไม่มีสิทธิ์เข้าหน้านี้")
		return
	}

	var users []models.User
	var tasks []models.Task

	config.DB.Find(&users)
	config.DB.Find(&tasks)

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"users": users,
		"tasks": tasks,
	})
}

// Admin ลบงาน
func DeleteTaskAdmin(c *gin.Context) {

	session := sessions.Default(c)

	role := session.Get("role")

	if role == nil || role != "admin" {
		c.String(http.StatusForbidden, "ไม่มีสิทธิ์")
		return
	}

	id := c.Param("id")

	config.DB.Delete(&models.Task{}, id)

	c.Redirect(http.StatusFound, "/admin")
}
