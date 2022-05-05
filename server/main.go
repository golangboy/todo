package main

import (
	"encoding/json"
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
	g.NoRoute(func(ctx *gin.Context) {

		if ctx.Request.Method == "GET" && ctx.Request.URL.Path == "/list" {
			var ret Todo
			bytes, _ := os.ReadFile("data.json")
			err := json.Unmarshal(bytes, &ret)
			if err == nil {
				ctx.JSON(http.StatusOK, Response{Err: err, Data: ret})
			} else {
				ctx.JSON(http.StatusOK, Response{Err: err, Data: nil})
			}
			return
		}
		if ctx.Request.Method == "POST" && ctx.Request.URL.Path == "/add" {
			var param Todo
			ctx.BindJSON(&param)
			bytes, _ := json.Marshal(param)
			err := os.WriteFile("data.json", bytes, 0666)
			if err == nil {
				ctx.JSON(http.StatusOK, Response{Err: err, Data: nil})
			} else {
				ctx.JSON(http.StatusOK, Response{Err: err, Data: nil})
			}
			return
		}
		var p string
		if p == "/" {
			p = "index.html"
		} else {
			p = ctx.Request.URL.Path[1:]
		}
		ctx.File("../dist/" + p)
	})
	g.Run(":8089")
}
