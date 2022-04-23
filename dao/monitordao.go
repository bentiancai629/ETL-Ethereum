package dao

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"ETL-Ethereum/conf"
	"ETL-Ethereum/models"
)

type DataBase struct {
	dbCfg  conf.DBConfig
	db     *gorm.DB
	backup bool
}

func NewLandDao(dbCfg conf.DBConfig) *DataBase {
	dao := &DataBase{
		dbCfg: dbCfg,
	}
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
	dbCfg.Scheme+"?charset=utf8&parseTime=true"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	dao.db = db
	return dao
}


func (dao *DataBase) SaveChain(chain *models.Chain) error {
	daoChain := new(models.Chain)
	res := dao.db.Where("chain_name = ?", chain.ChainName).First(daoChain)
	if res.Error != nil {
		if res.Error.Error() == "record not found" {
			dao.db.Save(chain)
			return nil
		}
		return res.Error
	}

	daoChain.Height = chain.Height
	daoChain.BackwardBlockNumber = chain.BackwardBlockNumber
	dao.db.Save(daoChain)
	return nil
}

func (dao *DataBase) GetChain(chainName string) (*models.Chain, error) {
	chain := new(models.Chain)
	res := dao.db.Where("chain_name = ?", chainName).First(chain)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record")
	}
	return chain, nil
}

func (dao *DataBase) UpdateChain(chain *models.Chain) error {
	if chain == nil {
		return fmt.Errorf("no value!")
	}
	if dao.backup {
		return nil
	}
	res := dao.db.Updates(chain)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("no update!")
	}
	return nil
}

func (dao *DataBase) SaveEvents(erc20Evts []*models.Erc20TransferEvent) error {
	if erc20Evts != nil && len(erc20Evts) > 0 {
		for _, event := range erc20Evts {
			res := dao.db.Save(event)
			if res.Error != nil {
				logs.Error("addEvent event %v err: %v", event, res.Error)
			}
		}
	}
	return nil
}

func (dao *DataBase) SaveEventsIgnoredByHashIndex(erc20Evts []*models.Erc20TransferEvent) error {
	if erc20Evts != nil && len(erc20Evts) > 0 {
		for _, event := range erc20Evts {
			res:= dao.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Save(event)
			if res.Error != nil {
				logs.Error("addEvent event %v err: %v", event, res.Error)
				return res.Error
			}
		}
	}
	return nil
}



func (dao *DataBase) GetChainInfo() []models.ChainInfo {
	var list []models.ChainInfo
	dao.db.Model(models.ChainInfo{}).Order("id desc").Find(&list)
	return list
}
