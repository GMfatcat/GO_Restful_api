/*
This file will be the server file (integration of frontend and backend)

User search result will be send to backend,
If passed regex check, go to database to look for query, return a flag
If flag --> return json file from datafolder(not database)
If no flag --> return empty json file
*/

package main

import (
	"encoding/json"
	"fmt"
	"go_Restful_api/action/common"
	"go_Restful_api/action/reader"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	port            = ":9487"
	dataDir         = "./data/"
	successMsg      = "Input received successfully"
	errorMsg        = "Error processing request"
	getDataError    = "Error getting data"
	httpMethodError = "http method not allowed"
)

func handleGetJSON(querySuccess bool, queryFilename string) ([]byte, error) {
	var emptyJSON = []byte("{}")
	if querySuccess {
		// 指定 JSON 文件的路徑
		jsonFilePath := fmt.Sprintf("%s%s.json", dataDir, queryFilename)

		// 讀取 JSON 文件的內容
		jsonData, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			return emptyJSON, err
		}
		// show json data(only for checking)
		// fmt.Println("JSON Data:")
		// fmt.Println(string(jsonData))
		return jsonData, nil
	}

	return emptyJSON, nil
}

func backendProcess(user, query string) (bool, error) {
	// v2 case -- > now can return the flag -- querySuccess
	fmt.Println("Backend processing...")
	querySuccess, queryErr := reader.ClientQuery2(user, query)
	return querySuccess, queryErr
}

// 處理接收前端傳來的字串
func handleReceiveInput(w http.ResponseWriter, r *http.Request) {
	// 限制http request為POST
	if r.Method != http.MethodPost {
		http.Error(w, httpMethodError, http.StatusMethodNotAllowed)
		log.Println(httpMethodError)
		return
	}
	// 紀錄使用者的ip
	userIP := r.RemoteAddr
	// 解碼 JSON 資料
	var inputData common.InputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		http.Error(w, errorMsg, http.StatusBadRequest)
		log.Println("Error decoding JSON:", err)
		return
	}
	// 印出接收到的字串和IP，並回傳回應給前端
	fmt.Printf("Received input: %s, User IP: %s\n", inputData.InputText, userIP)

	// 後端執行查詢
	querySuccess, queryErr := backendProcess(userIP, inputData.InputText)
	fmt.Println("Query status:", querySuccess)
	if queryErr != nil {
		log.Println("Query error:", queryErr)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}
	// get json data(real data or empty json)
	jsonData, getDataErr := handleGetJSON(querySuccess, inputData.InputText)
	if getDataErr != nil {
		log.Println("Get Data error:", getDataErr)
		http.Error(w, getDataError, http.StatusInternalServerError)
		return
	}

	// 將 JSON 內容回傳給前端
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Println("Send JSON to frontend...")
}

func main() {
	fmt.Println("Starting...")
	// Router
	router := http.NewServeMux()
	router.HandleFunc("/receiveInput", handleReceiveInput)
	// 設定靜態資源伺服器，指向存放HTML文件的資料夾
	router.Handle("/", http.FileServer(http.Dir("frontend")))

	// 啟動 HTTP 服務器，監聽在 9487 端口
	fmt.Printf("Server is running on http://localhost%s\n", port)
	http.ListenAndServe(port, router)
}
