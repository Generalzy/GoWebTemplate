package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/Generalzy/GeneralSaaS/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitMysql(c *conf.MysqlConf) (*gorm.DB, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(c.Timeout)*time.Second)
	// 避免其他地方忘记 cancel，且重复调用不影响
	defer cancel()

	flagCh := make(chan struct{})
	defer close(flagCh)

	var (
		db  *gorm.DB
		err error
	)

	// 若InitMysql提前关闭则本函数会泄露
	// 对Context欠缺理解此处手动造成协程泄露
	go func(flagCh chan struct{}) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Username, c.Password, c.Host, c.Port, c.DB,
		)

		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         191,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			// 不使用默认的事务
			SkipDefaultTransaction: false,
			// 不建立实际的外键约束,约束靠代码实现
			DisableForeignKeyConstraintWhenMigrating: true,
			// 表名迁移为单数
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			// 禁止打印日志
			Logger: logger.Default.LogMode(logger.Silent),
		})

		flagCh <- struct{}{}
	}(flagCh)

	for {
		select {
		case <-flagCh:
			if err != nil {
				return nil, err
			}
			return db, nil
		case <-ctx.Done():
			return nil, errors.New("init gorm timeout")
		}
	}
}
