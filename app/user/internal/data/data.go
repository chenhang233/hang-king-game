package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hang-king-game/app/user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	conf  *Config
	db    map[string]*gorm.DB
	redis map[int]*redis.Client
}
type Config struct {
	dataConf *conf.Data
}

func (d *Data) GetDB(name string) *gorm.DB {
	g, ok := d.db[name]
	if !ok {
		panic(fmt.Sprintf("the database does not exist: %s", name))
	}
	return g
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	d := &Data{
		conf:  &Config{dataConf: c},
		db:    make(map[string]*gorm.DB),
		redis: make(map[int]*redis.Client),
	}
	for _, database := range c.Databases {
		if database.Driver == "mysql" || database.Driver == "tidb" {
			dsn := database.Source
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				return d, nil, err
			}
			d.db[database.Name] = db.Debug()
		}
	}
	for _, rd := range c.Redis {
		nc := redis.NewClient(&redis.Options{
			Addr:         rd.Addr,
			ReadTimeout:  rd.ReadTimeout.AsDuration(),
			WriteTimeout: rd.WriteTimeout.AsDuration(),
			DB:           int(rd.DbIndex),
			Password:     rd.Password,
		})
		d.redis[int(rd.DbIndex)] = nc
	}
	for _, client := range d.redis {
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			return d, nil, err
		}
	}

	cleanup := func() {
		for _, db := range d.db {
			sqlDB, _ := db.DB()
			_ = sqlDB.Close()
		}
		for _, client := range d.redis {
			_ = client.Close()
		}
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}
