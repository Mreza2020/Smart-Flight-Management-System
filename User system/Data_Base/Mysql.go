package Data_Base

import (
	"User_system/Error_Handling"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DBMysql() (*sql.DB, string) {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	db, err := sql.Open("mysql", Error_Handling.LoadConfigFile().DBMysql.DB)
	if err != nil {
		zap.L().Error("not open mysql server", zap.String("err", err.Error()))
		return nil, "err Connect"
	}
	if err = db.Ping(); err != nil {
		zap.L().Error("not ping", zap.String("err", err.Error()))
		return nil, "err ping"

	}
	zap.L().Info("Server started successfully")
	return db, ""
}
