package data

import (
	"context"
	"fmt"
	"time"

	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/conf"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewArticleRepo, NewCommentRepo, NewTagRepo, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, r *redis.Client) (*Data, func(), error) {

	d := &Data{db: db, redis: r}
	cleanup := func() {
		fmt.Printf("closing the data resources..")
		if _, err := d.db.DB(); err != nil {
			fmt.Printf("err:%+v", err)
		}
		if err := d.redis.Close(); err != nil {
			fmt.Printf("err:%+v", err)
		}
	}
	return d, cleanup, nil
}

// NewRedis .
func NewRedis(c *conf.Data) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:            c.Redis.Addr,
		Password:        c.Redis.Password,
		DB:              int(c.Redis.Db),
		DialTimeout:     c.Redis.DialTimeout.AsDuration(),
		WriteTimeout:    c.Redis.WriteTimeout.AsDuration(), //
		ReadTimeout:     c.Redis.ReadTimeout.AsDuration(),  //
		Network:         "tcp",                             //默认（不需要更改）
		PoolSize:        10,                                // too much pointless
		MinIdleConns:    1,                                 //空闲个数
		MaxIdleConns:    5,                                 //空闲个数
		MaxActiveConns:  10,                                //最大打开个数
		ConnMaxIdleTime: time.Second * 90,                  //超出空闲个数的Conn超过100后无响应则关闭
		ConnMaxLifetime: time.Minute * 5,                   //空闲的Conn超过5分钟后无响应则关闭
	})
	if err := r.Ping(context.Background()).Err(); err != nil {
		panic("failed to connect redis")
	}
	return r
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:MyNewPass4!@tcp(192.168.56.129:3306)/kratos_demo?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect mysql,err:%+v", err))
	}
	if err = InitDB(db); err != nil {
		panic(fmt.Sprintf("failed set db config into: %v", err))
	}
	return db
}

func InitDB(db *gorm.DB) error {
	return db.AutoMigrate(&biz.User{})
}
