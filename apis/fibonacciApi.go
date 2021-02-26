package apis

import (
	"fibunacci/daos"
	"fibunacci/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetFibonacciSeries godoc
// @Summary Return a Fibonacci series upto the given count
// @Produce json
// @Param count path integer true "count"
// @Success 200 {object} string
// @Router /fseries/{count} [get]
func GetFibonacciSeries(c *gin.Context) {
	start := time.Now()
	log.Println("Processing GetFibonacciSeries request")

	s := services.NewFibonacciService(daos.NewFibonacciDao())
	count, _ := strconv.ParseInt(c.Param("count"), 10, 32)

	if count <= 0 {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		fseries := s.Start(int(count))
		res := fmt.Sprintf(" %s : GetFibonacciSeries upto (%d) processed in %s", fseries, count, time.Since(start))
		c.JSON(http.StatusOK, res)
	}

	duration := time.Since(start)
	log.Println("Processing GetFibonacciSeries request finished in : ", duration)
}

// GetFibonacciSeries godoc
// @Summary Return Zero Fibonacci Series
// @Produce json
// @Success 200 {object} int
// @Router /fzero/ [get]
func GetAllFibonacciSeries(c *gin.Context) {
	start := time.Now()
	log.Println("Processing GetFibonacciSeries request")

	s := services.NewFibonacciService(daos.NewFibonacciDao())

	fcount := s.Count()
	res := fmt.Sprintf(" %d : GetAllFibonacciSeries processed in %s", fcount, time.Since(start))
	c.JSON(http.StatusOK, res)
	duration := time.Since(start)

	log.Println("Processing GetAllFibonacciSeries request finished in : ", duration)
}

// GetFibonacciSeries godoc
// @Summary Return the Fibonacci number given an ordinal
// @Produce json
// @Param ordinal path integer true "ordinal"
// @Success 200 {object} string
// @Router /fnumber/{ordinal} [get]
func GetFibonacciNumberForOrdinal(c *gin.Context) {
	start := time.Now()
	log.Println("Processing GetFibonacciNumberForOrdinal request")

	ordinal, _ := strconv.ParseInt(c.Param("ordinal"), 10, 32)
	if ordinal < 0 {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		fibNumber := services.FetchFibNumber(int(ordinal))
		res := fmt.Sprintf(" %s : GetFibonacciNumberForOrdinal(%d) processed in %s", fibNumber, ordinal, time.Since(start))
		c.JSON(http.StatusOK, res)
	}

	duration := time.Since(start)
	log.Println("Processing GetFibonacciNumberForOrdinal request finished in : ", duration)
}
