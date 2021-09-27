package initials

var (
	AppCfg *AppConfig
	ConfPath string
)

type AppConfig struct {
	Mysql struct {
		User     string `ini:"user"`
		Password string `ini:"password"`
		Host     string `ini:"host"`
		Port     string `ini:"port"`
		DBName   string `ini:"dbname"`
		CharSet  string `ini:"charset"`
	}
}
