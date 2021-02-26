package daos

import (
	"fibunacci/models"
	"fibunacci/util"
	"reflect"
	"testing"
)

func init() {
	util.LoadEnvFromFileForTests()
}

func TestFibonacciDao_Get(t *testing.T) {
	// Create a test fibonacci
	sdao := NewSqlDao("../config/db.yaml")
	testModel := models.Fibonacci{}
	testModel.PreviousNumber = 10
	testModel.NextNumber = 11

	tx, _ := gConn.Begin()
	ctx := util.TxContext{TX: tx}

	sdao.InsertFibonacci(&testModel, ctx.TX)

	type args struct {
		pk  int
		ctx util.TxContext
	}
	tests := struct {
		name string
		args args
		want models.Fibonacci
	}{
		"TestFibonacciDao_Get", args{pk: testModel.Id, ctx: ctx}, testModel,
	}
	t.Run(tests.name, func(t *testing.T) {
		dao := FibonacciDao{}
		if got := dao.Get(tests.args.pk, tests.args.ctx); !reflect.DeepEqual(got, tests.want) {
			t.Errorf("Get() = %v, want %v", got, tests.want)
		}
	})

	tx.Rollback()
}

func TestNewFibonacciDao(t *testing.T) {
	tt := struct {
		name string
		want *FibonacciDao
	}{
		"TestNewFibonacciDao", &FibonacciDao{},
	}

	t.Run(tt.name, func(t *testing.T) {
		if got := NewFibonacciDao(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewFibonacciDao() = %v, want %v", got, tt.want)
		}
	})

}

func TestFibonacciDao_Save(t *testing.T) {
	type args struct {
		fm  models.Fibonacci
		ctx util.TxContext
	}
	tx, _ := gConn.Begin()
	ctx := util.TxContext{TX: tx}

	tt := struct {
		name   string
		args   args
		wantPk int
	}{
		"TestFibonacciDao_Save", args{fm: models.Fibonacci{PreviousNumber: 0, NextNumber: 1}, ctx: ctx}, -1,
	}

	t.Run(tt.name, func(t *testing.T) {
		dao := FibonacciDao{}
		if gotPk := dao.Save(tt.args.fm, tt.args.ctx); gotPk < tt.wantPk {
			t.Errorf("Save() = %v, want %v", gotPk, tt.wantPk)
		}
	})

	tx.Rollback()
}

func TestFibonacciDao_Begin(t *testing.T) {
	t.Run("TestFibonacciDao_Begin", func(t *testing.T) {
		dao := &FibonacciDao{}
		if gotCtx := dao.Begin(); gotCtx.TX == nil {
			t.Errorf("Begin() TX failed with nil tx handle")
		}
	})

}
