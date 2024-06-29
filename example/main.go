package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
	"github.com/zhwei820/gconv"
)

func main() {
	ob := orderbook.NewOrderBook()
	fmt.Println(ob)

	TestLimitPlace(ob)
}

func TestLimitPlace(ob *orderbook.OrderBook) {
	quantity := decimal.New(2, 0)
	N := 20
	wg := sync.WaitGroup{}
	wg.Add(N * 2)
	t := time.Now()
	for k := 0; k < N; k += 1 {
		go func() {
			defer wg.Done()
			for i := 1; i <= 100; i += 1 {
				done, partial, partialQty, err := ob.ProcessLimitOrder(orderbook.Buy, fmt.Sprintf("buy-%d-%d", i, k), quantity, decimal.New(int64(i), 0))
				_ = done
				// if len(done) != 0 {
				// 	panic("OrderBook failed to process limit order (done is not empty)")
				// }
				if partial != nil {
					panic("OrderBook failed to process limit order (partial is not empty)")
				}
				if partialQty.Sign() != 0 {
					panic("OrderBook failed to process limit order (partialQty is not zero)")
				}
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	for k := 0; k < N; k += 1 {
		go func() {
			defer wg.Done()
			for i := 50; i < 150; i += 1 {
				done, partial, partialQty, err := ob.ProcessLimitOrder(orderbook.Sell, fmt.Sprintf("sell-%d-%d", i, k), quantity, decimal.New(int64(i), 0))
				_ = done
				// fmt.Println("done", done)
				// if len(done) != 0 {
				// 	panic("OrderBook failed to process limit order (done is not empty)")
				// }
				if partial != nil {
					panic("OrderBook failed to process limit order (partial is not empty)")
				}
				if partialQty.Sign() != 0 {
					panic("OrderBook failed to process limit order (partialQty is not zero)")
				}
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()

	gap := time.Now().Sub(t) / time.Millisecond

	fmt.Println("\n===>>")
	fmt.Println("NN: ", N*2*100)
	fmt.Println("gap: ", gap, " ms")

	fmt.Println("ob", gconv.Export(ob))

	if ob.Order("fake") != nil {
		panic("can get fake order")
	}

	if ob.Order("sell-100") == nil {
		panic("can't get real order")
	}

	// fmt.Println("ob.Depth()", gconv.Export(ob.Depth()))
	return
}
