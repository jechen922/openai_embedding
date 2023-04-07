package config

const (
	RunModeDebug      = "debug"
	RunModeProduction = "production"
	RunModeLocal      = "local"
)
const (
	LogLevelInfo  = "info"
	LogLevelError = "error"
)

type Env struct {
	SystemENV
	MysqlENV
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
