package initials

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

func InitCfg() {
	var err error
	AppCfg = new(AppConfig)
	fmt.Printf(ConfPath)
	cfg, err := ini.Load(ConfPath)
	if err != nil {
		fmt.Printf("loading file err")
		os.Exit(1)
	}
	err = cfg.MapTo(AppCfg)
	if err != nil {
		fmt.Printf("map struct err")
		os.Exit(1)
	}
}