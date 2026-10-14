package main

import (
	"os"
	"studyplanner/config"
	"studyplanner/controllers"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {

	// เชื่อมฐานข้อมูล
	config.ConnectDatabase()

	// สร้าง Gin Engine
	r := gin.Default()

	store := cookie.NewStore([]byte("studyplanner-secret-key"))
	r.Use(sessions.Sessions("studyplanner-session", store))

	// โหลดไฟล์ HTML
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// หน้า Register
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	// หน้า Login
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	// สมัครสมาชิก
	r.POST("/register", controllers.Register)

	// เข้าสู่ระบบ
	r.POST("/login", controllers.Login)

	r.GET("/dashboard", controllers.Dashboard)

	r.POST("/task/create", controllers.CreateTask)

	r.GET("/task/delete/:id", controllers.DeleteTask)

	r.GET("/task/edit/:id", controllers.EditTask)

	r.POST("/task/update/:id", controllers.UpdateTask)

	r.GET("/admin", controllers.AdminPage)

	r.GET("/logout", controllers.Logout)

	r.GET("/admin/delete/:id", controllers.DeleteTaskAdmin)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
