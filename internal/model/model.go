package model

import (
	"fmt"
)

// Model 通用模型
type Model struct {
	ID         uint32 `gorm: "primary_key" json:"id"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreateTime uint32 `json:"create_time"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := fmt.Sprintf(format, a)
	db, err := gorm.Open(databaseSetting.DBType, s, 
		databaseSetting.name,
		databaseSetting.pw,
		databaseSetting.host,
		databaseSetting.dbname
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
