package init

import (
	"context"
	"fmt"
	"jiaqizuoye2/src/basic/config"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client

func RedisInit() {
	RedisConfig := config.Gen.Redis
	Addr := fmt.Sprintf("%s:%d",
		RedisConfig.Host,
		RedisConfig.Port,
	)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: RedisConfig.Password, // no password set
		DB:       RedisConfig.Database, // use default DB
	})
	fmt.Println("redis连接成功")

}
