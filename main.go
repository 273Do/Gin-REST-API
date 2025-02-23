// reference:https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a

package main

import (
	"gin-rest-api/controller"
	"gin-rest-api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	// middleware
	router.Use(middleware.RecordUaAndTime)

	// CRUD 書籍
	// ルーティングをグループ化
	bookRouter := router.Group("/book")
	{
		// さらにバージョンでグループ化
		v1 := bookRouter.Group("/v1")
		{
			// book/v1/add
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}

	router.Run()
}
