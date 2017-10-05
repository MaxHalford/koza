# Command-line interface (CLI)

## Development

```sh
cd xgp/cmd
go run main.go fit ../examples/boston/train.csv -m mse
go run main.go predict ../examples/boston/test.csv
```

```sh
cd xgp/cmd
go run main.go fit ../examples/iris/train.csv -m accuracy -y target
go run main.go predict ../examples/iris/test.csv -m accuracy -y target
```