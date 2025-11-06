package Authentication_System

import (
	"User_system/Capabilities"
	"User_system/Data_Base"
	"User_system/Error_Handling"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func GenerateOTP(length int) string {
	otp := ""
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		otp += fmt.Sprintf("%d", n.Int64())

	}
	return otp
}

func SecuritySing1(email string) string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	status := Data_Base.RunRedisServer()
	zap.L().Warn("Redis status", zap.String("status", status))
	if status == "ok" {
		otp := GenerateOTP(10)
		note := Capabilities.SendEmailOtp("Email_sign_in.html", otp, "Smart Flight Management System", email, "احراز هویت", "کد احراز هویت پیشرفته شما {%s}")
		if note == "sent" {
			note1 := Capabilities.SaveOtpToRedis(email, otp, 5)
			if note1 == "ok" {
				return note1
			} else {
				return note1
			}
		} else {
			return "err"
		}
	} else {
		return "err"
	}

}

func SecuritySing2(email, otp1 string) string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	status := Data_Base.RunRedisServer()
	zap.L().Warn("Redis status", zap.String("status", status))
	if status == "ok" {
		otp := Capabilities.GetOtpFromRedis(email)
		switch otp {
		case "ok":
			return "ok"
		case "not":
			return "not"
		case "Error":
			return "err"
		default:
			return "err"
		}
	} else {
		return "err"
	}

}

func SecuritySing3(username, password, level, code, email string) string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	hash, err := hashPassword(password)
	if err != nil {
		zap.L().Error("GenerateFromPassword error", zap.Error(err))
		return ""
	}

	db, status1 := Data_Base.DBMysql()
	if status1 != nil {
		zap.L().Error("Database connection failed", zap.Error(status1))
		return ""
	}
	defer func(db *sql.DB) {
		err1 := db.Close()
		if err1 != nil {
			zap.L().Error("Close database failed", zap.Error(err1))
		}
	}(db)
	zap.L().Info("Connected to MySQL successfully!")

	insertQuery := "INSERT INTO sing(Username, Password, Level, Code, Email) VALUES (?,?,?,?,?)"
	_, err = db.Exec(insertQuery, username, hash, level, code, email)

	if err != nil {
		zap.L().Error("The table could not be added because", zap.Error(err))
		return ""
	}
	return "ok"

}
