## Usage

The tool leverages the Cobra CLI library to provide multiple functionalities. You can either convert a price to a tick
index or convert a tick index to a price using the underlying Osmosis Concentrated Liquidity math module.

## Convert Price to Tick Index

To convert a price to a tick index, run the following command at the root of the project:

```bash
go run main.go price-to-tick <price>
```

Example:

```bash
go run main.go price-to-tick 1.23
```

## Convert Tick Index to Price

To convert a tick index to a price, run the following command at the root of the project:

```bash
go run main.go tick-to-price <tick>
```

For negative tick values, use:

```bash
go run main.go tick-to-price -- <-tick>
````

Example:

```bash
go run main.go tick-to-price 456
go run main.go tick-to-price -- -456
```
