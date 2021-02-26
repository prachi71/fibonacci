package daos

import (
	"fibunacci/models"
	"fibunacci/util"
	"reflect"
	"testing"
)

func init() {
	initialize("../config/db.yaml")
	util.LoadEnvFromFileForTests()
}

func TestNewSqlDao(t *testing.T) {
	type args struct {
		fn string
	}
	tt := struct {
		name string
		args args
		want *SqlDao
	}{
		"TestSqlDao_GetByPk", args{"../config./db.yaml"}, &SqlDao{},
	}

	t.Run(tt.name, func(t *testing.T) {
		if got := NewSqlDao(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewSqlDao() = %v, want %v", got, tt.want)
		}
	})

}

func TestSqlDao_GetByPk(t *testing.T) {
	// Create a test fibonacci
	sdao, testModel := getMockModel()

	tx, _ := gConn.Begin()

	sdao.InsertFibonacci(&testModel, tx)

	type args struct {
		pk int
	}
	tt := struct {
		name string
		args args
		want models.Fibonacci
	}{
		"TestFibonacciDao_Get", args{testModel.Id}, testModel,
	}

	t.Run(tt.name, func(t *testing.T) {
		sqlDao := &SqlDao{}
		if got := sqlDao.GetByPk(tt.args.pk, tx); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetByPk() = %v, want %v", got, tt.want)
		}
	})

	tx.Rollback()

}

func getMockModel() (*SqlDao, models.Fibonacci) {
	sdao := NewSqlDao("../config/db.yaml")
	testModel := models.Fibonacci{}
	testModel.PreviousNumber = 10
	testModel.NextNumber = 11
	return sdao, testModel
}

func TestSqlDao_InsertFibonacci(t *testing.T) {
	// Create a test fibonacci
	sdao, testModel := getMockModel()

	tx, _ := gConn.Begin()

	sdao.InsertFibonacci(&testModel, tx)
	type args struct {
		fibonacci *models.Fibonacci
	}
	tt := struct {
		name string
		args args
	}{
		"TestSqlDao_InsertFibonacci", args{&testModel},
	}
	t.Run(tt.name, func(t *testing.T) {
		sqlDao := &SqlDao{}
		if got := sqlDao.GetByPk(tt.args.fibonacci.Id, tx); !reflect.DeepEqual(got, *tt.args.fibonacci) {
			t.Errorf("GetByPk() = %v, want %v", got, tt.args.fibonacci)
		}
	})

	tx.Rollback()
}

func Test_createTable(t *testing.T) {

	tt := struct {
		name string
	}{
		"Test_createTable",
	}

	t.Run(tt.name, func(t *testing.T) {
		db := getConnection()
		if _, err := db.Exec("select * from fibonacci"); err != nil {
			t.Error("createTable() , err")
		}
	})

}

func Test_getConnection(t *testing.T) {
	t.Run("TestConnection", func(t *testing.T) {
		conn := getConnection()
		if err := conn.Ping(); err != nil {
			t.Error("getConnection() , err")
		}
	})

}

func Test_initialize(t *testing.T) {
	if crudConfig.Ddl.Create == "" {
		t.Error("Initialization failure, db.yaml config not loaded properly")
	}
}

func Test_initializeDb(t *testing.T) {

}
