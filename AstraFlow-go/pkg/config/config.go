package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	MQ     MQConfig     `mapstructure:"mq"`
	Flask  FlaskConfig  `mapstructure:"flask"`
}

type ServerConfig struct {
	Port      int    `mapstructure:"port"`
	JwtSecret string `mapstructure:"jwt_secret"`
}

type MySQLConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

type MQConfig struct {
	URL string `mapstructure:"url"`
}

type FlaskConfig struct {
	BaseURL string `mapstructure:"base_url"`
}

var Cfg Config

// InitConfig 读取 config.yml
func InitConfig() {
	v := viper.New()

	// 设置环境变量前缀
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// 设置默认值
	v.SetDefault("flask.base_url", "http://localhost:5000")

	// 设置配置文件
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./pkg/config") // 你的 config.yml 放在这里

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("读取配置文件失败: %v, 使用环境变量", err)
	}

	// 映射环境变量到配置结构
	v.BindEnv("flask.base_url", "FLASK_BASE_URL")
	v.BindEnv("mq.url", "MQ_URL")
	v.BindEnv("mysql.host", "DB_HOST")
	v.BindEnv("mysql.user", "DB_USER")
	v.BindEnv("mysql.password", "DB_PASSWORD")
	v.BindEnv("mysql.database", "DB_NAME")
	v.BindEnv("server.port", "SERVER_PORT")

	if err := v.Unmarshal(&Cfg); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	log.Println("配置文件加载成功")
}
