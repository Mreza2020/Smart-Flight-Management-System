package Error_Handling

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ConfigST struct {
	Log        ConfigLog      `json:"log"`
	Redis      ConfigRedis    `json:"redis"`
	DBMysql    ConfigDB       `json:"DBSql"`
	Email      ConfigEmail    `json:"email"`
	DBPostgres ConfigPostgres `json:"DBPostgres"`
}
type ConfigLog struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxSize"`
	MaxBackups int    `json:"maxBackups"`
	MaxAge     int    `json:"maxAge"`
	Compress   bool   `json:"compress"`
}

type ConfigRedis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ConfigDB struct {
	DB string `json:"DB"`
}

type ConfigEmail struct {
	Email           string `json:"Email"`
	EmailPassword   string `json:"EmailPassword"`
	EmailSmtpServer string `json:"EmailSmtpServer"`
	EmailPort       int    `json:"EmailPort"`
}

type ConfigPostgres struct {
	Postgres string `json:"Postgres"`
}

var (
	address string
	data    ConfigST
)

func LoadConfigFile() ConfigST {
	InitLogger(zapcore.InfoLevel)
	address = "config.json"

	fileData, err := os.ReadFile(address)
	if err != nil {
		zap.L().Error("Err OpenFile", zap.String("err", err.Error()))
		return ConfigST{}
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		zap.L().Error("Err Unmarshal JSON", zap.String("err", err.Error()))
		return ConfigST{}
	}

	return data
}
