package main

import "github.com/btcthirst/angular-go-psql/internal/api"

func main() {
	api.Initialize()
	api.Start()
}
