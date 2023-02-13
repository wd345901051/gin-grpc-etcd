package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 数据库的连接
var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	fmt.Println(host, port, database)
	err := Database(dsn)
	if err != nil {
		panic(err)
	}
}

// gorm的定义
func Database(dsn string) error {
	var ormLogger logger.Interface // 在终端输出原生的sql语句
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // 默认string类型的长度
		DisableDatetimePrecision:  true,  // 禁用datatime的精度，mysql 5.6 之前的数据是不支持的
		DontSupportRenameIndex:    true,  // 重命名索引的时候采用删除并新建的方式.因为mysql 5.7 之前的数据库是不支持重命名的
		DontSupportRenameColumn:   true,  // 用 change 重命名列， mysql 8 之前的数据库是不支持重命名列的
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 复数形式 User的表名应该是users
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置空闲连接池
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
	return nil
}
