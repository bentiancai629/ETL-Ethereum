package tools

import (
	"ETL-Ethereum/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ETL-Ethereum/conf"
	"ETL-Ethereum/models"
)

func Deploy(cfg *conf.Config) {
	dbCfg := cfg.DBConfig
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})

	if err != nil {
		panic(err)
	}
	err = db.Debug().AutoMigrate(
		&models.Chain{},
		&models.Erc20TransferEvent{},
		&models.ChainInfo{},
	)
	if err != nil {
		panic(err)
	}

	dao := dao.NewLandDao(cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}

	chain := &models.Chain{BackwardBlockNumber: cfg.ChainConfig.BackwardBlockNumber}

	dao.SaveChain(chain)
}
