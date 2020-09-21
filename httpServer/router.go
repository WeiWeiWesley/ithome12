package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//Router
const (
	RouterRedirectPrefix = "/api/redirect"
	AccessLog            = "access_log"
)

//Server Server
type Server struct {
	Router *gin.Engine
	Port   string
}

//Router api router
func Router() (server []Server) {
	if os.Getenv("ENV") != LOCAL {
		gin.SetMode(gin.ReleaseMode) //trun off gin debug log
	}

	r := gin.New()
	r.Use(
		gin.Recovery(),     // Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
		crosMiddleware(),   // Allow Header
		accessMiddleware(), // Some important access log
	)

	//open gin access log
	if os.Getenv("ENV") == LOCAL {
		r.Use(gin.Logger())
	}

	server = append(server, Server{
		Router: r,
		Port:   ":80",
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/echo", Echo)
	r.GET("/book/:page", BookPage)
	r.POST("/body/json", JSONBody)
	r.POST("/body/form_data", FormData)

	return
}

// Header Allow
func crosMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
			"Accept-Encoding, X-CSRF-Token, Authorization, token, Accept, Referer, Origin, User-Agent")

		if c.Request.Method == "OPTIONS" {
			c.String(200, "ok")
			return
		}
		c.Next()
	}
}

// Some important access log
func accessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, RouterRedirectPrefix) {
			accseeLog := (map[string]interface{}{
				AccessLog: map[string]interface{}{
					"address":        c.Request.Header.Get("X-Forwarded-For"),
					"method":         c.Request.Method,
					"uri":            c.Request.RequestURI,
					"header":         c.Request.Header,
					"content_length": c.Request.ContentLength,
					"body":           "",
				},
			})
			log.Println(accseeLog)
		}
	}
}
