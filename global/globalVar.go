package global

import (
	"poetry/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
    Settings config.ServerConfig
	Lg *zap.Logger
	Trans ut.Translator
	DB *gorm.DB
	Redis  *redis.Client
	JWTConfig config.JWTConfig
)
