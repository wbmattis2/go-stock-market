package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stock struct {
  name string
  price float32
}

func randomNumberGen(min float32, max float32) float32 {	
	return min + (max-min)*rand.Float32()
}

func (stock *Stock) updateStock() {
  change := randomNumberGen(-10000, 10000)
  stock.price += change
}

func displayMarket(market []Stock) {
  for i := range market {
    fmt.Printf("%v - %v\n", market[i].name, market[i].price)
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
	stockMarket := []Stock{{"GOOG", 2313.5}, {"AAPL", 157.28}, {"FB", 203.77}, {"TWTR", 50.00}}
  displayMarket(stockMarket)
  (&(stockMarket[3])).updateStock()
  fmt.Println("***UPDATE***")
  displayMarket(stockMarket)
}