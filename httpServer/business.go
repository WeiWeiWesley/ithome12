package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Echo print param that request carried
func Echo(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(http.StatusPaymentRequired, "name required")
		return
	}
	age := c.Query("age")

	c.String(http.StatusOK, fmt.Sprintf("Receive name: %s, age: %s", name, age))
}

//BookPage get param in url
func BookPage(c *gin.Context) {
	page := c.Param("page")
	c.String(http.StatusOK, fmt.Sprintf("Now we got page: %s", page))
}

//JSONBody receive body with row json
func JSONBody(c *gin.Context) {
	type request struct {
		Category string `json:"category"`
		Page     int    `json:"page"`
	}

	var req request
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}
	fmt.Println("category", req.Category, "page", req.Page)

	c.JSON(http.StatusOK, req)
}

//FormData receive body with form-data
func FormData(c *gin.Context) {
	category := c.PostForm("category")
	page := c.PostForm("page")

	pInt, err := strconv.Atoi(page)
	if err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}

	resp := struct {
		Category string `json:"category"`
		Page     int    `json:"page"`
	}{category, pInt}

	c.JSON(http.StatusOK, resp)
}
