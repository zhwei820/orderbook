package example

import (
	"fmt"

	ob "github.com/muzykantov/orderbook"
)

func main() {
	orderBook := ob.NewOrderBook()
	fmt.Println(orderBook)
}
