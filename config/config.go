package config

import (
	"app/utils/file"
	"errors"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type CommandArgs struct {
	ConfigDir  string `json:"config_dir"`
	ConfigFile string `json:"config_file"`
}

type AppConfig struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	Env  string `json:"env"`
}

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `mapstructure:"DB_NAME"`
}

type LogPath struct {
	Path string `json:"path" mapstructure:"PATH"`
	Name string `json:"name" mapstructure:"NAME"`
}

type LoggerConfig struct {
	MaxSize    int       `mapstructure:"MAX_SIZE"`
	MaxBackups int       `mapstructure:"MAX_BACKUPS"`
	MaxAge     int       `mapstructure:"MAX_AGE"`
	Compress   bool      `mapstructure:"COMPRESS"`
	Business   []LogPath `json:"business"`
}

type Config struct {
	App AppConfig    `json:"app"`
	DB  DbConfig     `json:"db"`
	Log LoggerConfig `json:"log"`
}

func InitConfig() (Config, error) {
	// 加载配置文件
	var args CommandArgs
	pflag.StringVar(&args.ConfigDir, "config_dir", "./config", "配置文件目录")
	pflag.StringVar(&args.ConfigFile, "config_file", "config.yaml", "配置文件")
	pflag.Parse()

	// 配置文件是否存在
	if !file.DirExists(args.ConfigDir) {
		return Config{}, errors.New("configDir not exists")
	}

	configPath := filepath.Join(args.ConfigDir, args.ConfigFile)
	if !file.FileExists(configPath) {
		return Config{}, errors.New("configFile not exists")
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
