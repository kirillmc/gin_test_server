package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	//"net/http/httputil"
	"time"

	"github.com/kirillmc/data_filler/pkg/filler"
	fil "github.com/kirillmc/data_filler/pkg/model"
	"github.com/kirillmc/gin_test_server/internal"
	"github.com/kirillmc/gin_test_server/internal/model"
)

//func SchetWithZagolovki(){
//	d, _ := httputil.DumpRequest(resp.Request, true)
//	log.Printf("requset: %v\nLEN: %v", string(d), len(d))
//	dump, err := httputil.DumpResponse(resp, true)
//	if err != nil {
//		fmt.Println("Error dumping response:", err)
//	}
//
//	fmt.Println(string(dump))
//
//	if err != nil {
//		log.Fatal("Failed to get number of programs:", err)
//	}
//	log.Printf("RESP REQUEST: %v\nlen(\"RESP REQUEST):%v\nRESP CONTENT LENGTH: %v\nRESPONSE:%v \nLENGTH:%v",
//		resp.Request, len([]byte(fmt.Sprintf("%v", resp.Request))),
//		resp.ContentLength,
//		resp,
//		len(dump))
//}

const (
	baseUrl     = "http://localhost:8081"
	newBaseUrl  = "https://ca32-93-100-98-132.ngrok-free.app"
	getPostfix  = "/programs/%d"
	postPostfix = "/programs/"
	avg         = 5
)

func getNProgramsClient(n int64) (internal.TrainPrograms, error) {
	resp, err := http.Get(fmt.Sprintf(baseUrl+getPostfix, n))
	if err != nil {
		log.Fatal("Failed to get programs:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return internal.TrainPrograms{}, fmt.Errorf("server code: %v", http.StatusNotFound)
	}

	if resp.StatusCode != http.StatusOK {
		return internal.TrainPrograms{}, fmt.Errorf("server code is not: %v, (code is: %v)", http.StatusOK, resp.StatusCode)
	}

	var programs internal.TrainPrograms
	if err = json.NewDecoder(resp.Body).Decode(&programs); err != nil {
		return internal.TrainPrograms{}, err
	}

	return programs, nil
}

func postNProgramsClient(programs fil.TrainPrograms) (model.Response, float64, error) {
	dataToSend, err := json.Marshal(programs)
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	resp, err := http.Post(fmt.Sprintf(baseUrl+postPostfix), "application/json", bytes.NewBuffer(dataToSend))
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return model.Response{Message: fmt.Sprintf("server code: %v", http.StatusNotFound)}, float64(len(dataToSend)), err
	}

	if resp.StatusCode != http.StatusCreated {
		return model.Response{Message: fmt.Sprintf("server code is not: %v, (code is: %v)", http.StatusOK, resp.StatusCode)}, float64(len(dataToSend)), err
	}

	var mess model.Response
	if err = json.NewDecoder(resp.Body).Decode(&mess); err != nil {
		log.Println(5)
		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	return mess, float64(len(dataToSend)), nil
}

func updateNProgramsClient(programs fil.TrainPrograms) (model.Response, float64, error) {
	dataToUpdate, err := json.Marshal(programs)
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(baseUrl+postPostfix), bytes.NewBuffer(dataToUpdate))
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return model.Response{Message: fmt.Sprintf("server code: %v", http.StatusNotFound)}, float64(len(dataToUpdate)), err
	}

	if resp.StatusCode != http.StatusOK {
		return model.Response{Message: fmt.Sprintf("server code is not: %v, (code is: %v)", http.StatusOK, resp.StatusCode)}, float64(len(dataToUpdate)), err

	}

	var mess model.Response
	if err = json.NewDecoder(resp.Body).Decode(&mess); err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	return mess, float64(len(dataToUpdate)), nil
}

func deleteNProgramsClient(n int64) (model.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(baseUrl+getPostfix, n), nil)
	if err != nil {
		return model.Response{Message: err.Error()}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Response{Message: err.Error()}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return model.Response{Message: fmt.Sprintf("server code: %v", http.StatusNotFound)}, err
	}

	if resp.StatusCode != http.StatusOK {
		return model.Response{Message: fmt.Sprintf("server code is not: %v, (code is: %v)", http.StatusOK, resp.StatusCode)}, err
	}

	var mess model.Response
	if err = json.NewDecoder(resp.Body).Decode(&mess); err != nil {
		return model.Response{Message: err.Error()}, err
	}

	return mess, nil
}

