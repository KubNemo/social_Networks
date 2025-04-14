package main

import "socialNetworks/internal/db"

func main() {
	db.InitDB()
	db.ApplyMigrations()
}
