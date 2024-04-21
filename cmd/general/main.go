package main

import "github.com/asaphin/all-databases-go/internal/infrstructure/postgres"

func main() {
	_ = postgres.NewDB("postgres")
}
