package mysql

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"sync"
)

var once sync.Once

func InitDB(dbPoolName string) (*gorm.DB,error) {

	db := &gorm.DB{}
	var err error

	once.Do(func() {
		db, err = lib.GetGormPool(dbPoolName)
	})

	return db,err

}
