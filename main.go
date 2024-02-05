package main

import (
	"os"
	"surveillance/server"
)

func main() {
	// Проверка наличия папки
	if _, err := os.Stat("screenshot_list/"); os.IsNotExist(err) {
		if err := os.Mkdir("screenshot_list", 0110); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	// Start the server
	server.Run()

}
