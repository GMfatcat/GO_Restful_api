package common

// Data structure for static json file
type Data struct {
	Lons   []float64 `json:"lons"`
	Lats   []float64 `json:"lats"`
	Counts []int     `json:"counts"`
}

// Data structure for Input from the frontend text box
type InputData struct {
	InputText string `json:"inputText"`
}
