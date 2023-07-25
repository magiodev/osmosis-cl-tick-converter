package main

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/osmomath"
	clmath "github.com/osmosis-labs/osmosis/v16/x/concentrated-liquidity/math"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <price>")
		os.Exit(1)
	}

	// TODO START - this is loosing precision below 1.0, fix
	//price, err := strconv.ParseFloat(os.Args[1], 64)
	//if err != nil {
	//	fmt.Println("Invalid price argument. Price must be a float.")
	//	os.Exit(1)
	//}
	//priceDec := sdk.NewDec(int64(price))
	// TODO END

	// START Replacement
	priceStr := os.Args[1]
	priceDec, err := sdk.NewDecFromStr(priceStr)
	if err != nil {
		fmt.Println("Invalid price argument. Price must be a float.")
		os.Exit(1)
	}
	// END Replacement

	tickIndex, err := clmath.CalculatePriceToTickDec(priceDec)

	if err != nil {
		fmt.Printf("Error calculating tick index: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Tick index for price %s is %s\n", priceStr, tickIndex.String())
}

func main_sqrt() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <price>")
		os.Exit(1)
	}

	// TODO START - this is loosing precision below 1.0, fix
	//price, err := strconv.ParseFloat(os.Args[1], 64)
	//if err != nil {
	//	fmt.Println("Invalid price argument. Price must be a float.")
	//	os.Exit(1)
	//}
	//sqrtPrice := math.Sqrt(price)
	//sqrtPriceBigDec := osmomath.NewBigDec(int64(sqrtPrice))
	// TODO END

	// START Replacement
	priceStr := os.Args[1]
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Invalid price argument. Price must be a float.")
		os.Exit(1)
	}
	sqrtPrice := math.Sqrt(price)
	sqrtPriceStr := strconv.FormatFloat(sqrtPrice, 'f', -1, 64)
	priceDec, err := osmomath.NewDecFromStr(sqrtPriceStr)
	if err != nil {
		fmt.Println("Invalid price argument. Price must be a float.")
		os.Exit(1)
	}
	// END Replacement

	tickIndex, err := clmath.CalculateSqrtPriceToTick(priceDec)
	if err != nil {
		fmt.Printf("Error calculating tick index: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Tick index for sqrt price %.2f is %d\n", sqrtPrice, tickIndex)
}