func main() {

	launchThirdTestWithNProgramsAndWg(25, 55)
}

func launchFirstTestForm0ToN(n int64) {
	fmt.Printf("SIZE - IS ONLY SIZE OF DATA(body)")

	fmt.Printf("\nMETHOD GET from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToGet)

	fmt.Printf("\nMETHOD POST from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToPost)

	fmt.Printf("\nMETHOD UPDATE from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToUpdate)

	fmt.Printf("\nMETHOD DELETE from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToDelete)

}

func launchSecondTestWithNProgramsAndWg(n int64, wgGroupCount int64) {
	fmt.Printf("SIZE - IS ONLY SIZE OF DATA(body)[%d]", n)

	fmt.Printf("\nMETHOD GET from 1 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToGet)

	fmt.Printf("\nMETHOD POST from 0 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToPost)

	fmt.Printf("\nMETHOD UPDATE from 0 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToUpdate)

	fmt.Printf("\nMETHOD DELETE from 0 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToDelete)
}

func launchThirdTestWithNProgramsAndWg(n int64, wgGroupCount int64) {
	fmt.Printf("SIZE - IS ONLY SIZE OF DATA(body) [USERS: %d]", wgGroupCount)

	//fmt.Printf("\nMETHOD GET from 1 to %d COUNTS:\n", n)
	//methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToGet)
	//
	//fmt.Printf("\nMETHOD POST from 0 to %d COUNTS:\n", n)
	//methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToPost)

	fmt.Printf("\nMETHOD UPDATE from 0 to %d COUNTS:\n", n)
	methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToUpdate)

	fmt.Printf("\nMETHOD DELETE from 0 to %d COUNTS:\n", n)
	methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToDelete)
}

// Будет увеличиваться объем данных от 0 до n, статично горутин wgGroupCount
func methodFrom0ToNWithAVGOfNGproutines(n int64, wgGroupCount int64, fun func(int64) (float64, float64)) {
	log.Printf("USERS;\tCOUNT;\tTIME(nanoS);\tSIZE(byte);\n")
	for i := int64(0); i <= n; i++ {
		printAvgOfGoroutines(i, wgGroupCount, fun)
	}
}

// Будет увеличиваться количетсово горутин от 0 до wgGroupCount, статично объем данных n
func methodWithConstAVGOfNGproutines(n int64, wgGroupCount int64, fun func(int64) (float64, float64)) {
	log.Printf("USERS;\tCOUNT;\tTIME(nanoS);\tSIZE(byte);\n")
	if wgGroupCount <= 0 {
		wgGroupCount = 1
	}
	for i := int64(1); i <= wgGroupCount; i++ {
		printAvgOfGoroutines(n, i, fun)
	}
}

func printAvgOfGoroutines(n int64, wgGroupCount int64, fun func(int64) (float64, float64)) {
	var durOfSize []float64
	var wg sync.WaitGroup
	wg.Add(int(wgGroupCount))
	// Создаем канал для передачи результатов из горутин
	resultСh := make(chan float64, wgGroupCount)
	sizeСhn := make(chan float64, wgGroupCount)

	// Создаем мьютекс для безопасного доступа к срезу result
	var mu sync.Mutex

	// Используем цикл для запуска горутин
	for i := int64(1); i <= wgGroupCount; i++ {
		go func(n int64) {
			defer wg.Done()
			dur, size := fun(n)
			resultСh <- dur
			sizeСhn <- size
		}(n)
	}

	wg.Wait()
	close(resultСh)
	close(sizeСhn)
	var size float64
	for res := range resultСh {
		mu.Lock()
		size = <-sizeСhn
		durOfSize = append(durOfSize, res)
		mu.Unlock()
	}

	log.Printf("\t%d;\t%d;\t%f;\t%f;\n", wgGroupCount, n, getAvgFromSlice(wgGroupCount, durOfSize), size)
}

func getAvgFromSlice(wgGroupCount int64, durOfSize []float64) float64 {
	var avgTime float64
	for i := int64(0); i < wgGroupCount; i++ {
		avgTime += durOfSize[i]
	}

	return avgTime / float64(wgGroupCount)
}

func methodFrom0ToNWithAVG(n int64, fun func(int64) (float64, float64)) {
	log.Printf("\tCOUNT;\tTIME(nanoS);\tSIZE(byte);\n")
	for i := int64(0); i <= n; i++ {
		printAvgOfConst(i, fun)
	}
}

func printAvgOfConst(n int64, fun func(int64) (float64, float64)) {
	var avgTime float64
	var avgSize float64
	for i := 1; i <= avg; i++ {
		avgTempTime, avgTempSize := fun(n)
		avgTime += avgTempTime
		avgSize += avgTempSize
	}

	log.Printf("\t%d;\t%f;\t%f;\n", n, avgTime/avg, avgSize/avg)
}

func oneToGet(n int64) (float64, float64) {
	start := time.Now()

	programs, err := getNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	numOfSets, err := json.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}

	return float64(end.Sub(start).Nanoseconds()), float64(len(numOfSets))
}

func oneToPost(n int64) (float64, float64) {
	start := time.Now()
	programs := filler.CreateOwnSetOfPrograms(int(n))
	_, postMessSize, err := postNProgramsClient(programs)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oneToUpdate(n int64) (float64, float64) {
	start := time.Now()
	programs := filler.CreateOwnSetOfPrograms(int(n))
	_, postMessSize, err := updateNProgramsClient(programs)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oneToDelete(n int64) (float64, float64) {
	start := time.Now()
	_, err := deleteNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	return float64(end.Sub(start).Nanoseconds()), 0.0
}

func getSizeInFormattedString(byteSize int64) string {
	if byteSize < 1024 {
		return fmt.Sprintf("%.3f байт\t", float64(byteSize))
	}
	if byteSize < 1024*1024 {
		return fmt.Sprintf("%.3f килобайт\t", float64(byteSize)/1024)
	} else {
		return fmt.Sprintf("%.3f мегабайт\t", float64(byteSize)/(1024*1024))
	}
}

func oldPrint(n int64) {
	start := time.Now()

	programs, err := getNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}

	end := time.Now()
	numOfSets, err := json.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}
	log.Printf("|\t\t\tGIN INFO: SIZE[%d]\t\t\t|\n", n)
	log.Printf("|\tTOTAL TIME TO GET PROGRAMS:\t%v\t\t|\n", end.Sub(start))
	log.Printf("|\tSIZE OF PROGRAMS:\t\t%s\t|\n", getSizeInFormattedString(int64(len(numOfSets))))
}

//func testRequest(i int) {
//	defer wg.Done()
//
//	start := time.Now()
//	_, err := getNProgramsClient(1)
//	if err != nil {
//		errCount++
//		//	log.Printf("failed to get user: ", err)
//	}
//	end := time.Now()
//	list = append(list, end.Sub(start))
//}
//
//func growGoroutine(n int) {
//	errCount = 0
//	start := time.Now()
//
//	wg.Add(n)
//	for i := 0; i < n; i++ {
//		go testRequest(i)
//	}
//	wg.Wait()
//	end := time.Now()
//	var total time.Duration
//	//	log.Printf("Last user info:%v\n", user)
//
//	for i := range list {
//		total += list[i]
//	}
//
//	avg := total.Nanoseconds() / (int64(len(list)))
//	log.Printf("%d get request;\t total time: %v;\tavg time: %s;\ttcount of errors: %d;\tpercentage of errors: %0.3f;\n",
//		n, end.Sub(start), time.Duration(avg), errCount, float64(errCount)/float64(n))
//}

//func main() {
//	//var user UserToGet
//	//var err error
//	//start := time.Now()
//	//for i := 0; i < 101; i++ {
//	//	user, err = getNProgramsClient(1)
//	//	if err != nil {
//	//		log.Fatal("failed to get user:", err)
//	//	}
//	//}
//	//end := time.Now()
//	//	var user UserToGet
//
//	//start := time.Now()
//	//n := 5
//	//wg.Add(n)
//	//for i := 0; i < n; i++ {
//	//	go testRequest(i)
//	//}
//	//wg.Wait()
//	//end := time.Now()
//	//var total time.Duration
//	////	log.Printf("Last user info:%v\n", user)
//	//
//	//for i := range list {
//	//	total += list[i]
//	//}
//	//
//	//avg := total.Nanoseconds() / (int64(len(list)))
//	//log.Printf("total requests time: %v\n", total)
//	//log.Printf("time for %d get requests: %v\n", n, end.Sub(start))
//	//log.Printf("avg time for %d get requests: %s\n", n, time.Duration(avg))
//	var n int64 = 55
//	log.Printf("\tCOUNT;\tTIME(nanoS);\tSIZE;\n")
//	for i := int64(1); i <= n; i++ {
//		printAvgOfConst(i)
//	}
//}
