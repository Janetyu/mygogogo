package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果制定了配置文件则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigName("yaml")               // 设置配置文件的格式为YAML
	viper.AutomaticEnv()                      // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER")           // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_") // 设置为前缀与配置名称用 _ 进行连接
	// export APISERVER_ADDR=:7777 直接设置环境变量进行变更配置
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	// viper 监控配置文件变更，如有变更则热更新程序
	// 所谓热更新是指：可以不重启 API 进程，使 API 加载最新配置项的值
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
