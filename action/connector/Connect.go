package connector

import (
	"database/sql"
	"fmt"
	"go_Restful_api/action/common"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	connStr := "user=postgres dbname=mydb password=123456 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkDBError(err)
	return db, nil

}

func checkDBError(err error) {
	if err != nil {
		panic(err)
	}
}

// Check if table exist, if no create table
func CheckDBTable(db *sql.DB, tableName string) error {
	// Check if table exists
	exists, err := tableExists(db, tableName)
	if err != nil {
		return err
	}
	// Create table if it doesn't exist
	if !exists {
		err := createTable(db, tableName)
		if err != nil {
			return err
		}
		log.Printf("Table %s created successfully", tableName)
	} else {
		log.Printf("Table %s already exists", tableName)
	}
	return nil
}

func tableExists(db *sql.DB, tableName string) (bool, error) {
	var result bool
	query := "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)"
	err := db.QueryRow(query, tableName).Scan(&result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func createTable(db *sql.DB, tableName string) error {
	query := fmt.Sprintf("CREATE TABLE %s (lons FLOAT, lats FLOAT, counts INTEGER)", tableName)
	_, err := db.Exec(query)
	return err

}

func InsertData(db *sql.DB, data common.Data, tableName string) error {
	query := fmt.Sprintf("INSERT INTO %s(lons, lats, counts) VALUES ($1, $2, $3)", tableName)
	// 將數據插入到表格中
	var rowsAdded int = 0
	for i := 0; i < len(data.Lons); i++ {
		_, err := db.Exec(query, data.Lons[i], data.Lats[i], data.Counts[i])
		if err != nil {
			return err
		}
		rowsAdded++
	}

	fmt.Printf("Inserted %d rows\n", rowsAdded)

	return nil
}

// Handle http code status, return http code & error //
/*
200 : ok
404 : not found
500 : internal server error
*/
func QueryData(db *sql.DB, queryTableName string) (int, error) {

	// Check if table exists
	exists, err := tableExists(db, queryTableName)
	if err != nil {
		return 500, err
	}

	if !exists {
		return 404, fmt.Errorf("TABLE : %s Not found", queryTableName)
	} else {
		// List all data of table
		err := listAllTableData(db, queryTableName)
		if err != nil {
			return 500, err
		}
	}

	return 200, nil

}

func listAllTableData(db *sql.DB, queryTableName string) error {

	rows, err := db.Query("SELECT * FROM " + queryTableName)
	if err != nil {
		return err
	}
	defer rows.Close()

	var lons, lats float64
	var counts int
	var rowCounts int = 0

	// List all data
	fmt.Printf("TABLE : %s\nrow | lons | lats | counts|\n", queryTableName)
	for rows.Next() {

		err := rows.Scan(&lons, &lats, &counts)
		if err != nil {
			return err
		}
		fmt.Printf("%d |%.4f| |%.4f| |%d|\n", rowCounts, lons, lats, counts)
		rowCounts++
	}
	fmt.Printf("\n(Total %d rows)\n", rowCounts)

	if err := rows.Err(); err != nil {
		return err
	}

	return nil

}
