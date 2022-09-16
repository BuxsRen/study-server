package models

import (
	"github.com/gohouse/gorose/v2"
	"study-server/app/libs/utils"
)

// Prize 奖品表
type Prize struct{}

func (p *Prize) TableName() string {
	return "prize"
}

func (p *Prize) GetList() []gorose.Data {
	list, e := DB(p).Get()
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	return list
}
