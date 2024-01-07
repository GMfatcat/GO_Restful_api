package main

import (
	"fmt"
	"go_Restful_api/action/reader"
	"log"
)

func main() {
	fmt.Println("Starting")

	// add data to sql section (do it if no table in sql)
	var addNewData bool = false
	if addNewData {

		// Read all data & save to SQL database
		folderPath := "./data"
		jsonFileCount, err := reader.ReadAllFiles(folderPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Total JSON files processed: %d\n", jsonFileCount)
	}

	// Query Testing
	fmt.Println("Querying...")
	// exist case
	fmt.Println("----------------")
	var user string = "Fatcat"
	var query string = "2022-01-01"
	queryErr := reader.ClientQuery(user, query)
	if queryErr != nil {

		log.Fatal(queryErr)

	}
	// // non-exist case
	// fmt.Println("----------------")
	// var user2 string = "Batman"
	// var query2 string = "2022-01-32"
	// query2Err := reader.ClientQuery(user2, query2)
	// if query2Err != nil {

	// 	log.Fatal(query2Err)

	// }
	// v2 case -- > now can return the flag -- querySuccess
	fmt.Println("----------------")
	var user3 string = "Batman3"
	var query3 string = "2022-01-02"
	querySuccess, query3Err := reader.ClientQuery2(user3, query3)
	if query3Err != nil {

		log.Fatal(query3Err)

	}
	fmt.Println("Query status:", querySuccess)

}
