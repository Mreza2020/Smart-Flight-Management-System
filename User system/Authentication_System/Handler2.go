package Authentication_System

import (
	"User_system/Error_Handling"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Username string
	Password string
	Level    string
	Code1    string
	Email    string
)

func Sign(c *gin.Context) {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	var s SignForm
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, ch := range s.Username {
		if strings.ContainsRune(LoginCheckU, ch) {
			zap.L().Warn("Invalid character in username", zap.String("char", string(ch)))
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Name", string(ch))})
			return
		}
	}

	for _, ch1 := range s.Password {
		if strings.ContainsRune(LoginCheckP, ch1) {
			zap.L().Warn("Invalid character in password", zap.String("char", string(ch1)))
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Password", string(ch1))})
			return
		}
	}
	isValid := false
	for _, allowed := range Code {
		if s.Code == allowed {
			isValid = true
			break
		}
	}
	if !isValid {
		zap.L().Warn("Invalid code", zap.String("code", s.Code))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong Code"})
		return
	}

	Username, Password, Level, Code1, Email =
		s.Username, s.Password, s.Level, s.Code, s.Email

	status := SecuritySing1(Email)
	if status == "ok" {
		c.JSON(http.StatusOK, gin.H{"ok": "SecuritySing1"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong Send Email"})
		return
	}
}

func Sign1(c *gin.Context) {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	var s SignForm1
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(s.Otp) != 6 || !regexp.MustCompile(`^[0-9]+$`).MatchString(s.Otp) {
		zap.L().Warn("OTP must be a 6-digit numeric string")
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP must be a 6-digit numeric string"})
		return
	}

	status := SecuritySing2(Email)
	switch status {
	case "not":
		c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
		return
	case "err":
		c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
		return
	default:
		if status == s.Otp {
			c.JSON(http.StatusOK, gin.H{"ok": "SecuritySing2"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
			return
		}
	}

}

func Sign2(c *gin.Context) {
	status := SecuritySing3(Username, Password, Level, Code1, Email)
	if status == "ok" {
		c.JSON(http.StatusOK, gin.H{"ok": "SecuritySing3"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
		return
	}
}
