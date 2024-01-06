package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NumberResponse struct {
	Number int `json:"number"`
}

func handleGetNumber(w http.ResponseWriter, r *http.Request) {
	// 在這裡可以生成一個數字，這裡使用了固定的數字 42 作為範例
	numberResponse := NumberResponse{Number: 42}

	// 將數字轉換為 JSON 格式並返回給前端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numberResponse)
}

func handleGetJSON(w http.ResponseWriter, r *http.Request) {
	// 指定 JSON 文件的路徑
	jsonFilePath := "./data/2022-01-01.json"

	// 讀取 JSON 文件的內容
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	// 將 JSON 內容回傳給前端
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	// 設定伺服器路由
	http.HandleFunc("/getNumber", handleGetNumber)
	http.HandleFunc("/getJSON", handleGetJSON)

	// 設定靜態資源伺服器，指向存放HTML文件的資料夾
	http.Handle("/", http.FileServer(http.Dir("static")))

	// 啟動 HTTP 服務器，監聽在 8080 端口
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}