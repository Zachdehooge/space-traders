module github.com/Zachdehooge/space-traders

go 1.24.0

require (
	github.com/Zachdehooge/space_trader_go v0.0.0-20250531013831-f3264188a9d3
	github.com/fatih/color v1.18.0
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/HOWZ1T/space_trader v0.1.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	golang.org/x/sys v0.32.0 // indirect
)

replace github.com/Zachdehooge/space_trader_go => ./api-wrapper-go
