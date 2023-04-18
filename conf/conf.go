package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Username string `toml:"username"`
}

func ReadConf() {
	var conf Config

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println("读取配置文件失败：", err)
		return
	}

	// 输出配置信息
	fmt.Println(conf.Username)
}
