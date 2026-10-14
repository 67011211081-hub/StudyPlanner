package controllers

import (
	"net/http"

	"studyplanner/config"
	"studyplanner/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// แสดงหน้ารายการงาน
func Dashboard(c *gin.Context) {

	session := sessions.Default(c)

	userID := session.Get("userID")
	username := session.Get("username")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var tasks []models.Task

	config.DB.Where("user_id = ?", userID).Find(&tasks)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"username": username,
		"tasks":    tasks,
	})
}

// เพิ่มงาน
func CreateTask(c *gin.Context) {

	session := sessions.Default(c)

	userID := session.Get("userID")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")

	task := models.Task{
		Title:       title,
		Description: description,
		UserID:      userID.(uint),
	}

	config.DB.Create(&task)

	c.Redirect(http.StatusFound, "/dashboard")
}

// ลบงาน
func DeleteTask(c *gin.Context) {

	session := sessions.Default(c)

	userID := session.Get("userID")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id := c.Param("id")

	var task models.Task

	result := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task)

	if result.Error != nil {
		c.String(http.StatusForbidden, "ไม่มีสิทธิ์ลบงานนี้")
		return
	}

	config.DB.Delete(&task)

	c.Redirect(http.StatusFound, "/dashboard")
}

// แสดงหน้าแก้ไขงาน
func EditTask(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("userID")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id := c.Param("id")

	var task models.Task

	result := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task)

	if result.Error != nil {
		c.String(http.StatusForbidden, "ไม่มีสิทธิ์")
		return
	}

	c.HTML(http.StatusOK, "edit_task.html", gin.H{
		"task": task,
	})
}

// บันทึกการแก้ไข
func UpdateTask(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("userID")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id := c.Param("id")

	var task models.Task

	result := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task)

	if result.Error != nil {
		c.String(http.StatusForbidden, "ไม่มีสิทธิ์")
		return
	}

	task.Title = c.PostForm("title")
	task.Description = c.PostForm("description")

	config.DB.Save(&task)

	c.Redirect(http.StatusFound, "/dashboard")
}
