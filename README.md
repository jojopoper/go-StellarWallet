# Go stellar wallet for windows and linux

*Description: In the testnet and the livenet, This grogram running all right.*

## Current func

	1. create account
	2. quary account info
	3. payment
	4. merge account
	5. quary account operations

## Command line

	--live 		for https://horizon.stellar.org
	--test 		for https://horizon-testnet.stellar.org/
	*note: default is testnet*

	--proxy 	format is  --proxy "IP;PORT;Username;Password"

	Example:
	 go run wallet.go --live --proxy "127.0.0.1;8181;;"

## Installation


```shell
go get github.com/jojopoper/go-StellarWallet

go build wallet.go
```
