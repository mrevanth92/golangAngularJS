package main

type Response struct  {
	Status Status `json:"status"`
	Outputs []*Outputs `json:"outputs"`
}

type Status struct {
	Code int64 `json:"code"`
	Description string `json:"description"`
}

type  Outputs struct{
	 OutputData *OutputData `json:"data"`
	 Input ResponseInput `json:"input"`
}

type ResponseInput struct {
	Id string `json:"id"`
	Data ResponseInputData `json:"data"`
}

type ResponseInputData struct {
	Image ResponseImageUrlData `json:"image"`
}

type ResponseImageUrlData struct {
	Url string `json:"url"`
}
type  OutputData struct{
	Concepts []*Concepts `json:"concepts"`
}

type Concepts struct {
	Id string   `json:"id"`
	Name string  `json:"name"`
	Value float64 `json:"value"`
	App_id string  `json:"app_id"`
}

type InputProbability struct {
	Probability float64
	Url string
	Index int
}