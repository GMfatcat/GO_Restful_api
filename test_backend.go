/*
v1: pure query no return, used for the backend query only.
v2: return flag, used for the server.
All queries are recorded in log.txt.
*/

package main

import (
	"fmt"
	"go_Restful_api/action/reader"
	"log"
)

// Query Testing
func main() {

	fmt.Println("Querying...")
	// v1 case -- > exist case
	fmt.Println("----------------")
	var user string = "Fatcat"
	var query string = "2022-01-01"
	queryErr := reader.ClientQuery(user, query)
	if queryErr != nil {
		log.Println(queryErr)
	}

	// v1 case -- > non-exist case
	fmt.Println("----------------")
	var user2 string = "Batman"
	var query2 string = "2022-01-32"
	query2Err := reader.ClientQuery(user2, query2)
	if query2Err != nil {
		log.Println(query2Err)
	}

	// v2 case -- > now can return the flag -- querySuccess
	fmt.Println("----------------")
	var user3 string = "Batman3"
	var query3 string = "2022-01-02"
	querySuccess, query3Err := reader.ClientQuery2(user3, query3)
	if query3Err != nil {
		log.Println(query3Err)
	}

	fmt.Println("Query status:", querySuccess)
}
