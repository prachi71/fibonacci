package services

import (
	"fibunacci/models"
	"fibunacci/util"
	"reflect"
	"testing"
)

var id int = 1

var lastInsert models.Fibonacci

type MockService struct {
}

func (mockService *MockService) Begin() (ctx util.TxContext) {
	id = 1
	return util.TxContext{}
}

func (mockService *MockService) End(ctx util.TxContext) {
	id = 1
}

func (mockService *MockService) Save(fibonacci models.Fibonacci, ctx util.TxContext) (pk int) {
	fibonacci.Id = id
	id++
	lastInsert = fibonacci
	return fibonacci.Id
}

func (mockService *MockService) Get(id int, ctx util.TxContext) models.Fibonacci {
	return lastInsert
}

func TestFibonacciServiceWrapper_Start(t *testing.T) {
	type fields struct {
		fs FibonacciService
	}
	type args struct {
		count int
	}

	testFSeries := getFseriesUpto10()

	tt := struct {
		name        string
		fields      fields
		args        args
		wantFseries []string
	}{
		"TestFibonacciServiceWrapper_Start", fields{fs: &MockService{}}, args{count: 10}, testFSeries,
	}

	t.Run(tt.name, func(t *testing.T) {
		f := FibonacciServiceWrapper{
			fs: tt.fields.fs,
		}
		if gotFseries := f.Start(tt.args.count); !reflect.DeepEqual(gotFseries, tt.wantFseries) {
			t.Errorf("Start() = %v, want %v", gotFseries, tt.wantFseries)
		}
	})

}

func getFseriesUpto10() []string {
	testFSeries := make([]string, 0)
	testFSeries = append(testFSeries, "0")
	testFSeries = append(testFSeries, "1")
	testFSeries = append(testFSeries, "1")
	testFSeries = append(testFSeries, "2")
	testFSeries = append(testFSeries, "3")
	testFSeries = append(testFSeries, "5")
	testFSeries = append(testFSeries, "8")
	testFSeries = append(testFSeries, "13")
	testFSeries = append(testFSeries, "21")
	testFSeries = append(testFSeries, "34")
	return testFSeries
}

func TestNewFibonacciService(t *testing.T) {
	type args struct {
		fs FibonacciService
	}
	tt := struct {
		name string
		args args
		want *FibonacciServiceWrapper
	}{
		"TestNewFibonacciService", args{fs: &MockService{}}, &FibonacciServiceWrapper{fs: &MockService{}},
	}
	t.Run(tt.name, func(t *testing.T) {
		if got := NewFibonacciService(tt.args.fs); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewFibonacciService() = %v, want %v", got, tt.want)
		}
	})
}

func TestFetchFibNumber(t *testing.T) {
	type args struct {
		ordinal int
	}
	tests := []struct {
		name          string
		args          args
		wantFibNumber string
	}{
		{
			"TestFetchFibNumber11", args{ordinal: 11}, "89",
		},
		{
			"TestFetchFibNumber12", args{ordinal: 12}, "144",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFibNumber := FetchFibNumber(tt.args.ordinal); gotFibNumber != tt.wantFibNumber {
				t.Errorf("FetchFibNumber() = %v, want %v", gotFibNumber, tt.wantFibNumber)
			}
		})
	}

}
