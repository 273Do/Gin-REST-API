package controller

import (
	"gin-rest-api/model"
	"gin-rest-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 書籍を追加する
// *gin.Contextでメモリから値を参照しているので効率的
func BookAdd(c *gin.Context) {

	// Book構造体を初期化
	book := model.Book{}

	// リクエストのBodyをパースしてBook構造体に入れる
	err := c.Bind(&book)

	// エラーがある場合は400を返す
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	// Bookに関するロジックはservice.BookServiceに記述
	// インスタンスを作成
	bookService := service.BookService{}

	// SetBookメソッドを呼び出し，追加する
	// エラーがあればerrに入れる
	err = bookService.SetBook(&book)

	// エラーがある場合は500を返す
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// 201を返す
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// 書籍一覧を取得する
func BookList(c *gin.Context) {
	// BookService構造体のインスタンスを作成
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()

	// gin.Hでjson形式のレスポンスを完結に記述
	// jsでいうとjson({})に該当
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}

// 書籍情報を更新する
func BookUpdate(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	// UpdateBookメソッドを呼び出し，更新する
	// エラーがあればerrに入れる
	err = bookService.UpdateBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// 書籍を削除する
func BookDelete(c *gin.Context) {
	// リクエストからidを取得(文字列)
	id := c.PostForm("id")

	// 文字列のidをintに変換
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	// DeleteBookメソッドを呼び出し，削除する
	err = bookService.DeleteBook(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
