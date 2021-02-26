package daos

import (
	"fibunacci/models"
	"fibunacci/util"
)

// Struct represent an instance of the dao
type FibonacciDao struct {
}

// TODO Check for toggle ( sql or orm)
func NewFibonacciDao() *FibonacciDao {
	return &FibonacciDao{}
}

func (dao *FibonacciDao) Save(fm models.Fibonacci, ctx util.TxContext) (pk int) {
	sd := NewSqlDao("config/db.yaml")
	sd.InsertFibonacci(&fm, ctx.TX)
	return fm.Id
}

func (dao *FibonacciDao) Get(pk int, ctx util.TxContext) models.Fibonacci {
	sd := NewSqlDao("config/db.yaml")
	return sd.GetByPk(pk, ctx.TX)
}

func (dao *FibonacciDao) Begin() (ctx util.TxContext) {
	NewSqlDao("config/db.yaml")
	tx, _ := gConn.Begin()
	return util.TxContext{TX: tx}
}

func (dao *FibonacciDao) End(ctx util.TxContext) {
	ctx.TX.Rollback()
}
