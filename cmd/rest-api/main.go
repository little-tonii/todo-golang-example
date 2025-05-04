package main

import "todo-golang-example/internal/infrastructure/config"

func main() {
	errors := config.LoadEnvironment()
	if errors != nil && len(errors) > 0 {
		panic(errors)
	}
	error := config.InitializeDatabase()
	if error != nil {
		panic(error)
	}
	defer config.Database.Close()
}
