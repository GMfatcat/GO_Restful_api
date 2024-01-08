/*
Put all new json files in one folder, and set the folderPath.
Each file will be one table in sql with the same name,
while the dash are transferred to lowercase.

In case of accidently executing this file, addNewData set to false in default.
*/

package main

import (
	"fmt"
	"go_Restful_api/action/reader"
	"log"
)

func main() {
	fmt.Println("Starting...")

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
		return
	}
	fmt.Println("addNewData status:", addNewData)
}
