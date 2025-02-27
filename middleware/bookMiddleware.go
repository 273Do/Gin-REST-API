package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ユーザーエージェントとリクエストの処理にかかった時間を記録するミドルウェア関数
func RecordUaAndTime(c *gin.Context) {
	// ロガーを初期化
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	// 現在の時刻を取得
	oldTime := time.Now()

	// リクエストの処理を続ける
	ua := c.GetHeader("User-Agent")
	c.Next()
	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("Ua", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
