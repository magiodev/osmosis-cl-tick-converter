package main

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clmath "github.com/osmosis-labs/osmosis/v19/x/concentrated-liquidity/math"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	// Price To Tick returns the tick related to a given price
	var cmdPriceToTick = &cobra.Command{
		Use:   "price-to-tick [price]",
		Short: "Convert price to tick index",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			priceStr := args[0]
			priceDec, err := sdk.NewDecFromStr(priceStr)
			if err != nil {
				fmt.Println("Invalid price argument. Price must be a float.")
				return
			}

			tickIndex, err := clmath.CalculatePriceToTickDec(priceDec)
			if err != nil {
				fmt.Printf("Error calculating tick index: %s\n", err.Error())
				return
			}

			fmt.Println(tickIndex.String())
		},
	}

	// Tick To Price returns the price related to a given tick index
	var cmdTickToPrice = &cobra.Command{
		Use:   "tick-to-price [tick]",
		Short: "Convert tick index to price",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tickStr := args[0]
			tickIndex, err := strconv.ParseInt(tickStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid tick index argument. Tick index must be an integer.")
				return
			}

			price, err := clmath.TickToPrice(tickIndex)
			if err != nil {
				fmt.Printf("Error calculating price: %s\n", err.Error())
				return
			}

			fmt.Println(price.String())
		},
	}

	rootCmd.AddCommand(cmdPriceToTick, cmdTickToPrice)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
