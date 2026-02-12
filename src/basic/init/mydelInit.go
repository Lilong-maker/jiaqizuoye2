package init

import (
	"fmt"
	"jiaqizuoye2/src/basic/config"
	"jiaqizuoye2/src/handler/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func MysqlInit() {
	MysqlConfig := config.Gen.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlConfig.User,
		MysqlConfig.Password,
		MysqlConfig.Host,
		MysqlConfig.Port,
		MysqlConfig.Database,
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	err = config.DB.AutoMigrate(&model.User{}, &model.Sp{})
	if err != nil {
		return
	}
	fmt.Println("表添加成功")
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := config.DB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
