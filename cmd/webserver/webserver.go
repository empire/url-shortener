package main

import (
	"fmt"
	"net/http"

	"github.com/empire/url-shortener/internal/grpc/shortener"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

const addr = "localhost:9000"

type Shorten struct {
	Hash string `json:"hash"`
	Url  string `json:"url" validate:"min=1,max=120"`
	Age  int32  `json:"age" validate:"min=1,max=30"`
}

func shortenHandler(ctx *gin.Context) {
	var shorten Shorten
	if err := ctx.ShouldBindJSON(&shorten); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := validator.Validate(shorten); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hash, err := shortener.Shorten(ctx.Request.Context(), shorten.Url, shorten.Age)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		ctx.JSON(400, gin.H{"code": "INVAILD_REQUEST", "message": "Invalid request"})
		return
	}
	// TODO the addr must be got from gin
	ctx.JSON(201, gin.H{"redirect": fmt.Sprintf("%s/%s", addr, hash)})
}

type Retrive struct {
	Code string `uri:"code" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/shorten", shortenHandler)
	r.GET("/:code", func(c *gin.Context) {
		var req Retrive
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		url, err := shortener.GetUrl(c.Request.Context(), req.Code)
		if err != nil {
			c.JSON(404, gin.H{"msg": err})
			return
		}
		c.JSON(307, gin.H{"link": url})
	})
	r.Run(addr)
}
