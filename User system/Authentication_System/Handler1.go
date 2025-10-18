package Authentication_System

import (
	"User_system/Error_Handling"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Login(c *gin.Context) {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	var l LoginForm
	if err := c.ShouldBindJSON(&l); err != nil {
		zap.L().Error("Login fault", zap.String("err Login", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, ch := range l.Username {
		if strings.ContainsRune(LoginCheckU, ch) {
			zap.L().Warn("Invalid character in username", zap.String("char", string(ch)))
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Name", string(ch))})
			return
		}
	}

	for _, ch1 := range l.Password {
		if strings.ContainsRune(LoginCheckP, ch1) {
			zap.L().Warn("Invalid character in password", zap.String("char", string(ch1)))
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Password", string(ch1))})
			return
		}
	}

	valid := false
	for _, ch2 := range LevelCheck {
		if ch2 != l.Level {
			valid = true
			break
		}
	}
	if !valid {
		zap.L().Warn("Invalid level", zap.String("level", l.Level))
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Level: %s", l.Level)})
		return
	}
	zap.L().Info("Login ok\"", zap.String("User", l.Username))
	zap.L().Info("Login ok\"", zap.String("Level", l.Level))

	c.JSON(200, gin.H{"Login ok": fmt.Sprintf("User %s", l.Level)})

}
