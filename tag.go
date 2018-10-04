package main 

import ("fmt"
"os"
"log"
"bufio"
"net/http"
"encoding/json"
"bytes"
"container/heap"
"strings"
)
type PriorityQueue []*InputProbability
var storedMap map[string]PriorityQueue
func main() {
	storedMap = make(map[string]PriorityQueue)
	readInputFile(storedMap)
	fmt.Println("Done with storing data")
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	searchWord := r.URL.RawQuery
	searchArray := strings.Split(searchWord, "%20")
	searchWord = ""
	for _,value := range searchArray {
		searchWord += value + " "
	}
	searchWord = string(strings.Trim(searchWord, " "))
	pq := storedMap[searchWord]
	jsonObject,error := json.Marshal(pq)
	if error != nil {
		fmt.Println("Error occured while marshalling json for angular")
	}
	fmt.Fprintf(w,string(jsonObject))
}


func readInputFile(storedMap map[string]PriorityQueue) {
	file, err := os.Open("./fileList.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error finding the path")
	}

	defer file.Close()	
	scanner := bufio.NewScanner(file)
	var url *Url
	var image *Image
	var data *Data
	inputs := &Inputs{}
	inputs.Data = make([] *Data, 128)
	var i int64
	for scanner.Scan() {
		i++
		url = &Url{}
		url.Value = scanner.Text()
		image = &Image{}
		image.Url = url
		data = &Data{}
		data.Image = image
		inputs.Data[i-1] = data
		if i%128 == 0 {
			clarifaiCall(inputs,storedMap)
			inputs.Data = make([] *Data, 128)
			i = 0
		}				
	}
	inputs.Data = inputs.Data[:i]
	if len(inputs.Data) != 0 {
		clarifaiCall(inputs,storedMap)
	}		
}


func clarifaiCall(inputs *Inputs, storedMap map[string]PriorityQueue) {
	b, jsonError := json.Marshal(*inputs)
	if jsonError != nil {
		fmt.Println(jsonError)
		log.Fatalln("Error in creating json object")
	}
	client := &http.Client {}
	reader := bytes.NewReader(b)
	req, reqError := http.NewRequest("POST","https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs",reader)
	if reqError != nil {
		fmt.Println(reqError)
		log.Fatalln("Error while creating request")
	}
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Authorization","Key 67aa00be41914b9cb89532bfb8ccc135")
	resp, respError := client.Do(req)
	if respError != nil {
		fmt.Println(respError)
		log.Fatalln("Error while sending request")
	}
	response := &Response{}
	decoderErr := json.NewDecoder(resp.Body).Decode(&response)
	if decoderErr != nil {
		fmt.Println(decoderErr)
		log.Fatalln("Error while decoding response Json")		
	}
	for _,output := range response.Outputs {
		for _,concept := range output.OutputData.Concepts {
			inputProbability := &InputProbability{}
			inputProbability.Url = output.Input.Data.Image.Url
			inputProbability.Probability = concept.Value
			priorityQueue := make(PriorityQueue,0,10)
			tag := concept.Name
			if storedMap[tag] != nil {
				priorityQueue = storedMap[tag]
			}
			if len(priorityQueue) >= 10 {
				heap.Init(&priorityQueue)
				popInputProbability := heap.Pop(&priorityQueue).(*InputProbability)
				if inputProbability.Probability > popInputProbability.Probability {
					inputProbability.Index = 10;
					heap.Push(&priorityQueue, inputProbability)
				} else {
					popInputProbability.Index = 10;
					heap.Push(&priorityQueue, popInputProbability)
				}
			} else if len(priorityQueue) < 10 {
				inputProbability.Index = len(priorityQueue)
				heap.Push(&priorityQueue, inputProbability)
			}
			
			storedMap[tag] = priorityQueue
		}
	} 	
}