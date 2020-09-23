package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Echo print param that request carried
curl --location --request GET 'localhost/echo?name=weiweiwesley&age=87'
*/
func Echo(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(http.StatusPaymentRequired, "name required")
		return
	}
	age := c.Query("age")

	c.String(http.StatusOK, fmt.Sprintf("Receive name: %s, age: %s", name, age))
}

/*
BookPage get param in url
curl --location --request GET 'localhost/book/3'
*/
func BookPage(c *gin.Context) {
	page := c.Param("page")
	c.String(http.StatusOK, fmt.Sprintf("Now we got page: %s", page))
}

/*
JSONBody receive body with row json
curl --location --request POST '127.0.0.1/body/form_data' \
--header 'Content-Type: application/json' \
--data-raw '{
    "category": "math",
    "page": 32
}'
*/
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

/*
FormData receive body with form-data
curl --location --request POST '127.0.0.1/body/form_data' \
--form 'category=statistic' \
--form 'page=32'
*/
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

/*
GetFilmList SELECT gorm.film data from mysql
curl --location --request GET 'localhost/film/list'
*/
func GetFilmList(c *gin.Context) {
	var list []FilmModel
	if err := sqlSlave.Table(FilmModel{}.TableName()).Find(&list).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

/*
SearchFilm SELECT gorm.film data from mysql wirh WHERE
curl --location --request GET 'localhost/film/list'
*/
func SearchFilm(c *gin.Context) {
	name := c.Query("name")
	category := c.Query("category")

	var list []FilmModel
	if err := sqlSlave.Table(FilmModel{}.TableName()).Where("category=? AND name=?", category, name).Find(&list).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

/*
CreateFilm insert film info to gorm.film
curl --location --request POST '127.0.0.1/film/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "決戰異次元",
    "category": "科幻"
}'
*/
func CreateFilm(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}

	film := FilmModel{
		Name:     req.Name,
		Category: req.Category,
	}

	if err := sqlMaster.Table(FilmModel{}.TableName()).Create(&film).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, film)
}

/*
UpdateFilmLength update gorm.film
curl --location --request PUT 'localhost/film/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "length_minute": 97
}'
*/
func UpdateFilmLength(c *gin.Context) {
	var req struct {
		ID        int64 `json:"id"`
		LengthMin int64 `json:"length_minute"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}

	if err := sqlMaster.Table(FilmModel{}.TableName()).Update("length", req.LengthMin).Where("id=?", req.ID).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := struct {
		Satus string `json:"status"`
	}{"ok"}

	c.JSON(http.StatusOK, resp)
}
