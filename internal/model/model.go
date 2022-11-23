package model

import (
	"errors"
	"fmt"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/pkg/setting"

	"github.com/haierspi/gormTracing"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Predicate string

var (
	// =
	Eq = Predicate("=")
	// <>
	Neq = Predicate("<>")
	// >
	Gt = Predicate(">")
	// >=
	Egt = Predicate(">=")
	//<
	Lt = Predicate("<")
	//<=
	Elt  = Predicate("<=")
	Like = Predicate("LIKE")
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	//对非MYSQL不支持
	if databaseSetting.DBType != "mysql" {
		return nil, errors.New("no support database type")
	}

	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   databaseSetting.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true,                        // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Config.Logger = logger.Default.LogMode(logger.Info)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	_ = db.Use(&gormTracing.OpentracingPlugin{})

	return db, nil
}
