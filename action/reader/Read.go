package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Data struct {
	Lons   []float64 `json:"lons"`
	Lats   []float64 `json:"lats"`
	Counts []int     `json:"counts"`
}

func JsonReader(filePath string) (Data, error) {
	// 讀取 JSON 檔案
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Data{}, fmt.Errorf("Error reading JSON file: %v", err)
	}

	// 解析 JSON 檔案
	var data Data
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return Data{}, fmt.Errorf("Error decoding JSON file: %v", err)
	}

	return data, nil

}

func ReadAllFiles(directoryPath string) (int, error) {

	// Count valid json files
	var jsonFileCount int = 0

	fmt.Println("Read All Files")
	// Get all files under directory
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

		data, err := JsonReader(filePath)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		// Show data
		fmt.Printf("File: %s\n", filePath)
		fmt.Println("lons:", data.Lons)
		fmt.Println("lons:", data.Lats)
		fmt.Println("Counts:", data.Counts)

		jsonFileCount++

	}

	return jsonFileCount, nil
}
