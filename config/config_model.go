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

type config struct {
	SystemENV
	MysqlENV
	PostgresENV
	LoggerEnv
}

var version = "v0.0.0"

// SystemENV : 系統環境變數
type SystemENV struct {
	//Project     string `env:"project" validate:"required"`
	//Environment string `env:"environment" validate:"required"`
	Port         int    `env:"port" envDefault:"8080"`
	RunMode      string `env:"runMode" envDefault:"local"`
	Timezone     string `env:"timezone" envDefault:"UTC"`
	ChatGPTToken string `env:"CHAT_GPT_TOKEN" envDefault:""`
}

type MysqlENV struct {
	DSNAccount string `env:"ACCOUNT_MYSQL_URL" envDefault:"stock_user:secret@tcp(db:3306)/stock?parseTime=true&loc=Local"`
}

type PostgresENV struct {
	DSNAccount string `env:"ACCOUNT_MYSQL_URL" envDefault:"stock_user:secret@tcp(db:3306)/stock?parseTime=true&loc=Local"`
}

type LoggerEnv struct {
	FilePath      string `env:"filePath"`
	Level         string `env:"level"`
	MaxSize       int    `env:"maxSize"`
	MaxBackups    int    `env:"maxBackups"`
	MaxAge        int    `env:"maxAge"`
	Compress      bool   `env:"compress"`
	ServiceName   string `env:"serviceName"`
	IsShowConsole bool   `env:"isShowConsole"`
}
