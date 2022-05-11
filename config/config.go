package config

type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	Mysqlinfo   MysqlConfig `mapstructure:"mysql"`
	RedisInfo   RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logsAddress"`
	JWTConfig   JWTConfig   `mapstructure:"jwt"`
}

type MysqlConfig struct {
	Host                string `mapstructure:"host"`
	Port                int    `mapstructure:"port"`
	UserName            string `mapstructure:"username"`
	Password            string `mapstructure:"password"`
	Database            string `mapstructure:"database"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	JwtSec string `mapstructure:"key"`
}
