package services

import (
	"fibunacci/models"
	"fibunacci/util"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// The service interface , typically implemented by DAO's
type FibonacciService interface {
	Save(models.Fibonacci, util.TxContext) (pk int)
	Get(id int, ctx util.TxContext) models.Fibonacci
	Begin() (ctx util.TxContext)
	End(ctx util.TxContext)
	GetCount() (count int)
}

// The type that wraps a implementation of the above service interface
type FibonacciServiceWrapper struct {
	fs FibonacciService
}

// Constructor that takes type that implements the service ie concrete DAO
func NewFibonacciService(fs FibonacciService) *FibonacciServiceWrapper {
	return &FibonacciServiceWrapper{fs}
}

func (f FibonacciServiceWrapper) Count() (count int) {
	return f.fs.GetCount()
}

func (f FibonacciServiceWrapper) Start(count int) (fseries []string) {

	log.Println("Service method start called with count : ", count)

	// Begin a TX, so that we can rollback all the inserts
	ctx := f.fs.Begin()

	// Load the state (next, previous) from the DB, even though the values are available in the model
	var PreviousNumber uint32 = 0
	var NextNumber uint32 = 1
	fseries = make([]string, 0)
	for i := 0; i <= count; i++ {
		log.Println(i, " ", PreviousNumber)
		fseries = append(fseries, strconv.FormatInt(int64(PreviousNumber), 10))
		fibonacci := models.Fibonacci{PreviousNumber: NextNumber, NextNumber: NextNumber + PreviousNumber}
		pk := f.fs.Save(fibonacci, ctx)
		//could have used the model itself, just to show the insert and read back again steps
		fnRefresh := f.fs.Get(pk, ctx)
		PreviousNumber = fnRefresh.PreviousNumber
		NextNumber = fnRefresh.NextNumber
	}

	//Rollback, typically we will commit but in case we are rollback as we do not want any persistant data
	f.fs.End(ctx)

	log.Println("Service method start finished")
	return fseries
}

func FetchFibNumber(ordinal int) (fibNumber string) {
	log.Println("Service method FetchFibNumber called with count : ", fibNumber)

	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	pow := math.Pow(phi, float64(ordinal))
	ret := math.Round(pow / sqrt5)

	log.Println("Finished Service method FetchFibNumber called with count : ", fibNumber)

	return strings.Split(fmt.Sprintf("%f", ret), ".")[0]
}
