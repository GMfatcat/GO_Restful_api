package reader

import (
	"encoding/json"
	"fmt"
	"go_Restful_api/action/common"
	"go_Restful_api/action/connector"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func JsonReader(filePath string) (common.Data, error) {
	// 讀取 JSON 檔案
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return common.Data{}, fmt.Errorf("Error reading JSON file: %v", err)
	}

	// 解析 JSON 檔案
	var data common.Data
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return common.Data{}, fmt.Errorf("Error decoding JSON file: %v", err)
	}

	return data, nil

}

func ReadAllFiles(directoryPath string) (int, error) {

	// Count valid json files
	var jsonFileCount int = 0

	// Connect to PostgreSQL database
	fmt.Println("Connect to Database")
	db, err := connector.ConnectDB()
	if err != nil {
		return jsonFileCount, fmt.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Get all files under directory
	fmt.Println("Read All Files")
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return jsonFileCount, fmt.Errorf("Error reading directory: %v", err)
	}

	// Read json one by one
	for _, file := range files {

		filePath := filepath.Join(directoryPath, file.Name())

		// 檢查是否為 JSON 檔案
		if filepath.Ext(filePath) != ".json" {
			continue // 如果不是 JSON 檔案，就跳過
		}

		// Read Data
		data, err := JsonReader(filePath)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		// Save data to Postgres Table
		convertedFileName, err := convertFileName(filePath)
		if err != nil {
			fmt.Println("Convert Name Error:", err)
		}
		// 檢查並創建表格
		err = connector.CheckDBTable(db, convertedFileName)
		if err != nil {
			fmt.Println("CheckTable Error:", err)
		}
		// Insert data into table
		err = connector.InsertData(db, data, convertedFileName)
		if err != nil {
			fmt.Println("Insert Error:", err)
		}

		// Show data
		fmt.Printf("File: %s\n", filePath)
		fmt.Println("lons:", data.Lons)
		fmt.Println("lons:", data.Lats)
		fmt.Println("Counts:", data.Counts)

		// Record Processed Files
		jsonFileCount++

	}

	return jsonFileCount, nil
}

func convertFileName(filePath string) (string, error) {
	// 提取檔名（去除路徑）
	baseName := filepath.Base(filePath)
	// 去除擴展名 .json
	baseNameWithoutExt := strings.TrimSuffix(baseName, ".json")
	// 將 - 替換為 _
	replacedUnderscore := strings.ReplaceAll(baseNameWithoutExt, "-", "_")
	// 在前面加上 data_
	result := "data_" + replacedUnderscore

	return result, nil
}
