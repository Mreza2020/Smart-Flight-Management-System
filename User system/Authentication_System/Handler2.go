package Authentication_System

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Sign(c *gin.Context) {
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

}
