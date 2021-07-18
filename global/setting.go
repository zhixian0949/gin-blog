package global

import (
	"github.com/zhixian0949/gin-blog/pkg/logger"
	"github.com/zhixian0949/gin-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
