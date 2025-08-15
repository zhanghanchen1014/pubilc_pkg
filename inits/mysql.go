package inits

import (
	"fmt"
	"github.com/zhanghanchen1014/pubilc_pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func MysqlInit() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConf.Mysql.User,
		config.AppConf.Mysql.Password,
		config.AppConf.Mysql.Host,
		config.AppConf.Mysql.Port,
		config.AppConf.Mysql.Database,
	)

	config.Once.Do(func() {
		config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
			return
		}
		log.Println("MySQL init success")

		sqlDB, _ := config.DB.DB()

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

	})

}
