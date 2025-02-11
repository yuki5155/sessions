package main

import (
	"gin-contrib/sessions"
	"sessions/cookie"
	"github.com/gin-gonic/gin"
  )
  
  func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
  
	r.GET("/hello", func(c *gin.Context) {
	  session := sessions.Default(c)
  
	  if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	  }
  
	  c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
  }