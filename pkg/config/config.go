package config

import (
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突

)

// Viper viper 实例
var Viper *viperlib.Viper
