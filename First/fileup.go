package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"time"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io"
	"encoding/json"
)

func main() {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			dir := "/var/lib/motioneye/Camera1/"
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			for i, _ := range files {
				fileupload(dir, files[i].Name())	
			}
		}		
	}()
	
	http.ListenAndServe(":8094", nil)
}

func fileupload(dir string, fname string){
	url := "http://142.93.209.83:9194/upload"
	method := "POST"
	jsonfile, err := ioutil.ReadFile("details.json")
	if err != nil{
		fmt.Println(err)
	}
	var data (map[string]interface{})
	json.Unmarshal([]byte(jsonfile), &data)
	results := data
	//for j := 0; j < len(results); j++ {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(dir + fname)
	defer file.Close()
	part1,
			errFile1 := writer.CreateFormFile("myFile", filepath.Base(dir + fname))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 !=nil {
			
		fmt.Println(errFile1)
	}
	_ = writer.WriteField("key", results["key"].(string))
	err1 := writer.Close()
	if err1 != nil {
		fmt.Println(err)
	}
	client := &http.Client {Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	if res.StatusCode != 200{
		time.Sleep(2 * time.Second)
	}
	os.Remove(dir + fname)
}


