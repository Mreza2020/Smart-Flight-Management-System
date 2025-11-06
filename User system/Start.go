package main

import (
	"User_system/Authentication_System"
	"User_system/Error_Handling"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	zap.L().Info("Starting server User system ...")
	Api := gin.Default()
	Api.POST("/Login", Authentication_System.Login)
	Api.POST("/Sign", Authentication_System.Sign)
	Api.POST("/Sign1", Authentication_System.Sign1)
	Api.POST("/Sign2", Authentication_System.Sign2)

	if err := Api.Run(":8081"); err != nil {
		zap.L().Error("Failed to start server", zap.Error(err))
		return
	}
}
