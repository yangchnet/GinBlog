package global

import (
	"GinBlog/pkg/logger"
	"GinBlog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)
