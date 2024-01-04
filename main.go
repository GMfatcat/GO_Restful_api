package main

import (
	"fmt"
	"go_Restful_api/action/creater"
	"go_Restful_api/action/reader"
	"log"
)

func main() {
	fmt.Println("Starting")
	createrValue := creater.Creater
	fmt.Println("Creater:", createrValue)

	// Read all data & save to SQL database
	folderPath := "./data"
	jsonFileCount, err := reader.ReadAllFiles(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total JSON files processed: %d\n", jsonFileCount)

}
