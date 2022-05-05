package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Todo []struct {
	AddTime    string `json:"addTime"`
	FinishTime string `json:"finishTime"`
	Title      string `json:"title"`
}

type Response struct {
	Err  error       `json:"err"`
	Data interface{} `json:"data"`
}

func main() {
	g := gin.Default()
	g.StaticFS("/web", http.Dir("../dist/"))
	g.POST("/add", func(ctx *gin.Context) {
		fmt.Println("1")
		var param Todo
		ctx.BindJSON(&param)
		bytes, _ := json.Marshal(param)
		fmt.Println("2")
		err := os.WriteFile("data.json", bytes, 0666)
		if err == nil {
			ctx.JSON(http.StatusOK, Response{Err: err, Data: nil})
		} else {
			ctx.JSON(http.StatusOK, Response{Err: err, Data: nil})
		}
	})
	g.GET("/list", func(ctx *gin.Context) {
		var ret Todo
		bytes, _ := os.ReadFile("data.json")
		err := json.Unmarshal(bytes, &ret)
		if err == nil {
			ctx.JSON(http.StatusOK, Response{Err: err, Data: ret})
		} else {
			ctx.JSON(http.StatusOK, Response{Err: err, Data: nil})
		}
	})
	g.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/web")
	})
	g.Run(":8089")
}
