package config

import (
	"compro/database/seeds"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionMysql() (*Mysql, error) {
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.DBName)

	db, err := gorm.Open(mysql.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[connection mysql -1] Failed to connect database " + cfg.Mysql.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[connection mysql -2] Failed to connect database ")
		return nil, err
	}

	seeds.SeedAdmin(db)

	sqlDB.SetMaxOpenConns(cfg.Mysql.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.Mysql.DBMaxIdle)

	return &Mysql{DB: db}, nil
}
