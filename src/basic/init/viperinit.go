package init

import (
	"fmt"
	"jiaqizuoye2/src/basic/config"

	"github.com/spf13/viper"
)

func ViperInit() {
	viper.SetConfigFile("../../../config.yml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config.Gen)
	if err != nil {
		return
	}
	fmt.Println("配置文件加载成功")
}
