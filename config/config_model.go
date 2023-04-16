package config

const (
	serviceName = "openai" // 不可修改，會影響資料一致性
)

type RunMode = string

const (
	RunModeDebug      RunMode = "debug"
	RunModeProduction RunMode = "production"
	RunModeLocal      RunMode = "local"
)

type Config struct {
	SystemENV
	MysqlENV
	Logger
}

var version = "v0.0.0"

// SystemENV : 系統環境變數
type SystemENV struct {
	//Project     string `env:"project" validate:"required"`
	//Environment string `env:"environment" validate:"required"`
	Port     int    `env:"port" envDefault:"8080"`
	RunMode  string `env:"runMode" envDefault:"local"`
	Timezone string `env:"timezone" envDefault:"UTC"`
}

type MysqlENV struct {
	DSNAccount string `env:"ACCOUNT_MYSQL_URL" envDefault:"stock_user:secret@tcp(db:3306)/stock?parseTime=true&loc=Local"`
}

type Logger struct {
	FilePath      string `env:"filePath"`
	Level         string `env:"level"`
	MaxSize       int    `env:"maxSize"`
	MaxBackups    int    `env:"maxBackups"`
	MaxAge        int    `env:"maxAge"`
	Compress      bool   `env:"compress"`
	ServiceName   string `env:"serviceName"`
	IsShowConsole bool   `env:"isShowConsole"`
}
