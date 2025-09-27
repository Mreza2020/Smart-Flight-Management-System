package Capabilities

import (
	"User_system/Data_Base"
	"User_system/Error_Handling"
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SaveOtpToRedis(email string, otp string, Storage int) string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := Data_Base.ClientRedis.Set(ctx, "otp:"+email, otp, time.Duration(Storage)*time.Minute).Err()
	if err != nil {
		zap.L().Error("Err Save Otp To Redis", zap.String("err", err.Error()))
		return "error"
	}
	zap.L().Info("Save Otp To Redis")
	return "ok"
}

func GetOtpFromRedis(email string) string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	otp, err := Data_Base.ClientRedis.Get(ctx, "otp:"+email).Result()
	if errors.Is(err, redis.Nil) {
		zap.L().Warn("OTP not found")
		return "not"
	} else if err != nil {
		zap.L().Error("Error retrieving OTP ", zap.String("err", err.Error()))
		return "Error"
	} else {
		zap.L().Info("The otp code is on the server")
		return otp
	}
}
