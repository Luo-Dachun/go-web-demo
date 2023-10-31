package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
	"web-demo/config"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server

	GVA_VP *viper.Viper

	GVA_LOG *zap.Logger

	lock sync.RWMutex
)
