/*
This file will be the server file (integration of frontend and backend)

User search result will be send to backend,
If passed regex check, go to database to look for query, return a flag
If flag --> return json file from datafolder(not database)
If no flag --> return empty json file
*/

package main

import "fmt"

func main() {
	fmt.Println("Starting")
}
