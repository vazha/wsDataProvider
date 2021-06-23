package wsDataProvider

import (
	"fmt"
	"github.com/vazha/wsDataProvider/pkg/services"
	"github.com/vazha/wsDataProvider/pkg/types"
)

func Run() error{
	bookStore := types.NewBookStore()

	bookStore.Set("Kraken", &types.Book{
		Pair: "BTC/USDT",
		LastUpdated: 12345678,
	})

	bookStore.Set("Kraken",  &types.Book{
		Pair: "LTC/USDT",
		LastUpdated: 12345678,
	})

	book, err := bookStore.Get("Kraken", "BTC/USDT")
	if err != nil {
		fmt.Println("Get fail:", err)
	}
	fmt.Printf("book: %+v\n", book)

	services.Engine()

	return nil
}

