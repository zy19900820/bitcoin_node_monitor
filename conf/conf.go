package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
}

var Conf Config

func ReadConf() error {
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		fmt.Println("读取配置文件失败：", err)
		return err
	}

	return nil
}
